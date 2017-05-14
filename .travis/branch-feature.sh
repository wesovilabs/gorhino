#!/bin/bash -e

: "${REPO?}" "${TAG?}"
: "${COMMIT?}" "${TRAVIS_BUILD_NUMBER?}"
: "${DOCKER_EMAIL?}" "${DOCKER_USER?}" "${DOCKER_PASS?}"


docker build -t $REPO:feature -f Dockerfile .
docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASS
docker tag $REPO:$COMMIT $REPO:feature
docker tag $REPO:$COMMIT $REPO:travis-$TRAVIS_BUILD_NUMBER
docker push $REPO

return 0