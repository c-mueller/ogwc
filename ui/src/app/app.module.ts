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
    FleetViewComponent
  ],
  imports: [
    BrowserModule,
    RouterModule.forRoot(appRoutes),
    NgbModule,
    NgbAlertModule,
    FormsModule,
    HttpClientModule,
    CommonModule,
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
