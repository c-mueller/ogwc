// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package core

type DistributionMode int

const (
	PERCENTAGE   DistributionMode = 0
	FIXED_AMOUNT DistributionMode = 1
	NONE         DistributionMode = 2
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

type CalculationResponse struct {
	TotalLoss   Resources `json:"total_loss"`
	TotalIncome Resources `json:"total_income"`
	TotalWin    Resources `json:"total_win"`

	LossPerParticipant      ResourcesMap `json:"loss_per_participant"`
	HarvestedPerParticipant ResourcesMap `json:"harvested_per_participant"`
	LootPerParticipant      ResourcesMap `json:"loot_per_participant"`
	IncomePerParticipant    ResourcesMap `json:"income_per_participant"`
	WinPerParticipant       ResourcesMap `json:"win_per_participant"`

	BalancePerParticipant ResourcesMap `json:"balance_per_participant"`
	ClaimedPerParticipant ResourcesMap `json:"claimed_per_participant"`
}

type ResourcesMap map[string]Resources
type ParticipantList []Participant

type Participant struct {
	Name                string           `json:"name"`
	DistribuitonMode    DistributionMode `json:"distribuiton_mode"`
	WinPercentage       float64          `json:"win_percentage"`
	FixedResourceAmount *Resources       `json:"fixed_resource_amount"`
	AdditionalLosses    *Fleet           `json:"additional_losses"`
}
