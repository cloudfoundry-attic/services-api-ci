#!/bin/bash

set -e
set -x

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

RELEASE_DIRECTORY=cf-release-develop
RELENG_ENV=${RELENG_ENV:-wasabi}

cd /workspace/$RELEASE_DIRECTORY

function make_manifest() {
  BOSH_RELEASES_DIR=/workspace \
  CF_RELEASE_DIR=/workspace/$RELEASE_DIRECTORY \
    ./generate_deployment_manifest aws \
    /workspace/deployments-services-api/$RELENG_ENV/cf-aws-stub.yml \
    /workspace/deployments-services-api/$RELENG_ENV/cf-shared-secrets.yml \
    > deployment.yml
  bosh -n deployment deployment.yml
}

function bosh_deploy() {
  bosh -n create release --force
  bosh -n upload release --skip-if-exists
  bosh -n deploy
}

bosh -n target bosh.$RELENG_ENV.cf-app.com 
make_manifest
bosh_deploy
