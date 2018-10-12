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

  constructor(private route: ActivatedRoute,
              private router: Router,
              private api: ApiService) {
  }

  ngOnInit() {
    this.route.paramMap.subscribe(p => {
      this.calculationID = p.get('id');
      this.api.fetchCalculation(this.calculationID).subscribe(e => {
        this.calculation = e;
        this.participants = this.calculation.participants;
      }, err => {
        console.log(err);
        this.router.navigate(['/404']);
      });
    });
  }

  getShortLink() {
    let baseUrl = window.location.href;
    baseUrl = baseUrl.split('#')[0].replace('ui/', '');
    return baseUrl + 'c/' + this.calculationID;
  }

}
