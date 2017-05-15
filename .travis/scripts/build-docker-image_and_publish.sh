#!/bin/sh

export VERSION="0.0.1"
export TAG="$VERSION-testing";


buildDockerImage(){
    docker build -t $REPO:$TAG -f Dockerfile .
    return 0
}

publishDockerImage(){
    docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASS
    docker tag $REPO:$COMMIT $REPO:$TAG
    docker tag $REPO:$COMMIT $REPO:travis-$TRAVIS_BUILD_NUMBER
    docker push $REPO
    return 0
}

if [ "$TRAVIS_BRANCH" == "master" ]
then
    export TAG="$VERSION";
    buildDockerImage
    publishDockerImage
fi

if [ "$TRAVIS_BRANCH" == "develop" ]
then
    export TAG="$VERSION-development";
    buildDockerImage
    publishDockerImage
fi




