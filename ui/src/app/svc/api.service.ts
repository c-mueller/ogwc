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

import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {CalculationResponse, SubmissionResponse} from './model';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) {
  }

  addApiKey(key: string, id: string) {
    return this.http.post<SubmissionResponse>(environment.applicationEndpoint + 'api/v1/calculation/' + id + '/add/' + key , null);
  }

  submitReport(key: string) {
    return this.http.post<SubmissionResponse>(environment.applicationEndpoint + 'api/v1/submit/' + key, null);
  }

  fetchCalculation(id: string) {
    return this.http.get<CalculationResponse>(environment.applicationEndpoint + 'api/v1/calculation/' + id);
  }
}
