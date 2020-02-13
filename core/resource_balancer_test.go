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
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Balancer(t *testing.T) {
	scenario := Balances{
		{
			Name: "P1",
			Rank: 1,
			Balance: Resources{
				Metal:   -500000000,
				Crystal: -500000000,
			},
		},
		{
			Name: "P2",
			Rank: 10,
			Balance: Resources{
				Metal:   1000000000,
				Crystal: 1000000000,
			},
		},
		{
			Name: "P3",
			Rank: 20,
			Balance: Resources{
				Metal:   -500000000,
				Crystal: -500000000,
			},
		},
	}

	transfers := scenario.GetTransferForBalancing()

	assert.True(t, len(transfers) == 2)
}

func TestBalances_Complex_1(t *testing.T) {
	scenario := Balances{
		{
			Name: "Jameson_4K",
			Rank: 5,
			Balance: Resources{
				Metal:   -3041232309,
				Crystal: -2218891447,
				Deuterium: 41506332,
			},
		},
		{
			Name: "Chris",
			Rank: 10,
			Balance: Resources{
				Metal:   2401730793,
				Crystal: 399041251,
				Deuterium: -13404501,
			},
		},
		{
			Name: "Who",
			Rank: 11,
			Balance: Resources{
				Metal:   1139501517,
				Crystal: 2319850198,
				Deuterium: -28101833,
			},
		},
		{
			Name: "AkSent",
			Rank: 12,
			Balance: Resources{
				Metal:   -500000000,
				Crystal: -500000000,
			},
		},
	}

	tr := scenario.GetTransferForBalancing()
	for _, transfer := range tr {
		fmt.Printf("From %q to %q\nMetal: %d\nCrystal: %d\nDeuterium: %d\n",
			transfer.From, transfer.To, transfer.Amount.Metal, transfer.Amount.Crystal, transfer.Amount.Deuterium)
	}
}
func TestBalances_Complex_2(t *testing.T) {
	scenario := Balances{
		{
			Name: "Jameson_4K",
			Rank: 50,
			Balance: Resources{
				Metal:   -3041232309,
				Crystal: -2218891447,
				Deuterium: 41506332,
			},
		},
		{
			Name: "Chris",
			Rank: 10,
			Balance: Resources{
				Metal:   2401730793,
				Crystal: 399041251,
				Deuterium: -13404501,
			},
		},
		{
			Name: "Who",
			Rank: 11,
			Balance: Resources{
				Metal:   1139501517,
				Crystal: 2319850198,
				Deuterium: -28101833,
			},
		},
		{
			Name: "AkSent",
			Rank: 12,
			Balance: Resources{
				Metal:   -500000000,
				Crystal: -500000000,
			},
		},
	}

	tr := scenario.GetTransferForBalancingRandomized(250)
	for _, transfer := range tr {
		fmt.Printf("From %q to %q\nMetal: %d\nCrystal: %d\nDeuterium: %d\n",
			transfer.From, transfer.To, transfer.Amount.Metal, transfer.Amount.Crystal, transfer.Amount.Deuterium)
	}
}
