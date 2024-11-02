import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrl: './nav.component.css'
})
export class NavComponent {

  navList = <any>[
    { dname: 'Home', route: 'home', icon: 'home' },
    { dname: 'Quotation', route: 'qi/quotation', icon: 'format_quote' },
    { dname: 'Invoice', route: 'qi/invoice', icon: 'monetization_on' },
    { dname: 'Master', route: 'master/master', icon: 'inventory_2' },
  ]
  userInfo: any;

  constructor(private routingNavigate: Router) { }

  ngOnInit(): void {
    if (!!localStorage.getItem("user")) {
      this.userInfo = JSON.parse(localStorage.getItem("user") || "")
    }
  }


  signOut() {
    localStorage.clear();
    this.routingNavigate.navigate(['/login']);
  }

}
