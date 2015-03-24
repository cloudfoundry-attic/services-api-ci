#!/bin/bash

set -e -x

echo $0

docker pull $DOCKER_IMAGE
docker run -a stdout -a stderr -w /workspace/services-api-ci \
    -v $PWD:/workspace\
    -e DEPLOYMENTS_REPO_NAME \
    $DOCKER_IMAGE \
    /bin/bash -l upgrade-pipeline/delete_deployment.sh

exit 0
