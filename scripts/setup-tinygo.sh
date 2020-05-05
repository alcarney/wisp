#!/bin/bash

# Script to bootstrap tinygo in Github Actions

url="https://github.com/tinygo-org/tinygo/releases/download/v${VERSION}/tinygo_${VERSION}_amd64.deb"
curl -s $url --output "tinygo_${VERSION}.deb"

apt install "tinygo_${VERSION}.deb"
tinygo