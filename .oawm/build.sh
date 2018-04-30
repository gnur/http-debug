#!/bin/sh

set -x
export PROJECT=http-debug
echo '------------------------------'

mkdir -p /go/src/github.com/gnur/
ln -s /src /go/src/github.com/gnur/${PROJECT}
cd /go/src/github.com/gnur/${PROJECT}/

export CGO_ENABLED=0
export GOOS=linux

cp -a /src/.oawm /dist/
go build -a -installsuffix cgo -o /dist/app .
