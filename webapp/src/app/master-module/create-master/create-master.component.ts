import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-create-master',
  templateUrl: './create-master.component.html',
  styleUrl: './create-master.component.css'
})
export class CreateMasterComponent {
  @Input() operation: any;
  @Input() type: any;
  @Input() inputInfo: any;
  @Output() closeEvent = new EventEmitter();


  constructor(){
    
  }
}
