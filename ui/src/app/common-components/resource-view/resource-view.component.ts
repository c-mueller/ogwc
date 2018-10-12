import {Component, Input, OnInit} from '@angular/core';
import {Resources} from '../../svc/model';

@Component({
  selector: 'app-resource-view',
  templateUrl: './resource-view.component.html',
  styleUrls: ['./resource-view.component.css']
})
export class ResourceViewComponent implements OnInit {

  @Input('resource')
  public resource: Resources;

  @Input('heading')
  public heading: string = null;

  constructor() {
  }

  ngOnInit() {
  }

}
