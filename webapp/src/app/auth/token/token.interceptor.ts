import { Injectable } from '@angular/core';
import {
  HttpInterceptor,
  HttpRequest,
  HttpResponse,
  HttpHandler,
  HttpEvent,
  HttpErrorResponse
} from '@angular/common/http';

import { ActivatedRoute, Router } from '@angular/router';
import moment from 'moment';

import { Observable } from 'rxjs'

@Injectable()
export class TokenInterceptor implements HttpInterceptor {
  constructor(private route: Router,
    private activeRoute: ActivatedRoute) {

  }
  intercept(request: HttpRequest<any>, next: HttpHandler): any {

    // add authorization header with jwt token if available
    const token = localStorage.getItem('x-token');
    if (token) {
      request = request.clone({
        setHeaders: {
          'x-token': token,
          'current-time': moment().format('YYYY-MM-DD HH:mm:ss'),
          'utc-time': moment().utc().format('YYYY-MM-DD HH:mm:s')
        }
      });
      return next.handle(request);
    } else {
      // Only for login time
      if (request.url.includes('login') || request.url.includes('otp_callback') || request.url.includes('general/products')) {
        request = request.clone({
          setHeaders: {
            'current-time': moment().format('YYYY-MM-DD HH:mm:ss'),
            'utc-time': moment().utc().format('YYYY-MM-DD HH:mm:s')
          }
        });
        return next.handle(request);
      } else {
        this.route.navigateByUrl('/login');
      }

    }
  }

}
