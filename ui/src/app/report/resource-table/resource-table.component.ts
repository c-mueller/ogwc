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
import {Resources} from '../../svc/model';

@Component({
  selector: 'app-resource-table',
  templateUrl: './resource-table.component.html',
  styleUrls: ['./resource-table.component.css']
})
export class ResourceTableComponent implements OnInit {
  @Input()
  public table: Map<string, Resources>;

  @Input('row-name')
  public rowName = 'Spieler';

  constructor() {
  }

  ngOnInit() {
  }

  getKeys(m: Map<string, Resources>) {
    const s: string[] = [];

    for (const k of Object.keys(m)) {
      console.log(k);
      s.push(k);
    }
    return s;
  }

}
