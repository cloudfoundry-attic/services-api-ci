#!/bin/bash

set -e
set -x

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

RELEASE_DIRECTORY=cf-release-develop
BOSH_LITE_IP=$(cat /workspace/api-address)

cd /workspace/$RELEASE_DIRECTORY

function latest_release_number() {
  ls dev_releases/cf/cf-* | \
  cut -d '-' -f2 | \
  cut -d '+' -f1 | \
  sort -n | \
  tail -1
}

function latest_dev_release_number() {
  ls dev_releases/cf-`latest_release_number`+dev.* | \
  cut -d '-' -f2                                   | \
  cut -d '.' -f2                                   | \
  sort -n                                          | \
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
}

function bosh_deploy() {
  bosh create release --force
  bosh -n upload release "dev_releases/cf/cf-`latest_release_number`+dev.`latest_dev_release_number`.yml" --skip-if-exists
  bosh -n deploy
}

export BOSH_USER=admin
export BOSH_PASSWORD=admin

bosh -n target ${BOSH_LITE_IP}.xip.io
make_manifest
customize_manifest
bosh_deploy
