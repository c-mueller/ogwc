#!/usr/bin/env fish
if test -d build
    rm -rf build/
end

mkdir build

cd core

if test -e rice-box.go
    rm rice-box.go
end

rice embed-go

cd ..

cd cmd
go build -v -o ../build/ogwc