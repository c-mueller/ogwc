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
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	serverCmd = kingpin.Command("server", "Launch the OGWC server Application")

	bindUrl = serverCmd.Flag("bind-address", "The address for the HTTP server to bind to").Default(":16666").String()

	redisAddr = serverCmd.Flag("redis-address", "The Address (including Port) to the redis server").Default("127.0.0.1:6379").String()
	redisPass = serverCmd.Flag("redis-password", "The password for the redis server").Default("").String()
	redisDb   = serverCmd.Flag("redis-database", "The index of the redis db to use").Default("0").Int()

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
		log.Info("Launching OGWC server...")
		app := ogwc.OGWCApplication{}
		app.Init(&redis.Options{
			Addr:     *redisAddr,
			DB:       *redisDb,
			Password: *redisPass,
		})
		app.Serve(*bindUrl)
	}
}

func initializeLogger() {
	stdoutBackend := logging.NewLogBackend(os.Stdout, "", 0)

	backendFormatter := logging.NewBackendFormatter(stdoutBackend, format)

	leveledBackend := logging.AddModuleLevel(backendFormatter)

	leveledBackend.SetLevel(logging.DEBUG, "")

	logging.SetBackend(leveledBackend)
}