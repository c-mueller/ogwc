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

type DeuteriumBalancer interface {
	Balance(totalDemand Resources) Resources
}

type RateDeuteriumBalancer struct {
	MetalProportion     float64 `json:"metal_proportion"`
	CrystalProportion   float64 `json:"crystal_proportion"`
	MetalPerDeuterium   float64 `json:"metal_per_deuterium"`
	CrystalPerDeuterium float64 `json:"crystal_per_deuterium"`
}

func (e RateDeuteriumBalancer) Balance(totalDemand Resources) Resources {
	if totalDemand.Deuterium > 0 {
		return totalDemand
	}
	totalDeut := float64(totalDemand.Deuterium) * -1.0

	metal := e.MetalProportion * totalDeut * e.MetalPerDeuterium
	crystal := e.MetalProportion * totalDeut * e.CrystalPerDeuterium

	return Resources{
		Metal:     totalDemand.Metal + int(metal),
		Crystal:   totalDemand.Crystal + int(crystal),
		Deuterium: 0,
	}
}

type NoBalancingDeuteriumBalancer struct {
}

func (NoBalancingDeuteriumBalancer) Balance(totalDemand Resources) Resources {
	return totalDemand
}
