import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MasterMainComponent } from './master-main/master-main.component';

const routes: Routes = [
  {
    path: ':type',
    component: MasterMainComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MasterModuleRoutingModule { }
