#!/usr/bin/env fish
./build.fish
docker build . -t halive/ogwc:latest
