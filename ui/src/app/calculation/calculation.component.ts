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
import {CalculationResponse, Participant} from '../svc/model';

@Component({
  selector: 'app-calculation',
  templateUrl: './calculation.component.html',
  styleUrls: ['./calculation.component.css']
})
export class CalculationComponent implements OnInit {

  public calculationID = '';
  public calculation: CalculationResponse = null;
  public participants: Participant[] = [];

  public apiKey: string;
  public newParticipantName: string;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private api: ApiService) {
  }

  loadReport() {
    this.router.navigate(['/calculation', this.calculationID, 'report']);
  }

  ngOnInit() {
    this.route.paramMap.subscribe(p => {
      this.calculationID = p.get('id');
      this.getCalculation();
    });
  }

  addParticipant() {
    if (this.newParticipantName.length > 0) {
      this.api.addParticipant(this.calculationID, this.newParticipantName).subscribe(e => {
        this.getCalculation();
        this.newParticipantName = '';
        alert('Teilnehmer wurde Hinzugefügt!');
      }, err => {
        console.log(err);
        alert('Der Teilnehmer konnte nicht angelegt werden.');
      });
    }
  }

  getCalculation() {
    this.api.fetchCalculation(this.calculationID).subscribe(e => {
      this.calculation = e;
      this.participants = this.calculation.participants;
    }, err => {
      console.log(err);
      this.router.navigate(['/404']);
    });
  }

  requestCalculationUpdate() {
    if (this.calculation != null) {
      this.getCalculation();
    }
  }

  addAPIKey() {
    this.api.addApiKey(this.apiKey, this.calculationID).subscribe(e => {
      this.getCalculation();
      alert('API Key wurde Hinzugefügt');
    }, err => {
      console.log(err);
      alert('API Key wurde nicht Hinzugefügt. Es ist ein Fehler aufgetreten.');
    });
    this.apiKey = '';
  }

  getShortLink() {
    let baseUrl = window.location.href;
    baseUrl = baseUrl.split('#')[0].replace('ui/', '');
    return baseUrl + 'c/' + this.calculationID;
  }

}
