// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/c-mueller/ogwc"
	"github.com/go-redis/redis"
)

func main() {
	app := ogwc.OGWCApplication{}
	app.Init(&redis.Options{
		Addr:"127.0.0.1:6379",
	})
	app.Serve(":16666")
}
