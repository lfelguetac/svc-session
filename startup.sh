#!/bin/sh

for file in $(ls /run/secrets/variables*.sh 2>/dev/null); do
  source $file
done

VAULT_CONFIG_FILE=""
if [ -f /run/secrets/config.hcl ]; then
  VAULT_CONFIG_FILE=/run/secrets/config.hcl
fi

if [ -n "$VAULT_CONFIG_FILE" ]; then
  envconsul -config=/run/secrets/config.hcl ./main
else
  ./main
fi
