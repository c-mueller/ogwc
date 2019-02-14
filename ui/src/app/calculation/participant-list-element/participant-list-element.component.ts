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

import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {CalculationResponse, Fleet, Participant, Resources} from '../../svc/model';
import {fleetKeys, winDistributionModes, zeroFleet} from '../../svc/constants';
import {ApiService} from '../../svc/api.service';

@Component({
  selector: 'app-participant-list-element',
  templateUrl: './participant-list-element.component.html',
  styleUrls: ['./participant-list-element.component.css']
})
export class ParticipantListElementComponent implements OnInit {

  @Output('updated')
  public updatePerformed: EventEmitter<any> = new EventEmitter();

  @Input('calculation')
  public calculation: CalculationResponse = null;

  @Input('calculationID')
  public calculationID: string;

  @Input('participant')
  public participant: Participant = null;

  @Input('playerName')
  public playerName: string = null;

  public collapse = true;

  public editWinDistribution = false;
  public winDistributionMode = 0;

  public fixedResourcesDistributionModeAmount: Resources;

  public winPercentage = 0;

  public distributionModeNames = winDistributionModes;

  public deletionClickCount = 0;
  public deletionTexts = [
    'Diesen Teilnehmer löschen',
    'Zum löschen des Teilnehmers erneut Klicken.'
  ];

  constructor(private api: ApiService) {
  }

  deleteParticipant() {
    this.deletionClickCount = (this.deletionClickCount + 1) % this.deletionTexts.length;

    if (this.deletionClickCount === 0) {
      this.api.deleteParticipant(this.calculationID, this.playerName).subscribe(e => {
        this.updatePerformed.emit(null);
        alert('Teilnehmer gelöscht!');
      }, e => {
        console.log(e);
        alert('Teilnehmer konnte nicht gelöscht werden!');
      });
    }
  }

  selectDistributionMode(mode: number) {
    this.winDistributionMode = mode;
  }

  getAdditionalFleetLoss() {
    if (this.participant.additional_losses == null) {
      const f: Fleet = {};

      for (const key of fleetKeys) {
        f[key] = 0;
      }

      return f;
    }
    return this.participant.additional_losses;
  }

  ngOnInit() {
    this.winDistributionMode = this.participant.distribuiton_mode;
    if (this.participant.fixed_resource_amount != null) {
      this.fixedResourcesDistributionModeAmount = this.participant.fixed_resource_amount;
    } else {
      this.fixedResourcesDistributionModeAmount = {
        metal: 0,
        crystal: 0,
        deuterium: 0,
      };
    }

    this.winPercentage = this.participant.win_percentage * 100;
  }

  saveDistributionMode() {
    if (this.winDistributionMode === 0) {
      this.api.updateWinDistributionModeToPercentage(this.calculationID, this.playerName, this.winPercentage / 100).subscribe(e => {
        this.onWinUpdateSuccess();
      }, error1 => {
        this.onWinUpdateFail();
      });
    } else if (this.winDistributionMode === 1) {
      this.api.updateWinDistributionModeToFixedAmount(this.calculationID, this.playerName,
        this.fixedResourcesDistributionModeAmount)
        .subscribe(e => {
          this.onWinUpdateSuccess();
        }, error1 => {
          this.onWinUpdateFail();
        });
    } else {
      this.api.updateWinDistributionModeTonone(this.calculationID, this.playerName).subscribe(e => {
        this.onWinUpdateSuccess();
      }, error1 => {
        this.onWinUpdateFail();
      });
    }
  }

  updateAdditionalLoss(f: Fleet) {
    this.participant.additional_losses = f;
    this.api.updateAdditionalFleetLoss(this.calculationID, this.playerName, f).subscribe(e => {
      alert('Flottenverluste Aktualisiert');
    }, e => {
      alert('Fehler beim Aktuallisieren der Flottenverluste');
    });
  }

  onWinUpdateSuccess() {
    this.editWinDistribution = false;
    alert('Gewinnverteilung angepasst');
    this.updatePerformed.emit(null);
  }

  onWinUpdateFail() {
    this.editWinDistribution = false;
    alert('Gewinnverteilung konnte nicht angepasst werden!');
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
