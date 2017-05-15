#!/bin/sh

export VERSION="0.0.1"
export TAG="$VERSION-testing";

if [ "$TRAVIS_BRANCH" == "master" ]
then
    export TAG="$VERSION";
fi

if [ "$TRAVIS_BRANCH" == "develop" ]
then
    export TAG="$VERSION-deveelopment";
fi

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
echo "$TAG"
buildDockerImage
publishDockerImage

