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

type Resources struct {
	Metal     int `json:"metal"`
	Crystal   int `json:"crystal"`
	Deuterium int `json:"deuterium"`
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
