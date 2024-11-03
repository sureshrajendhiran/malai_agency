import { Component } from '@angular/core';

@Component({
  selector: 'app-stock-manage',
  templateUrl: './stock-manage.component.html',
  styleUrl: './stock-manage.component.css'
})
export class StockManageComponent {


  dataList = <any>[1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5];

  onEnter(event: Event): void {
    const inputValue = (event.target as HTMLInputElement).value;
    console.log('Enter pressed, input value:', inputValue);
    // Do something with the input value or event
  }

}
