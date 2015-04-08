#!/bin/bash

set -e -x

echo $0

docker pull $DOCKER_IMAGE
docker run -a stdout -a stderr -w /workspace/services-api-ci \
    -v $PWD:/workspace\
    -e "RELENG_ENV=$RELENG_ENV" \
    -e "BOSH_USER=$BOSH_USER" \
    -e "BOSH_PASSWORD=$BOSH_PASSWORD" \
    $DOCKER_IMAGE \
    /bin/bash -l upgrade-pipeline/deploy_develop_runtime.sh

exit 0
