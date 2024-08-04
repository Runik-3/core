#!/bin/bash

GREEN=$'\e[0;32m'
RESET=$'\e[0m'

if [[ -z $1 ]];
  then
  echo "please provide a release version"
fi

if [[ $(git describe --tags) != $1 ]];
then 
  echo $1 > version
  git add version
  git commit -m "build: $1"
  git tag $1

  echo "$GREEN
$1 created. To finalize release, push the tag with$RESET
  git push tag origin [tagname]"
fi
