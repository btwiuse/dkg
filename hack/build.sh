#!/bin/bash

set -e -x -u

# makes builds reproducible
export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath"

go fmt

# export GOOS=linux GOARCH=amd64
go build $repro_flags -o dkg .
./dkg version

echo "Success"
