import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './lookingglass/components/home/home.component';
import { IpinfoComponent } from './lookingglass/components/ipinfo/ipinfo.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'ip', component: IpinfoComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
