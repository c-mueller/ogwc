package ogwc

import (
	"encoding/json"
	"time"
)

type HarvestReport struct {
	Generic HRGeneric `json:"generic"`
}

type HRGeneric struct {
	RrID                 string      `json:"rr_id"`
	EventTime            time.Time   `json:"event_time"`
	EventTimestamp       int         `json:"event_timestamp"`
	Coordinates          string      `json:"coordinates"`
	RecyclerCount        json.Number `json:"recycler_count"`
	RecyclerCapacity     json.Number `json:"recycler_capacity"`
	MetalInDebrisField   json.Number `json:"metal_in_debris_field"`
	CrystalInDebrisField json.Number `json:"crystal_in_debris_field"`
	MetalRetrieved       json.Number `json:"metal_retrieved"`
	CrystalRetrieved     json.Number `json:"crystal_retrieved"`
	OwnerName            string      `json:"owner_name"`
	OwnerAllianceName    string      `json:"owner_alliance_name"`
	OwnerAllianceTag     string      `json:"owner_alliance_tag"`
}

func (h HarvestReport) ToResources() Resources {
	m, _ := h.Generic.MetalRetrieved.Int64()
	c, _ := h.Generic.CrystalRetrieved.Int64()

	return Resources{
		Metal:     int(m),
		Crystal:   int(c),
		Deuterium: 0,
	}
}
