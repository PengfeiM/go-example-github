#!/bin/bash

PWD=$(
  cd "$(dirname "$0")"
  pwd
)
ROOT_PATH=${PWD}/..

# remove old output
rm -f $ROOT_PATH/output/*

cd $ROOT_PATH
make
