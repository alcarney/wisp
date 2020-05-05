#!/bin/bash

# Script to bootstrap tinygo in Github Actions

url="https://github.com/tinygo-org/tinygo/releases/download/v${VERSION}/tinygo_${VERSION}_amd64.deb"
curl -s $url -o "tinygo_${VERSION}.deb"

ls -lh

sudo apt install "./tinygo_${VERSION}.deb"
tinygo