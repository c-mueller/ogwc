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

import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {ApiService} from '../svc/api.service';
import {CalculationReport, Resources} from '../svc/model';

@Component({
  selector: 'app-report',
  templateUrl: './report.component.html',
  styleUrls: ['./report.component.css']
})
export class ReportComponent implements OnInit {

  public calculationID = '';
  public report: CalculationReport = null;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private api: ApiService
  ) {
  }

  ngOnInit() {
    this.route.paramMap.subscribe(p => {
      this.calculationID = p.get('id');
      this.getReport();
    });
  }

  getCommonMap(): any {
    const m = {
      'Gesamteinkommen': this.report.total_income,
      'Gesamtverlust': this.report.total_loss,
      'Gesamtgewinn': this.report.total_win
    };
    return m;
  }

  getReport() {
    this.api.fetchReport(this.calculationID).subscribe(data => {
      this.report = data;
    }, error1 => this.router.navigate(['/404']));
  }

  getShortLink() {
    let baseUrl = window.location.href;
    baseUrl = baseUrl.split('#')[0].replace('ui/', '');
    return baseUrl + 'r/' + this.calculationID;
  }

  navigateToCalculation() {
    this.router.navigate(['/calculation', this.calculationID]);
  }

}
