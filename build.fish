#!/usr/bin/env fish
if test -d build
    rm -rf build/
end

if test -d app-ui
    rm -rf app-ui
end

mkdir build
mkdir app-ui

cd ui

ng build --prod

cd dist/ui
cp * ../../../app-ui/
cd ../../../

if test -e rice-box.go
    rm rice-box.go
end

rice embed-go

cd core

if test -e rice-box.go
    rm rice-box.go
end

rice embed-go

cd ..

cd cmd
go build -v -o ../build/ogwc