import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { StockManageComponent } from './stock-manage/stock-manage.component';
import { StockModuleRoutingModule } from './stock-routing';
import { MaterialModule } from '../material-module';



@NgModule({
  declarations: [StockManageComponent],
  imports: [
    CommonModule,
    StockModuleRoutingModule,
    MaterialModule
  ]
})
export class StockManageModule { }
