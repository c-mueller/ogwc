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
