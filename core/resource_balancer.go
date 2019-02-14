// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
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
	"math/rand"
	"sort"
	"time"
)

type PlayerBalance struct {
	Name    string
	Rank    int
	Balance Resources
}

type Transfer struct {
	From   string    `json:"from"`
	To     string    `json:"to"`
	Amount Resources `json:"amount"`
}

type Balances []PlayerBalance
type Transfers []Transfer

type numberMap map[string]int

type toInt func(Resources) int
type toRes func(int) Resources

type functionPair struct {
	ToInteger   toInt
	ToResources toRes
}

var functions = []functionPair{
	{
		ToInteger: func(resources Resources) int {
			return resources.Metal
		},
		ToResources: func(i int) Resources {
			return Resources{
				Metal: i,
			}
		},
	},
	{
		ToInteger: func(resources Resources) int {
			return resources.Crystal
		},
		ToResources: func(i int) Resources {
			return Resources{
				Crystal: i,
			}
		},
	},
	{
		ToInteger: func(resources Resources) int {
			return resources.Deuterium
		},
		ToResources: func(i int) Resources {
			return Resources{
				Deuterium: i,
			}
		},
	},
}

func (b Balances) GetTransferForBalancingRandomized(cycles int) []Transfer {
	if !b.validate() {
		return nil
	}

	bestLen := int(^uint(0) >> 1)
	bestResult := make([]Transfer, 0)

	for i := 0; i < cycles; i++ {
		transfers := b.getTransfersForBalancing(func(ufMap numberMap) []string {
			keys := make([]string, 0)
			for k, _ := range ufMap {
				keys = append(keys, k)
			}
			return shuffleStringArray(keys)
		})

		if len(transfers) < bestLen {
			bestLen = len(transfers)
			bestResult = transfers
		}
	}

	return bestResult
}

func (b Balances) GetTransferForBalancing() []Transfer {
	if !b.validate() {
		return nil
	}

	return b.getTransfersForBalancing(func(ufMap numberMap) []string {
		return ufMap.GetKeysOrderedByRank(b)
	})
}

func (b Balances) getTransfersForBalancing(undeflowOrderFunc func(ufMap numberMap) []string) []Transfer {
	transfers := make(Transfers, 0)

	// The following operations get executed for metal, crystal and deuterium
	// seperatly
	for _, functions := range functions {
		// Contains Players with too many or exacty enough of this type
		overflow := make(numberMap)
		// Contains Players with
		underflow := make(numberMap)

		// Sorts all balances into two "buckets"
		// the Overflow Bucket (a player having too much of a resource)
		// and a underflow bucket (a player that needs resources to get a balance of 0)
		// If the balance value for a resource is zero, the player will be ignored in this resource,
		// because the balance is fulfiled anyways
		for _, balance := range b {
			value := functions.ToInteger(balance.Balance)
			if value < 0 {
				underflow[balance.Name] = value
			} else if value > 0 {
				overflow[balance.Name] = value
			}
		}

		// Iterate over all overflowing players
		// ordered ascendingly by their rank
		// The intention here is to minimize the amount of transfers requiring a Ticket
		for _, name := range overflow.GetKeysOrderedByRank(b) {
			// When iterating over the underflow map we use a parameter defined function
			// to determine the order.
			// The default implementation shuffles the keys.
			for _, uName := range undeflowOrderFunc(underflow) {
				// Get the Current Overflow Value for the overflowing player
				// retrieving this within the inner loop is neccessary, to update the value
				// if it gets partialy used (eg. iteration 1 subtracts 1000 from a total overflow of 2000 -> remaining 1000)
				overflowAmt := overflow[name]

				// get the current underflow value for the underflowing player
				underflowAmt := underflow[uName]
				remaining := overflowAmt + underflowAmt

				// If the remaining value is positive, the overflowing player has
				// more resources of this type than, the amount that is needed to balance the
				// underflowing player
				if remaining >= 0 {
					transfers = append(transfers, Transfer{
						From:   name,
						To:     uName,
						Amount: functions.ToResources(underflowAmt * -1),
					})

					// Delete underflowing player, because the resouce balance is 0
					delete(underflow, uName)

					// If the remainder is Zero (with tolerance)
					// also delete the overflowing value, because its balance is also zero
					if remaining <= 100 && remaining >= -100 {
						delete(overflow, name)
					} else { // Otherwise update the overflowing value
						overflow[name] = remaining
					}
				} else if remaining < 0 {
					// if the remaining value is less than zero, the balance of the underflowing player
					// cannot be fulfiled by this overflowing player alone, another player has to transfer
					// resources to this player
					transfers = append(transfers, Transfer{
						From:   name,
						To:     uName,
						Amount: functions.ToResources(overflowAmt),
					})

					// Remove overflowing value because the balance is zero
					delete(overflow, name)

					// update underflowing value
					underflow[uName] = remaining
					break
				}

			}
		}
	}

	return transfers.Normalize()
}

// Normalizes the Transfers array
// This means that all transfers with the same
// source and target (aka. from and to) will be added together
func (t Transfers) Normalize() Transfers {
	transferMap := make(map[string]map[string]*Resources)

	for _, transfer := range t {
		if transferMap[transfer.From] == nil {
			transferMap[transfer.From] = make(map[string]*Resources)
		}

		res := Resources{}

		if transferMap[transfer.From][transfer.To] != nil {
			res = *transferMap[transfer.From][transfer.To]
		}

		res = res.Add(transfer.Amount)

		transferMap[transfer.From][transfer.To] = &res
	}

	transfers := make(Transfers, 0)

	for from, innerMap := range transferMap {

		for to, amount := range innerMap {
			transfers = append(transfers, Transfer{
				From:   from,
				To:     to,
				Amount: *amount,
			})
		}
	}

	return transfers
}

// Checks if the sum of all balances is zero (with small tolerance)
func (b Balances) validate() bool {
	r := Resources{}

	for _, v := range b {
		r = r.Add(v.Balance)
	}

	return r.IsZeroWithTolerance()
}

func (o numberMap) GetKeysOrderedByRank(b Balances) []string {
	order := make([]string, 0)

	p := make(rnkPairs, 0)
	for k, _ := range o {
		p = append(p, sortPair{k, b.GetRank(k)})
	}

	sort.Sort(p)

	for _, v := range p {
		order = append(order, v.Key)
	}

	return order
}

// Returns the Keys of the map odered by the value
func (o numberMap) GetKeysOrderedByValue() []string {
	order := make([]string, 0)

	p := make(ofPairs, 0)
	for k, v := range o {
		p = append(p, sortPair{k, v})
	}

	sort.Sort(p)

	for _, v := range p {
		order = append(order, v.Key)
	}

	return order
}

type sortPair struct {
	Key   string
	Value int
}

func shuffleStringArray(vals []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

type ofPairs []sortPair

func (p ofPairs) Len() int           { return len(p) }
func (p ofPairs) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p ofPairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type rnkPairs []sortPair

func (p rnkPairs) Len() int           { return len(p) }
func (p rnkPairs) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p rnkPairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (b Balances) GetRank(name string) int {
	for _, v := range b {
		if v.Name == name {
			return v.Rank
		}
	}
	return 0
}
