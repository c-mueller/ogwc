package ogwc

import (
	"encoding/json"
	"github.com/GeertJohan/go.rice"
)

var Entities map[string]EntityType

func init() {
	cfgBucket := rice.MustFindBox("config")
	data, _ := cfgBucket.Bytes("entities.json")

	json.Unmarshal(data, &Entities)
}

type EntityType struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Cost          Resources `json:"cost"`
	CargoCapacity uint      `json:"cargo_capacity"`
}
