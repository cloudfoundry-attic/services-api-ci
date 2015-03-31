#!/bin/bash

set -e
set -x

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

RELEASE_DIRECTORY=cf-release-master

cd /workspace/$RELEASE_DIRECTORY

function latest_release() {
  ls releases/cf-*            | \
  cut -d '-' -f2              | \
  cut -d '.' -f1              | \
  sort -n                     | \
  tail -1
}

function make_manifest() {
  BOSH_RELEASES_DIR=/workspace \
  CF_RELEASE_DIR=/workspace/$RELEASE_DIRECTORY \
  ./generate_deployment_manifest aws \
    /workspace/deployments-services-api/wasabi/cf-aws-stub.yml \
    /workspace/deployments-services-api/wasabi/cf-shared-secrets.yml \
    > deployment.yml
  bosh -n deployment deployment.yml
}

function bosh_deploy() {
  FINAL_RELEASE_VERSION=$(latest_release)
  bosh -n --parallel 3 upload release releases/cf-${FINAL_RELEASE_VERSION}.yml --skip-if-exists
  bosh -n deploy
}

bosh -n target bosh.wasabi.cf-app.com 
make_manifest
bosh_deploy
