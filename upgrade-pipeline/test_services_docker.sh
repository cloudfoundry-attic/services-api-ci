#!/bin/bash

set -e -x

echo $0

docker pull $DOCKER_IMAGE
docker run -a stdout -a stderr -w /workspace/services-api-ci \
    -v $PWD:/workspace\
    -e CF_PASSWORD \
    -e RELENG_ENV \
    -e CF_TRACE \
    -e SLEEP_TIME \
    $DOCKER_IMAGE \
    /bin/bash -l upgrade-pipeline/test_services.sh

exit 0
