package core

import (
	"fmt"
	"testing"
)

func TestRebalancer(t *testing.T) {
	rb := RateDeuteriumBalancer{
		MetalProportion:     0.5,
		CrystalProportion:   0.5,
		MetalPerDeuterium:   2.5,
		CrystalPerDeuterium: 1.5,
	}

	res := rb.Balance(Resources{0, 0, -150000000})

	fmt.Println(res)
}
