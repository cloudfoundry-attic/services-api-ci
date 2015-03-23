#!/bin/bash

set -e
set -x

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

set -u

RELEASE_DIRECTORY=cf-release-master
BOSH_LITE_IP=$(cat /workspace/api-address)

cd /workspace/$RELEASE_DIRECTORY

function latest_release() {
  ls releases/cf-*            | \
  cut -d '-' -f3              | \
  cut -d '.' -f1              | \
  sort -n                     | \
  tail -1
}

function make_manifest() {
  BOSH_RELEASES_DIR=/workspace \
  CF_RELEASE_DIR=/workspace/$RELEASE_DIRECTORY \
  ./bosh-lite/make_manifest
}

function customize_manifest() {
  MANIFEST_FILE=$PWD/bosh-lite/manifests/cf-manifest.yml
  sed -i "s/10.244.0.34.xip.io/$BOSH_LITE_IP.xip.io/g" $MANIFEST_FILE
  # sed -i "s/^name\: cf/name\: cf-warden/g" $MANIFEST_FILE
}

function bosh_deploy() {
  FINAL_RELEASE_VERSION=$(latest_release)
  bosh -n upload release releases/cf-${FINAL_RELEASE_VERSION}.yml --skip-if-exists
  bosh -n deploy
}

export BOSH_USER=admin
export BOSH_PASSWORD=admin

bosh -n target ${BOSH_LITE_IP}.xip.io
make_manifest
customize_manifest
bosh_deploy
