#!/bin/bash

set -e -x -u

./hack/build.sh

# makes builds reproducible
export CGO_ENABLED=0
repro_flags="-ldflags=-buildid= -trimpath"

GOOS=darwin  GOARCH=amd64 go build $repro_flags -o dkg-darwin-amd64      .
GOOS=linux   GOARCH=amd64 go build $repro_flags -o dkg-linux-amd64       .
GOOS=windows GOARCH=amd64 go build $repro_flags -o dkg-windows-amd64.exe .

shasum -a 256 ./dkg-*-amd64*
