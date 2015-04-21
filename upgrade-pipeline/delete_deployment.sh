#!/bin/bash

set -x
set -e

RELEASE_NAME=cf

cd /workspace

source /usr/local/share/chruby/chruby.sh
chruby 2.1.6

export BOSH_RELEASES_DIR=/workspace
RELENG_ENV=$RELENG_ENV:-wasabi}
export DEPLOYMENT_NAME=cf-wasabi

bosh -n target bosh.wasabi.cf-app.com 

set +e # Bosh will exit with error if the deployment is not found
bosh -n delete deployment ${DEPLOYMENT_NAME} --force
bosh -n delete release ${RELEASE_NAME}
set -e
