import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MaterialModule } from '../material/material.module';
import { ToolbarComponent } from './components/toolbar/toolbar.component';
import { HomeComponent } from './components/home/home.component';
import { IpinfoComponent } from './components/ipinfo/ipinfo.component';
import { IpinfoAddEditComponent } from './components/ipinfo-add-edit/ipinfo-add-edit.component';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { IpinfoDeleteComponent } from './components/ipinfo-delete/ipinfo-delete.component';

@NgModule({
  declarations: [
    ToolbarComponent,
    HomeComponent,
    IpinfoComponent,
    IpinfoAddEditComponent,
    IpinfoDeleteComponent,
  ],
  imports: [
    CommonModule,
    MaterialModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  exports: [ToolbarComponent, HomeComponent, IpinfoComponent],
})
export class LookingglassModule {}
