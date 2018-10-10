package ogwc

import "fmt"

type DistribuitonMode int

const (
	PERCENTAGE   DistribuitonMode = 0
	FIXED_AMOUNT DistribuitonMode = 1
	NONE         DistribuitonMode = 2
)

type CombatReportCalculation struct {
	Participants ParticipantList `json:"participants"`

	Attacker           bool                   `json:"attacker"`
	InitialFleet       map[string]Fleet       `json:"initial_fleet"`
	Losses             map[string]Fleet       `json:"losses"`
	TotalLoot          Resources              `json:"total_loot"`
	LootPerParticipant map[string][]Resources `json:"loot_per_participant"`
	HarvestReports     map[string][]Resources `json:"harvest_reports"`

	InitialFleetOpponent map[string]Fleet `json:"initial_fleet_opponent"`
	LossesOpponent       map[string]Fleet `json:"losses_opponent"`

	RawCombatReports  []CombatReport  `json:"raw_combat_reports"`
	RawHarvestReports []HarvestReport `json:"raw_harvest_reports"`
}

type ParticipantList []Participant

type Participant struct {
	Name                string           `json:"name"`
	DistribuitonMode    DistribuitonMode `json:"distribuiton_mode"`
	WinPercentage       float64          `json:"win_percentage"`
	FixedResourceAmount *Resources       `json:"fixed_resource_amount"`
	AdditionalLosses    *Fleet           `json:"additional_losses"`
}

func (p ParticipantList) Find(name string) *Participant {
	for _, v := range p {
		if name == v.Name {
			return &v
		}
	}
	return nil
}

func (p ParticipantList) IsPresent(name string) bool {
	for _, v := range p {
		if name == v.Name {
			return true
		}
	}
	return false
}

func (c *CombatReportCalculation) GetReport() {
	lossesPerPlayer := make(map[string]Resources)
	incomePerPlayer := make(map[string]Resources)

	totalLosses := Resources{}
	totalIncome := Resources{}

	participantsByDistibutionType := make(map[DistribuitonMode][]Participant)

	for _, participant := range c.Participants {
		loss := Resources{}

		if participant.AdditionalLosses != nil {
			loss = loss.Add(participant.AdditionalLosses.ToResources())
		}

		loss = loss.Add(c.Losses[participant.Name].ToResources())

		lossesPerPlayer[participant.Name] = loss

		income := Resources{}

		for _, v := range c.LootPerParticipant[participant.Name] {
			income = income.Add(v)
		}

		for _, v := range c.HarvestReports[participant.Name] {
			income = income.Add(v)
		}

		incomePerPlayer[participant.Name] = income

		totalIncome = totalIncome.Add(income)
		totalLosses = totalLosses.Add(loss)

		if participantsByDistibutionType[participant.DistribuitonMode] == nil {
			participantsByDistibutionType[participant.DistribuitonMode] = make([]Participant, 0)
		}
		participantsByDistibutionType[participant.DistribuitonMode] = append(participantsByDistibutionType[participant.DistribuitonMode], participant)
	}

	totalWin := totalIncome.Sub(totalLosses)

	totalWinNoFixed := totalWin.Add(Resources{})

	fmt.Println(totalWin)

	for _, v := range participantsByDistibutionType[FIXED_AMOUNT] {
		if v.FixedResourceAmount == nil {
			continue
		}

		totalWinNoFixed = totalWinNoFixed.Sub(*v.FixedResourceAmount)
	}

}

func (c *CombatReportCalculation) AddParticipant(p Participant) {
	//TODO Implement checking for valid Distribution Percentage
	if !c.Participants.IsPresent(p.Name) {
		c.Participants = append(c.Participants, p)
	}
}

func (c *CombatReportCalculation) RebalanceDistributionPercentage() {
	count := 0

	for _, v := range c.Participants {
		if v.DistribuitonMode == PERCENTAGE {
			count++
		}
	}

	for _, v := range c.Participants {
		if v.DistribuitonMode == PERCENTAGE {
			v.WinPercentage = float64(1) / float64(count)
		}
	}
}

func (c *CombatReportCalculation) AddHarvestReport(h HarvestReport) {
	m := c.HarvestReports[h.Generic.OwnerName]
	if m == nil {
		m = make([]Resources, 0)
	}
	m = append(m, h.ToResources())
	c.HarvestReports[h.Generic.OwnerName] = m

	c.RawHarvestReports = append(c.RawHarvestReports, h)

	if !c.Participants.IsPresent(h.Generic.OwnerName) {
		c.Participants = append(c.Participants, Participant{
			Name:             h.Generic.OwnerName,
			DistribuitonMode: NONE,
		})
	}
}

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
		RawCombatReports:     []CombatReport{*r},
		HarvestReports:       make(map[string][]Resources),
	}
}

func (r *CombatReport) getParticipantsAndLootDistribution(initialFleet map[string]Fleet, loot Resources) ([]Participant, map[string][]Resources) {
	count := len(initialFleet)
	defaultWinPercentage := float64(100) / float64(count)
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
