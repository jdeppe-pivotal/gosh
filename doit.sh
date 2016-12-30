#!/usr/bin/env bash

set -e -x

WORKDIR=$( cd $(dirname $0); pwd )

pushd $WORKDIR
  GOPATH=$PWD go build gosh
  ./gosh 400
popd
