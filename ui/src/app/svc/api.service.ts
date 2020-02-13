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

import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {CalculationReport, CalculationResponse, Fleet, Resources, ResourceTransfer, SubmissionResponse, VersionInfo} from './model';
import {environment} from '../../environments/environment';
import {resource} from 'selenium-webdriver/http';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) {
  }

  updateAdditionalFleetLoss(calcID: string, name: string, fleet: Fleet) {
    return this.http.post(environment.applicationEndpoint +
      'api/v1/calculation/' + calcID + '/participant/fleet-loss', {
      'name': name,
      'lost_fleet': fleet
    }, {});
  }

  addParticipant(id: string, name: string) {
    return this.http.post(environment.applicationEndpoint +
      'api/v1/calculation/' + id + '/participant/add?name=' + name, null);
  }

  deleteParticipant(id: string, name: string) {
    return this.http.post(environment.applicationEndpoint +
      'api/v1/calculation/' + id + '/participant/delete', null, {
      params: {
        'name': name,
      }
    });
  }

  updateWinDistributionModeToPercentage(id: string, name: string, percentage: number) {
    return this.http.post(environment.applicationEndpoint + 'api/v1/calculation/' + id + '/participant/win/percentage', null, {
      params: {
        'name': name,
        'percentage': percentage.toString(10)
      }
    });
  }

  updateWinDistributionModeToFixedAmount(id: string, name: string, amount: Resources) {
    return this.http.post(environment.applicationEndpoint +
      'api/v1/calculation/' + id + '/participant/win/fixed', amount, {
      params: {
        'name': name
      }
    });
  }

  updateWinDistributionModeTonone(id: string, name: string) {
    return this.http.post(environment.applicationEndpoint +
      'api/v1/calculation/' + id + '/participant/win/none', null, {
      params: {
        'name': name
      }
    });
  }

  addApiKey(key: string, id: string) {
    return this.http.post<SubmissionResponse>(environment.applicationEndpoint + 'api/v1/calculation/' + id + '/add/' + key, null);
  }

  submitReport(key: string) {
    return this.http.post<SubmissionResponse>(environment.applicationEndpoint + 'api/v1/submit/' + key, null);
  }

  fetchCalculation(id: string) {
    return this.http.get<CalculationResponse>(environment.applicationEndpoint + 'api/v1/calculation/' + id);
  }

  fetchReport(id: string) {
    return this.http.get<CalculationReport>(environment.applicationEndpoint + 'api/v1/calculation/' + id + '/report');
  }

  fetchTransfers(id: string) {
    return this.http.get<ResourceTransfer[]>(environment.applicationEndpoint + 'api/v1/calculation/' + id + '/report/transfers');
  }
  fetchVersionInfo() {
    return this.http.get<VersionInfo>(environment.applicationEndpoint + 'api/v1/version');
  }
}
