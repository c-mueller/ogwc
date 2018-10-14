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

echo "Building Angular UI"

ng build --prod

cd dist/ui
cp * ../../../app-ui/