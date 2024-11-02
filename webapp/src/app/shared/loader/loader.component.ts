import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-loader',
  templateUrl: './loader.component.html',
  styleUrls: ['./loader.component.css']
})
export class LoaderComponent implements OnInit {

  @Input() size:any;
  constructor() { }

  ngOnInit(): void {
    if(!this.size){
      this.size='60'
    }
  }
}
