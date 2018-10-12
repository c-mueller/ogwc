import {Component, Input, OnInit} from '@angular/core';
import {Fleet} from '../../svc/model';
import {fleetKeys, fleetNames, FleetNames} from '../../svc/constants';

@Component({
  selector: 'app-fleet-view',
  templateUrl: './fleet-view.component.html',
  styleUrls: ['./fleet-view.component.css']
})
export class FleetViewComponent implements OnInit {

  @Input('fleet')
  public fleet: Fleet;

  public names: FleetNames = fleetNames;

  getKeys(): string[] {
    return fleetKeys;
  }

  constructor() {
  }

  ngOnInit() {
  }

}

