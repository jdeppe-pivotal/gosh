#!/usr/bin/env bash

#!/usr/local/bin/dumb-init /bin/bash

set -x

WORKDIR=$( cd $(dirname $0); pwd )

pushd $WORKDIR
  gcc -o zombie zombie.c
  X=0
  while [ $X -lt $1 ]; do
    ./zombie
    ((X = X + 1))
  done
popd
