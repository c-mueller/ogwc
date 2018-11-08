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

export interface VersionInfo {
  build_context: string;
  build_timestamp: string;
  build_revision: string;
  version: string;
  frontend_hashsum: string;
}

export interface CalculationReport {
  total_loss: Resources;
  total_income: Resources;
  total_win: Resources;

  loss_per_participant: Map<string, Resources>;
  harvested_per_participant: Map<string, Resources>;
  loot_per_participant: Map<string, Resources>;
  income_per_participant: Map<string, Resources>;
  win_per_participant: Map<string, Resources>;
  balance_per_participant: Map<string, Resources>;
  claimed_per_participant: Map<string, Resources>;
  fleet_loss_per_participant: Map<string, Fleet>;
}

export interface ResourceTransfer {
  from: string;
  to: string;
  amount: Resources;
}

export interface SubmissionResponse {
  code: number;
  calculation_id: string;
  calculation_url: string;
}

export interface Resources {
  metal: number;
  crystal: number;
  deuterium: number;
}

export interface Fleet {
  [key: string]: number;
}

export interface FleetMap {
  [key: string]: Fleet;
}

export interface ResourcesMap {
  [key: string]: Resources;
}

export interface ResourcesArrayMap {
  [key: string]: Resources[];
}

export interface CalculationResponse {
  participants: Participant[];
  attacker: boolean;
  initial_fleet: FleetMap;
  losses: FleetMap;
  total_loot: Resources;
  loot_per_participant: ResourcesArrayMap;
  harvest_reports: ResourcesArrayMap;
  initial_fleet_opponent: FleetMap;
  losses_opponent: FleetMap;
}

export interface Participant {
  name: string;
  distribuiton_mode: number;
  win_percentage: number;
  fixed_resource_amount?: Resources;
  additional_losses?: Fleet;
}

