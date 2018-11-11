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

package core

import "github.com/google/uuid"

type Resources struct {
	Metal     int `json:"metal"`
	Crystal   int `json:"crystal"`
	Deuterium int `json:"deuterium"`
}

type IdentifiableResources struct {
	Resources
	ID string `json:"id"`
}

func (r *Resources) ToIdentifiableResources() IdentifiableResources {
	uid := uuid.New().String()

	return IdentifiableResources{*r, uid}
}

func (r Resources) IsZeroWithTolerance() bool {
	v := r.Metal >= -100 && r.Metal <= 100
	v = v && r.Crystal >= -100 && r.Crystal <= 100
	v = v && r.Deuterium >= -100 && r.Deuterium <= 100

	return v
}

func (r Resources) Add(b Resources) Resources {
	return Resources{
		Metal:     r.Metal + b.Metal,
		Crystal:   r.Crystal + b.Crystal,
		Deuterium: r.Deuterium + b.Deuterium,
	}
}

func (r Resources) Mul(x int) Resources {
	return Resources{
		Metal:     r.Metal * x,
		Crystal:   r.Crystal * x,
		Deuterium: r.Deuterium * x,
	}
}

func (r Resources) MulF(x float64) Resources {
	return Resources{
		Metal:     int(float64(r.Metal) * x),
		Crystal:   int(float64(r.Crystal) * x),
		Deuterium: int(float64(r.Deuterium) * x),
	}
}

func (r Resources) Sub(b Resources) Resources {
	return Resources{
		Metal:     r.Metal - b.Metal,
		Crystal:   r.Crystal - b.Crystal,
		Deuterium: r.Deuterium - b.Deuterium,
	}
}
