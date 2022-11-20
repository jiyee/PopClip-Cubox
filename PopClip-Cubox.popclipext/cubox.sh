#!/bin/sh

export BIN="./cubox-macos"
export ARCH=`uname -m`

if [ "$ARCH" = "x86_64" ]; then
  BIN="$BIN-amd64"
else
  BIN="$BIN-arm64"
fi

$BIN "$@"
