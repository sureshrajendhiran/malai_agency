import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MasterMainComponent } from './master-main/master-main.component';
import { MasterModuleRoutingModule } from './master-routing';
import { MaterialModule } from '../material-module';
import { CreateMasterComponent } from './create-master/create-master.component';



@NgModule({
  declarations: [
    MasterMainComponent,
    CreateMasterComponent
  ],
  imports: [
    CommonModule,
    MasterModuleRoutingModule,
    MaterialModule
  ]
})
export class MasterModuleModule { }
