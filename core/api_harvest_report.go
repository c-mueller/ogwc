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

type HarvestReport struct {
	Generic HRGeneric `json:"generic"`
}

type HRGeneric struct {
	RrID                 string      `json:"rr_id"`
	EventTime            time.Time   `json:"event_time"`
	EventTimestamp       int         `json:"event_timestamp"`
	Coordinates          string      `json:"coordinates"`
	RecyclerCount        json.Number `json:"recycler_count"`
	RecyclerCapacity     json.Number `json:"recycler_capacity"`
	MetalInDebrisField   json.Number `json:"metal_in_debris_field"`
	CrystalInDebrisField json.Number `json:"crystal_in_debris_field"`
	MetalRetrieved       json.Number `json:"metal_retrieved"`
	CrystalRetrieved     json.Number `json:"crystal_retrieved"`
	OwnerName            string      `json:"owner_name"`
	OwnerAllianceName    string      `json:"owner_alliance_name"`
	OwnerAllianceTag     string      `json:"owner_alliance_tag"`
}

func (h HarvestReport) ToResources() Resources {
	m, _ := h.Generic.MetalRetrieved.Int64()
	c, _ := h.Generic.CrystalRetrieved.Int64()

	return Resources{
		Metal:     int(m),
		Crystal:   int(c),
		Deuterium: 0,
	}
}
