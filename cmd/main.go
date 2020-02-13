// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
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
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
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

	useRedis = serverCmd.Flag("redis", "Use Redis based repository").Short('R').Bool()
	boltPath = serverCmd.Flag("bolt-path", "Path to the BoltDB repo").Short('p').Default("ogwc.db").String()

	ogameApiUrl = serverCmd.Flag("ogame-api-proxy-url", "The url to the Ogame API Proxy").Default("https://ogapi.rest/v1/report/%s/0").String()

	metricsUsers = serverCmd.Flag("metrics-user", "Add a user for Metrics").Short('u').Strings()

	prodFlag = serverCmd.Flag("dev", "Set to true to make Gin run in Development mode").Bool()

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
		if !*prodFlag {
			gin.SetMode(gin.ReleaseMode)
		}

		log.Infof("Running revision %q built at %s on %s", Revision, BuildTimestamp, BuildContext)
		log.Info("Launching OGWC server...")

		accs := make([]ogwc.MetricsUserAccount, 0)
		for _, us := range *metricsUsers {
			splitUser := strings.Split(us, ":")
			if len(splitUser) != 2 {
				log.Errorf("Invalid Metrics user definition. %q", us)
				log.Error("Users have to be defined in the following format:")
				log.Error("<USERNAME>:<PASSWORD>, where USERNAME and Password may\n" +
					"only consist out of numbers (0-9), upper and lower case letters (a-z and A-Z).")
				os.Exit(1)
			}
			acc := ogwc.MetricsUserAccount{
				Username: splitUser[0],
				Password: splitUser[1],
			}

			log.Infof("Added Metrics account with Username: %q and Password: %q", acc.Username, acc.Password)
			accs = append(accs, acc)
		}

		app := ogwc.OGWCApplication{
			APIUrlTemplate: *ogameApiUrl,
		}

		if len(accs) > 0 {
			app.UserAccounts = accs
		}

		app.InitVersionInfo(BuildContext, Revision, Version, BuildTimestamp)
		app.Init(*useRedis, *boltPath, &redis.Options{
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
