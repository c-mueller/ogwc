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

package repo

import (
	"encoding/json"
	"github.com/c-mueller/ogwc/core"
	"github.com/go-redis/redis"
)

type Repository struct {
	Options redis.Options
	client  *redis.Client
}

func (r *Repository) Connect() {
	r.client = redis.NewClient(&r.Options)
}

func (r *Repository) Insert(calculation core.CombatReportCalculation) string {
	data, _ := json.Marshal(calculation)

	uidString := ""
	var e *core.CombatReportCalculation
	for e != nil || len(uidString) == 0 {
		uidString = generateID()
		e = r.Get(uidString)
	}

	r.client.Set(uidString, data, -1)

	return uidString
}

func (r *Repository) Update(id string, calc core.CombatReportCalculation) error {
	data, _ := json.Marshal(calc)

	cmd := r.client.Set(id, data, -1)

	return cmd.Err()
}

func (r *Repository) Get(id string) *core.CombatReportCalculation {
	cmd := r.client.Get(id)

	if cmd.Err() != nil {
		return nil
	}

	data, err := cmd.Bytes()
	if err != nil {
		return nil
	}

	var calc core.CombatReportCalculation

	err = json.Unmarshal(data, &calc)
	if err != nil {
		return nil
	}

	return &calc
}
