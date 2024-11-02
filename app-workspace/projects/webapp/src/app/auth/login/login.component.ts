import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { CommonApiService } from '../../services/common-api.service'
import { Router } from '@angular/router';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginFormGroup = new FormGroup({
    email: new FormControl('', Validators.email),
    password: new FormControl('', Validators.required)
  });
  loginUserId = ''
  isNewDevice: boolean = false;
  invalidOtp = false;
  otpExpTime = 100;
  otpApiBody = {
    otp: '',
    type: '',
    save_device: false
  };
  otpValue = ['', '', '', '', '', ''];
  passwordErrorCallback = '';
  emailErrorCallback = '';
  limitExitErrorCallback = '';
  passFieldType = 'password'
  timer: any;
  isSubmiting: boolean = false;
  isVerfying: boolean = false;
  isResending: boolean = false;
  appId: any;
  productList = <any>[];
  productLoading:boolean=true;
  constructor(private commonApiService: CommonApiService,
    private router: Router,
    private snackBar: MatSnackBar) { }

  ngOnInit(): void {
    this.signOut();
  }
  onClick(event:any) {
    event.target.value = '';
  }


  verify() {

  }
  reSend() {

  }




  login() {
    this.emailErrorCallback = ''
    this.passwordErrorCallback = ''
    this.limitExitErrorCallback = ''
    this.isSubmiting = true;
    if (!this.loginFormGroup.valid) {
      if (this.loginFormGroup.controls.email.status != "VALID" || !this.loginFormGroup.controls.email.value) {
        this.emailErrorCallback = 'Please enter a vaild email ID'
      }
      if (this.loginFormGroup.controls.password.status != "VALID") {
        this.passwordErrorCallback = 'Password is required';
      }
      this.isSubmiting = false;
    } else {
      const body = <any>{};
      Object.assign(body, this.loginFormGroup.value);
      this.commonApiService.login(body).subscribe((res:any)=> {
        this.isSubmiting = false;
        if (res.statusCode === 200) {
          if(res.info.success || (!!res.info.token && !!res.info.user_info)){
            this.toast('Logged in Successfully');
            this.afterLoginPageSet(res);
          }
        }
      })
    }
  }
  // After successfull login handle
  afterLoginPageSet(res:any) {
    localStorage.setItem("user",JSON.stringify(res.info.user_info));
    localStorage.setItem("x-token",res.info.token);
    this.router.navigate(['/main/qi/quotation']);
  }

  openCRM() {
    window.open('https://talk2ship.myairliftusa.com/', '_self');
  }

  // Toast message
  toast(value:any) {

  }

  signOut() {
    localStorage.removeItem('user');
    localStorage.removeItem('x-token');
  }



}
