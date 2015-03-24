#!/bin/bash

set -e
set -x

DEPLOYMENT_NAME=v1-dummy-broker
RELEASE_NAME=v1-dummy-broker

cd /workspace/v1-dummy-broker-release
source /usr/local/share/chruby/chruby.sh
chruby 2.1.4


source /workspace/deployments-aws/${RELENG_ENV}/bosh_environment
bosh -n target ${BOSH_PREFIX}.${RELENG_ENV}.cf-app.com
./generate_deployment_manifest aws /workspace/deployments-aws/$RELENG_ENV/cf-aws-stub.yml /workspace/deployments-aws/$RELENG_ENV/cf-shared-secrets.yml > deployment.yml
bosh deployment deployment.yml

bosh create release --force --name $RELEASE_NAME
bosh -n upload release --skip-if-exists
bosh -n deploy
