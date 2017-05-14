#!/bin/sh

if [ "$TRAVIS_BRANCH" == /^master/ ]; then
    export TAG="latest";
else if [ "$TRAVIS_BRANCH" == /^develop$/ ]; then
    export TAG="develop";
else
    export TAG="feature";
fi

docker build -t $REPO:$TAG -f Dockerfile .

docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASS
docker tag $REPO:$COMMIT $REPO:$TAG
docker tag $REPO:$COMMIT $REPO:travis-$TRAVIS_BUILD_NUMBER
docker push $REPO