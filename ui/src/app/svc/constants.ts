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

import {Fleet} from './model';

export const winDistributionModes = [
  'Prozentualer Anteil',
  'Feste Rohstoffmenge',
  'Kein Gewinnanteil',
];

export const zeroFleet: Fleet = {
  'light_fighter': 0,
  'heavy_fighter': 0,
  'cruiser': 0,
  'battleship': 0,
  'battlecruiser': 0,
  'bomber': 0,
  'destroyer': 0,
  'deathstar': 0,
  'small_cargo': 0,
  'large_cargo': 0,
  'recycler': 0,
  'colonyship': 0,
  'probe': 0,
  'satellite': 0,
  'rocket_launcher': 0,
  'light_laser': 0,
  'heavy_laser': 0,
  'gauss_cannon': 0,
  'ion_cannon': 0,
  'plasma_cannon': 0,
  'small_shield': 0,
  'large_shield': 0
};

export const fleetNames: FleetNames = {
  'light_fighter': 'Leichter Jäger',
  'heavy_fighter': 'Schwerer Jäger',
  'cruiser': 'Kreuzer',
  'battleship': 'Schlachtschiff',
  'battlecruiser': 'Schlachtkreuzer',
  'bomber': 'Bomber',
  'destroyer': 'Zerstörer',
  'deathstar': 'Todesstern',
  'small_cargo': 'Kleiner Transporter',
  'large_cargo': 'Großer Transporter',
  'recycler': 'Recycler',
  'colonyship': 'Kolonieschiff',
  'probe': 'Spionagesonde',
  'satellite': 'Solarsatelit',
  'crawler':'Crawler',
  'reaper':'Reaper',
  'pathfinder':'Pathfinder',
  'rocket_launcher': 'Raketenwerfer',
  'light_laser': 'Leichtes Lasergeschütz',
  'heavy_laser': 'Schweres Lasergeschütz',
  'gauss_cannon': 'Gaußkannone',
  'ion_cannon': 'Ionenkannone',
  'plasma_cannon': 'Plasmawerfer',
  'small_shield': 'Kleine Schildkuppel',
  'large_shield': 'Große Schildkuppel'
};

export const fleetKeys: string[] = [
  'light_fighter',
  'heavy_fighter',
  'cruiser',
  'battleship',
  'battlecruiser',
  'bomber',
  'destroyer',
  'deathstar',
  'small_cargo',
  'large_cargo',
  'recycler',
  'colonyship',
  'probe',
  'satellite',
  'crawler',
  'reaper',
  'pathfinder',
  'rocket_launcher',
  'light_laser',
  'heavy_laser',
  'gauss_cannon',
  'ion_cannon',
  'plasma_cannon',
  'small_shield',
  'large_shield'
];

export interface FleetNames {
  [key: string]: string;
}
