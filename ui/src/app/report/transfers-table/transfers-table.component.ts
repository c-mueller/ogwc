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

import {Component, Input, OnInit} from '@angular/core';
import {ResourceTransfer} from '../../svc/model';
import {ApiService} from '../../svc/api.service';

@Component({
  selector: 'app-transfers-table',
  templateUrl: './transfers-table.component.html',
  styleUrls: ['./transfers-table.component.css']
})
export class TransfersTableComponent implements OnInit {

  @Input('calulationID')
  public calcID: string;

  public transfers: ResourceTransfer[];

  public visible = false;

  constructor(public api: ApiService) {
  }

  ngOnInit() {
    this.api.fetchTransfers(this.calcID).subscribe(e => {
      this.transfers = e;
      this.visible = true;
    });
  }

}
