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

