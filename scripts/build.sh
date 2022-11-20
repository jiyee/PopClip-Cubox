#!/bin/bash

for arch in "arm64" "amd64"
do
  CGO_ENABLED=0 GOOS=darwin GOARCH="${arch}" go build -ldflags "-s -w" -o "./PopClip-Cubox.popclipext/cubox-macos-${arch}" ./src/cubox
done

zip -r "PopClip-Cubox.zip" "./PopClip-Cubox.popclipext"
mv -f "PopClip-Cubox.zip" "./Downloads/PopClip-Cubox.popclipextz"
