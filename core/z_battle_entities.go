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

package core

// NOTE: The name of this file is intentional, to ensure the init method of this file is invoked after the one of rice-box.go

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
)

var Entities map[string]EntityType

func init() {
	cfgBucket := rice.MustFindBox("config")
	data, _ := cfgBucket.Bytes("entities.json")

	json.Unmarshal(data, &Entities)
}

type EntityType struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Cost          Resources `json:"cost"`
	CargoCapacity uint      `json:"cargo_capacity"`
}
