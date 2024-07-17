import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-toolbar',
  templateUrl: './toolbar.component.html',
  styleUrls: ['./toolbar.component.scss'],
})
export class ToolbarComponent {
  title: string = 'Looking Glass';
  version: string = 'v2.0.0';

  constructor(private router: Router) {}

  navigate(url: string) {
    this.router.navigateByUrl(url);
  }
}
