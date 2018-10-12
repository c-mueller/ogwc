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
