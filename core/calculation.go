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

func (c *CombatReportCalculation) AddAdditionalLossForParticipant(name string, fleet Fleet) {
	for idx, p := range c.Participants {
		if p.Name == name {
			if p.AdditionalLosses == nil {
				p.AdditionalLosses = &fleet
			} else {
				totalFleetLoss := p.AdditionalLosses.Add(fleet)
				p.AdditionalLosses = &totalFleetLoss
			}
			c.Participants[idx] = p
			return
		}
	}
}

func (c *CombatReportCalculation) GetReport() CalculationResponse {
	lossesPerPlayer := make(ResourcesMap)
	fleetLossPerPlayer := make(FleetMap)
	incomePerPlayer := make(ResourcesMap)

	lootPerPlayer := make(ResourcesMap)
	harvestedPerPlayer := make(ResourcesMap)

	winPerPlayer := make(ResourcesMap)

	claimedPerPlayer := make(ResourcesMap)
	balancePerPlayer := make(ResourcesMap)

	totalLosses := Resources{}
	totalIncome := Resources{}

	participantsByDistibutionType := make(map[DistributionMode][]Participant)

	for _, participant := range c.Participants {
		loss := Resources{}

		fleetLoss := Fleet{}

		// Add Additional fleet losses to total fleet losses (if present)
		if participant.AdditionalLosses != nil {
			loss = loss.Add(participant.AdditionalLosses.ToResources())
			fleetLoss = fleetLoss.Add(*participant.AdditionalLosses)
		}

		// Add Fleet loss ass Resources
		loss = loss.Add(c.Losses[participant.Name].ToResources())

		// Add additional Resource losses
		// e.g. flight costs
		if participant.AdditionalResourceLosses != nil {
			loss = loss.Add(*participant.AdditionalResourceLosses)
		}

		// Add Losses from Missile Reports
		for _, v := range c.MissileReports[participant.Name] {
			loss = loss.Add(v)
		}

		fleetLoss = fleetLoss.Add(c.Losses[participant.Name])

		lossesPerPlayer[participant.Name] = loss
		fleetLossPerPlayer[participant.Name] = fleetLoss

		income := Resources{}
		loot := Resources{}
		harvested := Resources{}

		for _, v := range c.LootPerParticipant[participant.Name] {
			income = income.Add(v)
			loot = loot.Add(v)
		}

		for _, v := range c.HarvestReports[participant.Name] {
			income = income.Add(v)
			harvested = harvested.Add(v)
		}

		lootPerPlayer[participant.Name] = loot
		harvestedPerPlayer[participant.Name] = harvested

		incomePerPlayer[participant.Name] = income

		totalIncome = totalIncome.Add(income)
		totalLosses = totalLosses.Add(loss)

		if participantsByDistibutionType[participant.DistribuitonMode] == nil {
			participantsByDistibutionType[participant.DistribuitonMode] = make([]Participant, 0)
		}
		participantsByDistibutionType[participant.DistribuitonMode] = append(participantsByDistibutionType[participant.DistribuitonMode], participant)
	}

	totalWin := totalIncome.Sub(totalLosses)
	totalWinUnbalanced := totalWin.Add(Resources{})

	enableRebalancer := c.DeuteriumBalancerActive && totalWin.Deuterium < 0

	// Todo find out where to put this calculation
	if enableRebalancer && false {
		totalWin = c.RebalancerConfig.Balance(totalWin)
	}

	totalWinNoFixed := totalWin.Add(Resources{})

	for _, p := range participantsByDistibutionType[FIXED_AMOUNT] {
		if p.FixedResourceAmount == nil {
			continue
		}

		totalWinNoFixed = totalWinNoFixed.Sub(*p.FixedResourceAmount)

		winPerPlayer[p.Name] = *p.FixedResourceAmount

	}

	for _, p := range participantsByDistibutionType[PERCENTAGE] {
		winPerPlayer[p.Name] = totalWinNoFixed.MulF(p.WinPercentage)
	}

	for _, p := range participantsByDistibutionType[NONE] {
		winPerPlayer[p.Name] = Resources{}
	}

	for _, p := range c.Participants {
		claimedPerPlayer[p.Name] = lossesPerPlayer[p.Name].Add(winPerPlayer[p.Name])

		balancePerPlayer[p.Name] = incomePerPlayer[p.Name].Sub(claimedPerPlayer[p.Name])
	}

	return CalculationResponse{
		TotalIncome:             totalIncome,
		TotalLoss:               totalLosses,
		TotalWin:                totalWin,
		TotalWinNoRebalancing:   totalWinUnbalanced,
		WinPerParticipant:       winPerPlayer,
		LossPerParticipant:      lossesPerPlayer,
		IncomePerParticipant:    incomePerPlayer,
		ClaimedPerParticipant:   claimedPerPlayer,
		BalancePerParticipant:   balancePerPlayer,
		HarvestedPerParticipant: harvestedPerPlayer,
		LootPerParticipant:      lootPerPlayer,
		FleetLossPerParticipant: fleetLossPerPlayer,
	}
}

func (c CalculationResponse) GetResourceTransfers() []Transfer {
	b := make(Balances, 0)

	for k, v := range c.BalancePerParticipant {
		b = append(b, PlayerBalance{
			Rank:    0,
			Name:    k,
			Balance: v,
		})
	}

	return b.GetTransferForBalancingRandomized(200)
}

func (c *CombatReportCalculation) AddParticipant(p Participant) {
	//TODO Implement checking for valid Distribution Percentage
	if !c.Participants.IsPresent(p.Name) {
		c.Participants = append(c.Participants, p)
	}
}

func (c *CombatReportCalculation) AddCombatReport(cr CombatReport, isAttacker bool) {
	getBP, getBRE, _, _ := cr.getCollectionFunctions(isAttacker)
	lossMap := cr.getFleetLosses(getBP, getBRE)
	loot := cr.getLoot()

	initialFleet := cr.getInitialFleet(getBP)
	_, lootPerPlayer := cr.getParticipantsAndLootDistribution(initialFleet, loot)

	for name, loot := range lootPerPlayer {
		losses := lossMap[name]

		initial := initialFleet[name]

		c.InitialFleet[name] = c.InitialFleet[name].Add(initial)

		c.Losses[name] = c.Losses[name].Add(losses)
		if c.LootPerParticipant[name] == nil {
			c.LootPerParticipant[name] = make([]Resources, 0)
		}
		c.LootPerParticipant[name] = append(c.LootPerParticipant[name], loot...)

		if _, p := c.Participants.Find(name); p == nil {
			c.AddParticipant(Participant{
				Name:             name,
				DistribuitonMode: NONE,
			})
		}
	}

	c.RawReports.CombatReports = append(c.RawReports.CombatReports, cr)
}

func (c *CombatReportCalculation) RebalanceDistributionPercentage() {
	count := 0

	for _, v := range c.Participants {
		if v.DistribuitonMode == PERCENTAGE {
			count++
		}
	}

	for i, v := range c.Participants {
		if v.DistribuitonMode == PERCENTAGE {
			v.WinPercentage = float64(1) / float64(count)
		}
		c.Participants[i] = v
	}
}

func (c *CombatReportCalculation) AddHarvestReport(h HarvestReport) {
	if c.HarvestReports == nil {
		c.HarvestReports = make(RessourceSliceMap)
	}


	m := c.HarvestReports[h.Generic.OwnerName]
	if m == nil {
		m = make([]Resources, 0)
	}
	m = append(m, h.ToResources())
	c.HarvestReports[h.Generic.OwnerName] = m

	c.RawReports.HarvestReports = append(c.RawReports.HarvestReports, h)

	if !c.Participants.IsPresent(h.Generic.OwnerName) {
		c.Participants = append(c.Participants, Participant{
			Name:             h.Generic.OwnerName,
			DistribuitonMode: NONE,
		})
	}
}

func (c *CombatReportCalculation) AddMissileReport(m MissileReport) {
	if c.MissileReports == nil {
		c.MissileReports = make(RessourceSliceMap)
	}

	mrs := c.MissileReports[m.Generic.AttackerName]
	if mrs == nil {
		mrs = make([]Resources, 0)
	}

	mrs = append(mrs, m.GetLosses(c.Attacker))

	c.MissileReports[m.Generic.AttackerName] = mrs

	c.RawReports.MissileReports = append(c.RawReports.MissileReports, m)

	if !c.Participants.IsPresent(m.Generic.AttackerName) {
		c.Participants = append(c.Participants, Participant{
			Name:             m.Generic.AttackerName,
			DistribuitonMode: NONE,
		})
	}
}
