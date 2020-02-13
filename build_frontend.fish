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