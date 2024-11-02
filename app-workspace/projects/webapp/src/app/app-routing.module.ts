import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NavComponent } from './nav/nav.component';

const routes: Routes = [
  { path: '', redirectTo: 'main/qi/quotation', pathMatch: 'full' },
  {
    path: 'main',
    component: NavComponent,
    children: [
      {
        path: 'qi',
        loadChildren: () => import('./qi/qi.module').then(m => m.QiModule)
      },
      {
        path: 'master',
        loadChildren: () => import('./master-module/master-module.module').then(m => m.MasterModuleModule)
      },
    ]
  },
  {
    path: 'login',
    loadChildren: () => import('./auth/auth.module').then(m => m.AuthModule)
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
