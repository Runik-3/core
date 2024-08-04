#!/bin/bash

if [[ -z $1 ]];
  then
  echo "please provide a release version"
fi

if [[ $(git describe --tags) != $1 ]];
then 
  $1 > version
  git add version
  git commit -m "build: $1"
  git tag $1
fi
