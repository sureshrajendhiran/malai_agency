import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MainQiComponent } from './main-qi/main-qi.component';

const routes: Routes = [
  {
    path: ':type',
    component: MainQiComponent
  },
  {
    path: ':type/:filter',
    component: MainQiComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class QIRoutingModule { }
