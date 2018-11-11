package core

import (
	"encoding/json"
	"time"
)

type MissileReport struct {
	Generic MissileReportGeneric `json:"generic"`
	Details MissileReportDetails `json:"details"`
}

type MissileReportDetails struct {
	Defense          []interface{} `json:"defense"`
	DefenseDestroyed []interface{} `json:"defense_destroyed"`
}

type MissileReportGeneric struct {
	MrID                      string      `json:"mr_id"`
	EventTime                 time.Time   `json:"event_time"`
	EventTimestamp            json.Number `json:"event_timestamp"`
	AttackerName              string      `json:"attacker_name"`
	AttackerPlanetName        string      `json:"attacker_planet_name"`
	AttackerPlanetCoordinates string      `json:"attacker_planet_coordinates"`
	AttackerPlanetType        json.Number `json:"attacker_planet_type"`
	DefenderName              string      `json:"defender_name"`
	DefenderPlanetName        string      `json:"defender_planet_name"`
	DefenderPlanetCoordinates string      `json:"defender_planet_coordinates"`
	DefenderPlanetType        json.Number `json:"defender_planet_type"`
	MissilesLostAttacker      json.Number `json:"missiles_lost_attacker"`
	MissilesLostDefender      json.Number `json:"missiles_lost_defender"`
}

func (r MissileReport) GetLosses(attacker bool) Resources {
	if attacker {
		v, _ := r.Generic.MissilesLostAttacker.Int64()
		return IPMCost.Mul(int(v))
	} else {
		//TODO Implement Computation of cost for lost defenses/ AntiIPMs
		return Resources{}
	}
}