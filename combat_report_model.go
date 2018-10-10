package ogwc

import (
	"encoding/json"
	"time"
)

type CombatReport struct {
	RepairedDefenses []RepairedDefenseEntity `json:"repaired_defenses"`
	Generic          CRGenericInformation    `json:"generic"`
	Attackers        []BattleParticipant     `json:"attackers"`
	Defenders        []BattleParticipant     `json:"defenders"`
	Rounds           []BattleRound           `json:"rounds"`
}

type BattleRound struct {
	RoundNumber        json.Number           `json:"round_number"`
	Statistics         BattleRoundStatistics `json:"statistics"`
	AttackerShips      []BattleRoundEntity   `json:"attacker_ships"`
	AttackerShipLosses []BattleRoundEntity   `json:"attacker_ship_losses"`
	DefenderShips      []BattleRoundEntity   `json:"defender_ships"`
	DefenderShipLosses []BattleRoundEntity   `json:"defender_ship_losses"`
}

type BattleRoundEntity struct {
	Owner    json.Number `json:"owner"`
	ShipType json.Number `json:"ship_type"`
	Count    json.Number `json:"count"`
}

type BattleRoundStatistics struct {
	AttackerHits         json.Number `json:"attacker_hits"`
	AttackerAbsorbed     json.Number `json:"attacker_absorbed"`
	AttackerFullstrength json.Number `json:"attacker_fullstrength"`
	DefenderHits         json.Number `json:"defender_hits"`
	DefenderAbsorbed     json.Number `json:"defender_absorbed"`
	DefenderFullstrength json.Number `json:"defender_fullstrength"`
}

type BattleParticipant struct {
	FleetOwner            string         `json:"fleet_owner"`
	FleetOwnerCoordinates string         `json:"fleet_owner_coordinates"`
	FleetOwnerPlanetType  json.Number    `json:"fleet_owner_planet_type"`
	FleetOwnerPlanetName  string         `json:"fleet_owner_planet_name"`
	FleetOwnerAlliance    string         `json:"fleet_owner_alliance"`
	FleetOwnerAllianceTag string         `json:"fleet_owner_alliance_tag"`
	FleetArmorPercentage  json.Number    `json:"fleet_armor_percentage"`
	FleetShieldPercentage json.Number    `json:"fleet_shield_percentage"`
	FleetWeaponPercentage json.Number    `json:"fleet_weapon_percentage"`
	FleetComposition      []BattleEntity `json:"fleet_composition"`
}

type BattleEntity struct {
	ShipType json.Number `json:"ship_type"`
	Armor    json.Number `json:"armor"`
	Shield   json.Number `json:"shield"`
	Weapon   json.Number `json:"weapon"`
	Count    json.Number `json:"count"`
}

type CRGenericInformation struct {
	CrID                string             `json:"cr_id"`
	EventTime           time.Time          `json:"event_time"`
	EventTimestamp      json.Number        `json:"event_timestamp"`
	CombatCoordinates   string             `json:"combat_coordinates"`
	CombatPlanetType    json.Number        `json:"combat_planet_type"`
	CombatRounds        json.Number        `json:"combat_rounds"`
	LootPercentage      json.Number        `json:"loot_percentage"`
	Winner              string             `json:"winner"`
	UnitsLostAttackers  json.Number        `json:"units_lost_attackers"`
	UnitsLostDefenders  json.Number        `json:"units_lost_defenders"`
	AttackerCount       json.Number        `json:"attacker_count"`
	DefenderCount       json.Number        `json:"defender_count"`
	LootMetal           json.Number        `json:"loot_metal"`
	LootCrystal         json.Number        `json:"loot_crystal"`
	LootDeuterium       json.Number        `json:"loot_deuterium"`
	CombatHonorable     bool               `json:"combat_honorable"`
	AttackerHonorable   bool               `json:"attacker_honorable"`
	AttackerHonorpoints json.Number        `json:"attacker_honorpoints"`
	DefenderHonorable   bool               `json:"defender_honorable"`
	DefenderHonorpoints json.Number        `json:"defender_honorpoints"`
	MoonCreated         bool               `json:"moon_created"`
	MoonChance          json.Number        `json:"moon_chance"`
	MoonSize            json.Number        `json:"moon_size"`
	MoonExists          bool               `json:"moon_exists"`
	DebrisMetal         json.Number        `json:"debris_metal"`
	DebrisCrystal       json.Number        `json:"debris_crystal"`
	Wreckfield          []WreckfieldEntity `json:"wreckfield"`
}

type WreckfieldEntity struct {
	ShipType json.Number `json:"ship_type"`
	Count    json.Number `json:"count"`
}

type RepairedDefenseEntity struct {
	RepairedType  json.Number `json:"repaired_type"`
	RepairedCount json.Number `json:"repaired_count"`
}
