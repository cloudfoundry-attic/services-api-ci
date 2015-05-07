#!/bin/bash

set -e
set -x

DEPLOYMENT_NAME=v1-dummy-broker
RELEASE_NAME=v1-dummy-broker

upload_release() {
  set +e
    bosh -n -N upload release --rebase &>output.txt
    local status=$?
  set -e

  if [[ $(cat output.txt) != *"Error 100: Rebase"* && $status -ne 0 ]]
  then
    echo "failing"
    exit 1
  fi
}

cd /workspace/v1-dummy-broker-release
source /usr/local/share/chruby/chruby.sh
chruby 2.1.4


source /workspace/deployments-runtime/${RELENG_ENV}/bosh_environment
bosh -n target ${BOSH_PREFIX}.${RELENG_ENV}.cf-app.com
./generate_deployment_manifest aws \
  /workspace/deployments-runtime/$RELENG_ENV/cf-aws-stub.yml \
  /workspace/deployments-runtime/$RELENG_ENV/cf-shared-secrets.yml \
  /workspace/deployments-core-services/$RELENG_ENV/v1-dummy-broker-secrets.yml \
  /workspace/deployments-core-services/$RELENG_ENV/cf-properties.yml \
  > deployment.yml
bosh deployment deployment.yml

bosh create release --force --name $RELEASE_NAME

upload_release

bosh -n deploy
