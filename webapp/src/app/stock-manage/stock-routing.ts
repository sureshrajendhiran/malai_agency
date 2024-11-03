import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { StockManageComponent } from './stock-manage/stock-manage.component';
const routes: Routes = [
  {
    path: '',
    component: StockManageComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class StockModuleRoutingModule { }

