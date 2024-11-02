import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { provideAnimationsAsync } from '@angular/platform-browser/animations/async';
import { NavComponent } from './nav/nav.component';
import { MaterialModule } from './material-module';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { TokenInterceptor } from './auth/token/token.interceptor';
import { ErrorsInterceptor } from './auth/errors/errors.interceptor';

@NgModule({
  declarations: [
    AppComponent,
    NavComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: TokenInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorsInterceptor, multi: true },
    provideAnimationsAsync(),
    // {
    //   provide: SwRegistrationOptions,
    //   useFactory: () => ({ enabled: environment.production }),
    // },

  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
