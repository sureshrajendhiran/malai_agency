import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PaginationComponent } from './pagination/pagination.component';
import {
  Safe, UtcToCurrentRelative, UtcToCheckCurrentRelative,
  DateToNowTime, UtcToLocalTime, MomentDate, NotAvailable, NotAvailableCSS, FormatDateTime
} from './pipes/common-pipe.pipe';
import { LoaderComponent } from './loader/loader.component';
import { MaterialModule } from '../material-module';

@NgModule({
  declarations: [Safe, UtcToCurrentRelative, UtcToCheckCurrentRelative, PaginationComponent,
    DateToNowTime, UtcToLocalTime, MomentDate, NotAvailable, NotAvailableCSS, FormatDateTime, LoaderComponent],
  imports: [
    CommonModule,
    MaterialModule
  ],
  exports: [
    Safe, UtcToCurrentRelative, UtcToCheckCurrentRelative, LoaderComponent, PaginationComponent,
    DateToNowTime, UtcToLocalTime, MomentDate, NotAvailable, NotAvailableCSS, FormatDateTime
  ]
})
export class SharedModule { }
