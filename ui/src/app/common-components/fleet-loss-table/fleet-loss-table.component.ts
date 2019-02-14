// ogwc (https://github.com/c-mueller/ogwc).
// Copyright (C) 2018-2019 Christian Müller <dev@c-mueller.xyz>.
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

import {Component, Input, OnInit} from '@angular/core';
import {Fleet} from '../../svc/model';
import {fleetKeys, fleetNames, FleetNames} from '../../svc/constants';

@Component({
  selector: 'app-fleet-loss-table',
  templateUrl: './fleet-loss-table.component.html',
  styleUrls: ['./fleet-loss-table.component.css']
})
export class FleetLossTableComponent implements OnInit {

  @Input()
  public data: Map<string, Fleet>;

  public names: FleetNames = fleetNames;

  constructor() {
  }

  ngOnInit() {
  }

  getFleetValue(player: string, shipKey: string): number {
    const value = this.data[player][shipKey];

    if (value == null) {
      return 0;
    }

    return value;
  }

  getFleetKeys(): string[] {
    return fleetKeys;
  }

  getParticipants(m: Map<string, Fleet>) {
    const s: string[] = [];

    for (const k of Object.keys(m)) {
      s.push(k);
    }
    return s;
  }

}
