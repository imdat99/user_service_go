#!/bin/bash

set -a
source .env
set +a

ARGS=$(awk -F= '!/^#/ && NF==2 {print "--build-arg " $1"=" $2}' .env)
echo "Building with args: $ARGS"
docker build $ARGS -t myapp .
