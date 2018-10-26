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

set -x revision (git rev-list -1 --abbrev-commit HEAD)
set -x build_time (date)
set -x ctx (whoami)"@"(hostname)

go build -v -ldflags "-X main.Revision=$revision -X \"main.BuildTimestamp=$build_time\" -X main.BuildContext=$ctx" -o ../build/ogwc