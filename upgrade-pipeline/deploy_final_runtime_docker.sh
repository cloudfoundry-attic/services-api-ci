#!/bin/bash

set -e -x

echo $0

docker pull $DOCKER_IMAGE
docker run -a stdout -a stderr -w /workspace/services-api-ci \
    -v $PWD:/workspace\
    $DOCKER_IMAGE \
    /bin/bash -l upgrade-pipeline/deploy_final_runtime.sh

exit 0
