#!/usr/bin/env fish
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

set -x version (git describe --exact-match --tags (git log -n1 --pretty='%h'))

if [ -z "$version" ]
  set -x version (git branch | grep \* | cut -d ' ' -f2)
end

set -x revision (git rev-list -1 --abbrev-commit HEAD)
set -x build_time (date)
set -x ctx (whoami)"@"(hostname)

echo "Building OGWC with Build Context:"
echo "Version: $version"
echo "Revision: $revision"
echo "Build Timestamp: $build_time"
echo "Build Environment: $ctx"

go build -v -ldflags "-X main.Revision=$revision -X \"main.Version=$version\" -X \"main.BuildTimestamp=$build_time\" -X main.BuildContext=$ctx" -o ../build/ogwc
