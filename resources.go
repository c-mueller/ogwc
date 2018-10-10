package ogwc

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
