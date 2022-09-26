FROM golang:alpine AS builder

RUN apk add git

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    APP_ENV=prod

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Copy and download dependency using go mod
RUN go mod download

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine

# Vault config
RUN apk add --update --upgrade --no-cache ca-certificates curl curl-dev bash
RUN curl https://releases.hashicorp.com/envconsul/0.8.0/envconsul_0.8.0_linux_amd64.zip -O && \
    unzip envconsul_0.8.0_linux_amd64.zip -d /bin && \
    rm envconsul_0.8.0_linux_amd64.zip

COPY --from=builder /dist/main /
COPY --from=builder /build/startup.sh /

# Command to run
CMD ["/bin/sh", "/startup.sh"]
