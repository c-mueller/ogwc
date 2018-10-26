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

import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {Fleet} from '../../svc/model';
import {fleetKeys, fleetNames, FleetNames} from '../../svc/constants';

@Component({
  selector: 'app-fleet-view',
  templateUrl: './fleet-view.component.html',
  styleUrls: ['./fleet-view.component.css']
})
export class FleetViewComponent implements OnInit {

  @Input('heading')
  public heading: string;

  @Input('fleet')
  public fleet: Fleet;

  @Input('edit')
  public edit = false;

  @Output('onedit')
  public editEmitter: EventEmitter<Fleet> = new EventEmitter<Fleet>();

  public names: FleetNames = fleetNames;

  public editActive = false;

  public editFleet: Fleet;

  getKeys(): string[] {
    return fleetKeys;
  }

  onFleetUpdate() {
    this.editActive = false;
    this.editEmitter.emit(this.editFleet);
  }

  constructor() {
  }

  ngOnInit() {
    this.editFleet = this.fleet;
  }

}

