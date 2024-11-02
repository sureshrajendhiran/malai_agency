import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MainQiComponent } from './main-qi/main-qi.component';
import { CreateQiComponent } from './create-qi/create-qi.component';
import { QIRoutingModule } from './qi-routing';
import { MaterialModule } from '../material-module';


@NgModule({
  declarations: [
    MainQiComponent,
    CreateQiComponent
  ],
  imports: [
    CommonModule,
    QIRoutingModule,
    MaterialModule
  ]
})
export class QiModule { }
