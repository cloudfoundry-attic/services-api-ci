#!/bin/bash

set -e
set -x

source /usr/local/share/chruby/chruby.sh
chruby 2.1.4

source /workspace/deployments-runtime/${RELENG_ENV}/bosh_environment
bosh -n target ${BOSH_PREFIX}.${RELENG_ENV}.cf-app.com

bosh download manifest v1-dummy-broker > v1-dummy-broker.yml
bosh deployment v1-dummy-broker.yml

echo "Running acceptance tests"
bosh run errand acceptance-tests

