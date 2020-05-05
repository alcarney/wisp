#!/bin/bash

# Script to bootstrap tinygo in Github Actions

url="https://github.com/tinygo-org/tinygo/releases/download/v${VERSION}/tinygo_${VERSION}_amd64.deb"
wget $url -O "tinygo_${VERSION}.deb"

ls -lh

sudo apt install "./tinygo_${VERSION}.deb"

# Fix the path since tinygo does not seem to install itself to a standard location?
export PATH=$PATH:/usr/local/tinygo/bin/
