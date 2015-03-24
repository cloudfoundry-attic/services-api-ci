#!/bin/bash

set -x
set -e

PRODUCT=$1
DEPLOYMENT_NAME=cf
RELEASE_NAME=cf

cd /workspace

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

# If there is a file with the api address, then we must be targetting
# a transient bosh-lite instance. Otherwise, we're targetting AWS.
if [ -f api-address ]; then
  export BOSH_USER=admin
  export BOSH_PASSWORD=admin
  export BOSH_RELEASES_DIR=/workspace
  export BOSH_LITE_IP=$(cat api-address)
  export DEPLOYMENT_NAME=cf
  bosh -n target ${BOSH_LITE_IP}

  set +e # Bosh will exit with error if the deployment is not found
  bosh -n delete deployment ${DEPLOYMENT_NAME} --force
  bosh -n delete release ${RELEASE_NAME}
  set -e
else
  echo "Not on bosh lite"
  exit 10
fi
