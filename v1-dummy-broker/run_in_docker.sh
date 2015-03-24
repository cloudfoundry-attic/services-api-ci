#!/bin/bash
set -x
REGISTRY="docker.gocd.cf-app.com:5000"

SSH_KEY_LOCATION=${SSH_KEY_LOCATION:-/var/vcap/jobs/gocd-agent/id_rsa}

./gocd-scripts/with_docker_cleanup docker run --privileged \
  -e ENV=${ENV} \
  -e RELENG_ENV=${RELENG_ENV} \
  -e BOSH_PREFIX=${BOSH_PREFIX} \
  -e RELEASE_ACCESS_KEY=${RELEASE_ACCESS_KEY} \
  -e RELEASE_SECRET_KEY=${RELEASE_SECRET_KEY} \
  -v `pwd`:/workspace \
  -v ${SSH_KEY_LOCATION}:/root/.ssh/id_rsa \
  $REGISTRY/library/services-ci ssh-agent $@
