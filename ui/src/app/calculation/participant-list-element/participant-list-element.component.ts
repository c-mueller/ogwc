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

import {Component, Input, OnInit} from '@angular/core';
import {CalculationResponse, Fleet, Participant, Resources} from '../../svc/model';
import {zeroFleet} from '../../svc/constants';

@Component({
  selector: 'app-participant-list-element',
  templateUrl: './participant-list-element.component.html',
  styleUrls: ['./participant-list-element.component.css']
})
export class ParticipantListElementComponent implements OnInit {

  @Input('calculation')
  public calculation: CalculationResponse = null;

  @Input('participant')
  public participant: Participant = null;

  @Input('playerName')
  public playerName: string = null;

  public collapse = true;

  public editWinDistribution = false;

  constructor() {
  }

  ngOnInit() {
  }

  toggle() {
    this.collapse = !this.collapse;
  }

  getLostFleet(): Fleet {
    if (this.calculation.losses[this.playerName] == null) {
      return zeroFleet;
    }
    return this.calculation.losses[this.playerName];
  }

  getLoot(): Resources {
    if (this.calculation.loot_per_participant[this.playerName] == null) {
      return {
        metal: 0,
        crystal: 0,
        deuterium: 0,
      };
    }
    return this.sum(this.calculation.loot_per_participant[this.playerName]);
  }

  getHarvested(): Resources {
    if (this.calculation.harvest_reports[this.playerName] == null) {
      return {
        metal: 0,
        crystal: 0,
        deuterium: 0,
      };
    }
    return this.sum(this.calculation.harvest_reports[this.playerName]);
  }

  sum(r: Resources[]): Resources {
    let result: Resources = {
      metal: 0,
      crystal: 0,
      deuterium: 0,
    };

    for (const elem of r) {
      result = {
        metal: elem.metal + result.metal,
        crystal: elem.crystal + result.crystal,
        deuterium: elem.deuterium + result.deuterium,
      };
    }

    return result;
  }
}
