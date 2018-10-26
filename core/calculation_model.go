// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (c) 2018 Christian Müller <cmueller.dev@gmail.com>.
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

type DistributionMode int

const (
	PERCENTAGE   DistributionMode = 0
	FIXED_AMOUNT DistributionMode = 1
	NONE         DistributionMode = 2
)

type FleetMap map[string]Fleet
type RessourceSliceMap map[string][]Resources

type CombatReportCalculation struct {
	Participants ParticipantList `json:"participants"`

	Attacker           bool              `json:"attacker"`
	InitialFleet       FleetMap          `json:"initial_fleet"`
	Losses             FleetMap          `json:"losses"`
	TotalLoot          Resources         `json:"total_loot"`
	LootPerParticipant RessourceSliceMap `json:"loot_per_participant"`
	HarvestReports     RessourceSliceMap `json:"harvest_reports"`

	InitialFleetOpponent FleetMap `json:"initial_fleet_opponent"`
	LossesOpponent       FleetMap `json:"losses_opponent"`

	RawCombatReports  []CombatReport  `json:"raw_combat_reports"`
	RawHarvestReports []HarvestReport `json:"raw_harvest_reports"`
}

type CalculationResponse struct {
	TotalLoss   Resources `json:"total_loss"`
	TotalIncome Resources `json:"total_income"`
	TotalWin    Resources `json:"total_win"`

	FleetLossPerParticipant FleetMap `json:"fleet_loss_per_participant"`

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
	Name                     string           `json:"name"`
	DistribuitonMode         DistributionMode `json:"distribuiton_mode"`
	WinPercentage            float64          `json:"win_percentage"`
	FixedResourceAmount      *Resources       `json:"fixed_resource_amount"`
	AdditionalLosses         *Fleet           `json:"additional_losses"`
	AdditionalResourceLosses *Resources       `json:"additional_resource_losses"`
}
