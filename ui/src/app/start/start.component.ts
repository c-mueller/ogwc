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

import {Component, OnInit} from '@angular/core';
import {ApiService} from '../svc/api.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-start',
  templateUrl: './start.component.html',
  styleUrls: ['./start.component.css']
})
export class StartComponent implements OnInit {

  public currentAPIKey = '';
  public currentID = '';

  constructor(private router: Router, private api: ApiService) {
  }

  ngOnInit() {
  }

  submitAPIKey() {
    this.api.submitReport(this.currentAPIKey).subscribe((v) => {
      console.log(v.calculation_id);
      this.router.navigate(['/calculation', v.calculation_id]);
    }, error1 => {
      console.log('Submission Failed');
    });
  }

  openCalculation() {
    this.router.navigate(['/calculation', this.currentID]);
  }

}
