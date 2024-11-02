import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';


@Component({
  selector: 'app-pagination',
  templateUrl: './pagination.component.html',
  styleUrl: './pagination.component.css'
})
export class PaginationComponent {
  @Input() inputObject: any;
  @Input() getObj: any;
  @Input() isLoading: boolean = false;
  @Output() pageChangeEmitter = new EventEmitter<any>();
  limitList = [5, 10, 25, 50, 100];

  constructor() { }

  ngOnInit(): void {
  }

  pagination(value?: any) {
    if(!this.isLoading){
      if (value === 1) {
        if (this.inputObject.dataList.length === this.getObj.limit) {
          this.getObj.page = this.getObj.page + 1;
          this.changePage();
        }
      } else if (value === -1) {
        if (this.getObj.page !== 0) {
          this.getObj.page = this.getObj.page - 1;
          this.changePage();
        }
      }
    }
  }
  changePage(){
    this.pageChangeEmitter.emit(this.getObj);
  }
}
