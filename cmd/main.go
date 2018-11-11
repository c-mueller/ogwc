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

package main

import (
	"fmt"
	"github.com/c-mueller/ogwc"
	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// Build Params
var IsDev = true
var Version = "master"
var Revision string
var BuildTimestamp string
var BuildContext string

var (
	serverCmd = kingpin.Command("server", "Launch the OGWC server Application")

	bindUrl = serverCmd.Flag("bind-address", "The address for the HTTP server to bind to").Default(":16666").String()

	redisAddr = serverCmd.Flag("redis-address", "The Address (including Port) to the redis server").Default("127.0.0.1:6379").String()
	redisPass = serverCmd.Flag("redis-password", "The password for the redis server").Default("").String()
	redisDb   = serverCmd.Flag("redis-database", "The index of the redis db to use").Default("0").Int()

	ogameApiUrl = serverCmd.Flag("ogame-api-proxy-url", "The url to the Ogame API Proxy").Default("https://ogapi.rest/v1/report/%s/0").String()

	versionCmd = kingpin.Command("version", "Show version information")

	cmd = ""
)

var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05} - %{level}] - %{module}:%{color:reset} %{message}`,
)

var log = logging.MustGetLogger("main")

func init() {
	initializeLogger()

	cmd = kingpin.Parse()
}

func main() {
	switch cmd {
	case "server":
		log.Infof("Running revision %q built at %s on %s", Revision, BuildTimestamp, BuildContext)
		log.Info("Launching OGWC server...")
		app := ogwc.OGWCApplication{
			APIUrlTemplate: *ogameApiUrl,
		}
		app.InitVersionInfo(BuildContext, Revision, Version, BuildTimestamp)
		app.Init(&redis.Options{
			Addr:     *redisAddr,
			DB:       *redisDb,
			Password: *redisPass,
		})
		app.Serve(*bindUrl)
	case "version":
		fmt.Printf("OGWC -  OGame Win Calculator\n"+
			"Version: %s.%s\n"+
			"UI Revision: %s\n"+
			"Built on: %s\n"+
			"Built at: %s\n", Version, Revision, ogwc.GetUIRevision(), BuildContext, BuildTimestamp)
	}
}

func initializeLogger() {
	stdoutBackend := logging.NewLogBackend(os.Stdout, "", 0)

	backendFormatter := logging.NewBackendFormatter(stdoutBackend, format)

	leveledBackend := logging.AddModuleLevel(backendFormatter)

	leveledBackend.SetLevel(logging.DEBUG, "")

	logging.SetBackend(leveledBackend)
}
