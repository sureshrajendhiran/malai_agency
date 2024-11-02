import { Injectable } from '@angular/core';
import { HttpRequest, HttpHandler, HttpEvent, HttpInterceptor } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

import { Router } from '@angular/router';

@Injectable()
export class ErrorsInterceptor implements HttpInterceptor {
  constructor(private route: Router) { }

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(request).pipe(catchError(err => {
      if (err.status === 401) {
        // auto logout if 401 response returned from api
        // this.authenticationService.logout();
        location.reload();
      }
      if (err.status === 403) {
        localStorage.removeItem('user');
        localStorage.removeItem('x-token');
        this.route.navigate(['login']);
        // this.toastService.success(err.error.message);
      }
      if (err.status >= 500) {
        // this.toastService.error('Contact Our Support Team..');
      }

      const error = err.error.message || err.statusText;
      return throwError(error);
    }));
  }
}
