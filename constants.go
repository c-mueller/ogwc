// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package ogwc

import "regexp"

var crRegex = regexp.MustCompile("cr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var rrRegex = regexp.MustCompile("rr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var mrRegex = regexp.MustCompile("mr-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")

var reportRegex = regexp.MustCompile("(rr|cr|mr)-[a-z]{2}-[0-9]{1,3}-[0-9a-f]{40}")
var uidRegex = regexp.MustCompile("[0-f]{8}-([0-f]{4}-){3}[0-f]{12}")
