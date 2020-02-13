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

func (r *CombatReport) ToReportCalculation(isAttacker bool) *CombatReportCalculation {
	getBP, getBRE, getBPOpponent, getBPEOpponent := r.getCollectionFunctions(isAttacker)

	loot := r.getLoot()

	initialFleet := r.getInitialFleet(getBP)
	initialOpponentFleet := r.getInitialFleet(getBPOpponent)

	lossMap := r.getFleetLosses(getBP, getBRE)
	opponentLossMap := r.getFleetLosses(getBPOpponent, getBPEOpponent)

	participants, lootPerPlayer := r.getParticipantsAndLootDistribution(initialFleet, loot)

	return &CombatReportCalculation{
		Participants:         participants,
		Attacker:             isAttacker,
		InitialFleet:         initialFleet,
		InitialFleetOpponent: initialOpponentFleet,
		LootPerParticipant:   lootPerPlayer,
		TotalLoot:            loot,
		Losses:               lossMap,
		LossesOpponent:       opponentLossMap,
		HarvestReports:       make(map[string][]Resources),
		RawReports: RawReports{
			CombatReports: []CombatReport{*r},
		},
	}
}

func (r *CombatReport) getParticipantsAndLootDistribution(initialFleet map[string]Fleet, loot Resources) ([]Participant, map[string][]Resources) {
	count := len(initialFleet)
	defaultWinPercentage := float64(1) / float64(count)
	participants := make([]Participant, 0)
	lootPerPlayer := make(map[string][]Resources)
	totalLoot := 0
	for name, fleet := range initialFleet {
		participants = append(participants, Participant{
			Name:             name,
			WinPercentage:    defaultWinPercentage,
			DistribuitonMode: PERCENTAGE,
		})

		lootPerPlayer[name] = make([]Resources, 0)

		totalLoot += int(fleet.GetCargoCapacity())
	}
	for name, fleet := range initialFleet {
		lootPercentage := float64(fleet. /*.Sub(lossMap[name])*/ GetCargoCapacity()) / float64(totalLoot)
		lootPerPlayer[name] = append(lootPerPlayer[name], loot.MulF(lootPercentage))
	}
	return participants, lootPerPlayer
}

func (r *CombatReport) getLoot() Resources {
	m, _ := r.Generic.LootMetal.Int64()
	c, _ := r.Generic.LootCrystal.Int64()
	d, _ := r.Generic.LootDeuterium.Int64()
	loot := Resources{
		Metal:     int(m),
		Crystal:   int(c),
		Deuterium: int(d),
	}
	return loot
}

func (r *CombatReport) getCollectionFunctions(isAttacker bool) (getBP, getBRE, getBP, getBRE) {
	getBP := func() []BattleParticipant {
		if isAttacker {
			return r.Attackers
		}
		return r.Defenders
	}
	getBRE := func(round BattleRound) []BattleRoundEntity {
		if isAttacker {
			return round.AttackerShipLosses
		}
		return round.DefenderShipLosses
	}
	getBPOpponent := func() []BattleParticipant {
		if !isAttacker {
			return r.Attackers
		}
		return r.Defenders
	}
	getBPEOpponent := func(round BattleRound) []BattleRoundEntity {
		if !isAttacker {
			return round.AttackerShipLosses
		}
		return round.DefenderShipLosses
	}
	return getBP, getBRE, getBPOpponent, getBPEOpponent
}

func (r *CombatReport) getInitialFleet(getBP getBP) map[string]Fleet {
	initialFleet := make(map[string]Fleet)
	for _, fleetPart := range getBP() {
		counts := make(map[uint]uint)

		for _, element := range fleetPart.FleetComposition {
			t, _ := element.ShipType.Int64()
			c, _ := element.Count.Int64()

			counts[uint(t)] = uint(c)
		}

		initialFleet[fleetPart.FleetOwner] = initialFleet[fleetPart.FleetOwner].Add(toFleet(counts))
	}

	return initialFleet
}

type getBP func() []BattleParticipant
type getBRE func(round BattleRound) []BattleRoundEntity

func (r *CombatReport) getFleetLosses(getPartyAsBattleParticipants getBP, getPartyAsBattleParticipantsForRound getBRE) map[string]Fleet {
	ids := r.getIndexToPlayerNameMap(getPartyAsBattleParticipants)
	shipLosses := r.countShipLossesForParty(getPartyAsBattleParticipantsForRound, ids)
	lossMap := make(map[string]Fleet)
	for player, losses := range shipLosses {
		l := toFleet(losses)

		lossMap[player] = l
	}

	return lossMap
}

func (r *CombatReport) countShipLossesForParty(getPartyAsBattleParticipantsForRound getBRE, ids map[uint]string) map[string]map[uint]uint {
	shipLosses := make(map[string]map[uint]uint)
	for _, round := range r.Rounds {
		for _, losses := range getPartyAsBattleParticipantsForRound(round) {
			o, _ := losses.Owner.Int64()
			owner := ids[uint(o)]
			lossesForOwner := shipLosses[owner]
			if lossesForOwner == nil {
				lossesForOwner = make(map[uint]uint)
			}

			i, _ := losses.Count.Int64()

			shipID, _ := losses.ShipType.Int64()

			lossesForOwner[uint(shipID)] += uint(i)

			shipLosses[owner] = lossesForOwner
		}
	}
	return shipLosses
}

func (r *CombatReport) getIndexToPlayerNameMap(getPartyAsBattleParticipants getBP) map[uint]string {
	ids := make(map[uint]string)

	for index, attackerParty := range getPartyAsBattleParticipants() {
		ids[uint(index)] = attackerParty.FleetOwner
	}

	return ids
}
