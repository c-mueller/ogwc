// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2020 Christian Müller <dev@c-mueller.xyz>.
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

package core

import (
	"encoding/json"
	"time"
)

type MissileReport struct {
	Generic MissileReportGeneric `json:"generic"`
	Details MissileReportDetails `json:"details"`
}

type MissileReportDetails struct {
	Defense          []interface{} `json:"defense"`
	DefenseDestroyed []interface{} `json:"defense_destroyed"`
}

type MissileReportGeneric struct {
	MrID                      string      `json:"mr_id"`
	EventTime                 time.Time   `json:"event_time"`
	EventTimestamp            json.Number `json:"event_timestamp"`
	AttackerName              string      `json:"attacker_name"`
	AttackerPlanetName        string      `json:"attacker_planet_name"`
	AttackerPlanetCoordinates string      `json:"attacker_planet_coordinates"`
	AttackerPlanetType        json.Number `json:"attacker_planet_type"`
	DefenderName              string      `json:"defender_name"`
	DefenderPlanetName        string      `json:"defender_planet_name"`
	DefenderPlanetCoordinates string      `json:"defender_planet_coordinates"`
	DefenderPlanetType        json.Number `json:"defender_planet_type"`
	MissilesLostAttacker      json.Number `json:"missiles_lost_attacker"`
	MissilesLostDefender      json.Number `json:"missiles_lost_defender"`
}

func (r MissileReport) GetLosses(attacker bool) Resources {
	if attacker {
		v, _ := r.Generic.MissilesLostAttacker.Int64()
		return IPMCost.Mul(int(v))
	} else {
		//TODO Implement Computation of cost for lost defenses/ AntiIPMs
		return Resources{}
	}
}