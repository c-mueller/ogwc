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

import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppComponent} from './app.component';
import {RouterModule, Routes} from '@angular/router';
import {StartComponent} from './start/start.component';
import {NotFoundComponent} from './not-found/not-found.component';
import {CalculationComponent} from './calculation/calculation.component';
import {ReportComponent} from './report/report.component';
import {AboutComponent} from './about/about.component';
import {NgbAlertModule, NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {CommonModule, HashLocationStrategy, LocationStrategy} from '@angular/common';
import {FormsModule} from '@angular/forms';
import {ApiService} from './svc/api.service';
import {HttpClientModule} from '@angular/common/http';
import {ParticipantListElementComponent} from './calculation/participant-list-element/participant-list-element.component';
import {ResourceViewComponent} from './common-components/resource-view/resource-view.component';
import {FleetViewComponent} from './common-components/fleet-view/fleet-view.component';
import {ResourceTableComponent} from './common-components/resource-table/resource-table.component';
import {ClipboardModule} from 'ngx-clipboard';
import { FleetLossTableComponent } from './common-components/fleet-loss-table/fleet-loss-table.component';
import { TransfersTableComponent } from './report/transfers-table/transfers-table.component';

const appRoutes: Routes = [
  {path: 'start', component: StartComponent},
  {path: 'about', component: AboutComponent},
  {path: 'calculation/:id', component: CalculationComponent},
  {path: 'calculation/:id/report', component: ReportComponent},
  {path: '', redirectTo: '/start', pathMatch: 'full'},
  {path: '**', component: NotFoundComponent},
  {path: '404', component: NotFoundComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    StartComponent,
    NotFoundComponent,
    CalculationComponent,
    ReportComponent,
    AboutComponent,
    ParticipantListElementComponent,
    ResourceViewComponent,
    FleetViewComponent,
    ResourceTableComponent,
    FleetLossTableComponent,
    TransfersTableComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    NgbModule,
    NgbAlertModule,
    FormsModule,
    HttpClientModule,
    CommonModule,
    ClipboardModule,
  ],
  providers: [
    {
      provide: LocationStrategy,
      useClass: HashLocationStrategy
    },
    ApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
