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
	"reflect"
	"strconv"
)

type Fleet struct {
	LightFighter  uint `json:"light_fighter" eid:"204"`
	HeavyFighter  uint `json:"heavy_fighter" eid:"205"`
	Cruiser       uint `json:"cruiser" eid:"206"`
	Battleship    uint `json:"battleship" eid:"207"`
	Battlecruiser uint `json:"battlecruiser" eid:"215"`
	Bomber        uint `json:"bomber" eid:"211"`
	Destroyer     uint `json:"destroyer" eid:"213"`
	Deathstar     uint `json:"deathstar" eid:"214"`

	Crawler    uint `json:"crawler" eid:"217"`
	Reaper     uint `json:"reaper" eid:"218"`
	Pathfinder uint `json:"pathfinder" eid:"219"`

	SmallCargo uint `json:"small_cargo" eid:"202"`
	LargeCargo uint `json:"large_cargo" eid:"203"`
	Recycler   uint `json:"recycler" eid:"209"`
	Colonyship uint `json:"colonyship"eid:"208"`
	Probe      uint `json:"probe" eid:"210"`
	Satellite  uint `json:"satellite" eid:"212"`

	RocketLauncher uint `json:"rocket_launcher" eid:"401"`
	LightLaser     uint `json:"light_laser" eid:"402"`
	HeavyLaser     uint `json:"heavy_laser" eid:"403"`
	GaussCannon    uint `json:"gauss_cannon" eid:"404"`
	IonCannon      uint `json:"ion_cannon" eid:"405"`
	PlasmaCannon   uint `json:"plasma_cannon" eid:"406"`
	SmallShield    uint `json:"small_shield" eid:"407"`
	LargeShield    uint `json:"large_shield" eid:"408"`
}

func (a Fleet) IsZero() bool {
	t := reflect.TypeOf(a)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		count := reflect.ValueOf(a).FieldByName(field.Name).Uint()

		if count != 0 {
			return false
		}
	}
	return true
}

func (a Fleet) GetCargoCapacity() int {
	var r int = 0

	t := reflect.TypeOf(a)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		id := field.Tag.Get("eid")

		cap := uint64(Entities[id].CargoCapacity)

		count := reflect.ValueOf(a).FieldByName(field.Name).Uint()

		r = r + int(int(cap)*int(count))
	}

	return r
}

func (a Fleet) ToResources() Resources {
	t := reflect.TypeOf(a)

	r := Resources{}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		id := field.Tag.Get("eid")

		cost := Entities[id].Cost

		count := reflect.ValueOf(a).FieldByName(field.Name).Uint()

		r = r.Add(cost.Mul(int(count)))
	}

	return r
}

func (a Fleet) Add(b Fleet) Fleet {
	c := Fleet{}
	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		valA := reflect.ValueOf(a).FieldByName(field.Name).Uint()
		valB := reflect.ValueOf(b).FieldByName(field.Name).Uint()

		reflect.ValueOf(&c).Elem().FieldByName(field.Name).SetUint(uint64(valA + valB))
	}

	return c
}

func (a Fleet) Sub(b Fleet) Fleet {
	c := Fleet{}
	t := reflect.TypeOf(c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		valA := reflect.ValueOf(a).FieldByName(field.Name).Uint()
		valB := reflect.ValueOf(b).FieldByName(field.Name).Uint()

		reflect.ValueOf(&c).Elem().FieldByName(field.Name).SetUint(uint64(valA - valB))
	}

	return c
}

func toFleet(m map[uint]uint) Fleet {
	l := Fleet{}
	t := reflect.TypeOf(l)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		id, _ := strconv.ParseInt(field.Tag.Get("eid"), 10, 32)

		count := m[uint(id)]

		reflect.ValueOf(&l).Elem().FieldByName(field.Name).SetUint(uint64(count))
	}

	return l
}
