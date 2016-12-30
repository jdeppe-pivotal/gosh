#!/usr/bin/env bash

#!/usr/local/bin/dumb-init /bin/bash

set -e -x

WORKDIR=$( cd $(dirname $0); pwd )

pushd $WORKDIR
  GOPATH=$PWD go build gosh
  ./gosh 400 &
  sleep 2
popd
