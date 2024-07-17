import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-ipinfo-delete',
  templateUrl: './ipinfo-delete.component.html',
  styleUrls: ['./ipinfo-delete.component.scss'],
})
export class IpinfoDeleteComponent {
  constructor(private _dialogRef: MatDialogRef<IpinfoDeleteComponent>) {}

  onOK() {
    this._dialogRef.close(true);
  }
}
