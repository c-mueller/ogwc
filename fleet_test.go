package ogwc

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

	assert.Equal(t, uint(5080000000), value.Metal)
	assert.Equal(t, uint(4070000000), value.Crystal)
	assert.Equal(t, uint(1025000000), value.Deuterium)
}
