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

  submitReport(key: string) {
    return this.http.post<SubmissionResponse>(environment.applicationEndpoint + 'api/v1/submit/' + key, null);
  }

  fetchCalculation(id: string) {
    return this.http.get<CalculationResponse>(environment.applicationEndpoint + 'api/v1/calculation/' + id);
  }
}
