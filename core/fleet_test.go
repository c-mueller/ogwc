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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFleet_Add(t *testing.T) {
	a := Fleet{
		Battlecruiser: 13337,
		Deathstar:     1,
		Recycler:      5000,
	}
	b := Fleet{
		Deathstar: 1337,
		Recycler:  50000,
	}

	c := a.Add(b)

	assert.Equal(t, uint(1338), c.Deathstar)
	assert.Equal(t, uint(55000), c.Recycler)
	assert.Equal(t, uint(13337), c.Battlecruiser)
	assert.Equal(t, uint(0), c.Battleship)
}

func Test_ToResources(t *testing.T) {
	a := Fleet{
		Battlecruiser: 1000,
		Deathstar:     1000,
		Recycler:      5000,
	}

	value := a.ToResources()

	assert.Equal(t, int(5080000000), value.Metal)
	assert.Equal(t, int(4070000000), value.Crystal)
	assert.Equal(t, int(1025000000), value.Deuterium)
}
