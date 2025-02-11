import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { IPInfo, Vendor } from '../../models';
import { CoreService } from '../../services/core.service';
import { IpService } from '../../services/ip.service';

@Component({
  selector: 'app-ipinfo-add-edit',
  templateUrl: './ipinfo-add-edit.component.html',
  styleUrls: ['./ipinfo-add-edit.component.scss'],
})
export class IpinfoAddEditComponent implements OnInit {
  ipForm: FormGroup;

  public vendors: Vendor[] = [];

  constructor(
    private _fb: FormBuilder,
    private _ip: IpService,
    private _dialogRef: MatDialogRef<IpinfoAddEditComponent>,
    @Inject(MAT_DIALOG_DATA) public formData: IPInfo,
    private _core: CoreService
  ) {
    this.ipForm = this._fb.group({
      ip: '',
      host: '',
      port: '',
      username: '',
      password: '',
      vendor: '',
    });
  }

  ngOnInit(): void {
    this.ipForm.patchValue(this.formData);
    this._ip.getVendors({}).subscribe({
      next: (res: Vendor[]) => {
        this.vendors = res;
      },
      error: (err: any) => {
        console.log(err);
      },
    });
  }

  onFormSubmit() {
    if (this.ipForm.valid) {
      // TODO -need to add precheck for password match
      if (this.formData) {
        this._ip.edit(this.formData.ip, this.ipForm.value).subscribe({
          next: (res: any) => {
            this._core.messageAlert('Record Updated Successfully');
            this._dialogRef.close(true);
          },
          error: (err: any) => {
            // console.log(err);
            this._core.messageAlert(`${err?.error?.error || "Record failed"}`, "ERROR");
            this._dialogRef.close(true);
          },
        });
      } else {
        this._ip.add(this.ipForm.value).subscribe({
          next: (res: any) => {
            this._core.messageAlert('New Record Added Successfully');
            this._dialogRef.close(true);
          },
          error: (err: any) => {
            // console.log(err);
            this._core.messageAlert(`${err?.error?.error || "Record failed"}`, "ERROR");
            this._dialogRef.close(true);
          },
        });
      }
    }
  }
}
