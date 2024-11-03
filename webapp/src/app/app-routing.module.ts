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
      {
        path: 'stock',
        loadChildren: () => import('./stock-manage/stock-manage.module').then(m => m.StockManageModule)
      },
      {
        path: 'home',
        loadChildren: () => import('./home/home.module').then(m => m.HomeModule)
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
