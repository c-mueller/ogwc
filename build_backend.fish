#!/usr/bin/env fish

## ogwc (https://github.com/c-mueller/ogwc).
## Copyright (C) 2018-2020 Christian Müller <dev@c-mueller.xyz>.
##
## This program is free software: you can redistribute it and/or modify
## it under the terms of the GNU Affero General Public License as published by
## the Free Software Foundation, either version 3 of the License, or
## (at your option) any later version.
##
## This program is distributed in the hope that it will be useful,
## but WITHOUT ANY WARRANTY; without even the implied warranty of
## MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
## GNU Affero General Public License for more details.
##
## You should have received a copy of the GNU Affero General Public License
## along with this program.  If not, see <https://www.gnu.org/licenses/>.

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
