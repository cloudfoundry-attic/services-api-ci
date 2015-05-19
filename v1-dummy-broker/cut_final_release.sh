#!/bin/bash

set -ex

SCRIPTS_DIR=/workspace/services-api-ci/v1-dummy-broker

DEVELOP_BRANCH=${DEVELOP_BRANCH:-develop}
MASTER_BRANCH=${MASTER_BRANCH:-master}

# We will remove chruby in #84317342
hash ruby 2>/dev/null || {
    set +u
    echo "ruby not found, using chruby"
    source /usr/local/share/chruby/chruby.sh
    chruby 2.1.4
    set -u
}

cd /workspace/v1-dummy-broker-release

# Merge release-candidate branch into master (should be a fast-forward)
ssh-add /root/.ssh/id_rsa
ssh-add -l
git fetch origin $MASTER_BRANCH

set +e
git log origin/develop ^origin/master | grep commit
EXIT_CODE=$?
if [[ $EXIT_CODE -eq 1 ]]; then
  echo "No changes to the release. Exiting early and not cutting a new final release."
  exit 0
fi
set -e

git checkout $MASTER_BRANCH
git merge origin/release-candidate

cat << EOF > config/private.yml
---
blobstore:
  s3:
    access_key_id: $RELEASE_ACCESS_KEY
    secret_access_key: $RELEASE_SECRET_KEY
EOF

#Creating final release
bosh -n create release --final --with-tarball

#Pushing final release to master
FINAL_RELEASE_VERSION=$(ruby $SCRIPTS_DIR/final_release_version.rb v1-dummy-broker)
if [ -z $? ]; then
  echo "No final release version found - exiting."
  exit 1
fi

git config --global user.email "gocd-bot@pivotal.io"
git config --global user.name "Final Release Builder"
git add .
git commit -m "Final release version $FINAL_RELEASE_VERSION"
git push origin $MASTER_BRANCH
git tag v$FINAL_RELEASE_VERSION
git push origin refs/tags/v$FINAL_RELEASE_VERSION:refs/tags/v$FINAL_RELEASE_VERSION

# Merge the new commit back into develop
git fetch origin $DEVELOP_BRANCH
git checkout $DEVELOP_BRANCH
git merge $MASTER_BRANCH
git push origin $DEVELOP_BRANCH
