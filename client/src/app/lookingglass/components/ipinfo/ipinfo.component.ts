import { Component, OnInit, ViewChild } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { IpService } from '../../services/ip.service';
import { IpinfoAddEditComponent } from '../ipinfo-add-edit/ipinfo-add-edit.component';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { IP, IPInfo, Audit } from '../../models';
import { CoreService } from '../../services/core.service';
import { IpinfoDeleteComponent } from '../ipinfo-delete/ipinfo-delete.component';

@Component({
  selector: 'app-ipinfo',
  templateUrl: './ipinfo.component.html',
  styleUrls: ['./ipinfo.component.scss'],
})
export class IpinfoComponent implements OnInit {
  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  displayedColumns: string[] = [
    // 'id',
    'host',
    'ip',
    'port',
    'username',
    'vendor',
    'action',
  ];
  dataSource!: MatTableDataSource<IP>;
  audit?: Audit;

  constructor(
    private _dialog: MatDialog,
    private _ip: IpService,
    private _core: CoreService
  ) {}

  ngOnInit(): void {
    this.getIPs({});

    this._ip.getAudit({}).subscribe({
      next: (res: { data: Audit; message: string }) => (this.audit = res.data),
      error: (err: any) => {
        console.log(err);
      },
    });
  }

  openForm(data: any = null) {
    let config = {};
    if (data) {
      config = { data };
    }
    const dialogRef = this._dialog.open(IpinfoAddEditComponent, { data });
    dialogRef.afterClosed().subscribe({
      next: (val: any) => {
        if (val) {
          this.getIPs({});
        }
      },
      error: (err: any) => {
        console.log(err);
      },
    });
  }

  getIPs(data: any) {
    this._ip.list(data).subscribe({
      next: (res: any) => {
        this.dataSource = new MatTableDataSource(res.data);
        this.dataSource.sort = this.sort;
        this.dataSource.paginator = this.paginator;
      },
      error: (err: any) => {
        console.log(err);
      },
    });
  }

  deteleIP(data: IPInfo) {
    const dialogRef = this._dialog.open(IpinfoDeleteComponent, { data });
    this._core.messageAlert('Record Deleting......');
    dialogRef.afterClosed().subscribe({
      next: (val: any) => {
        if (val) {
          this._ip.delete(data.ip, {}).subscribe({
            next: (res: any) => {
              this._core.messageAlert('Record Deletion Succeed');
              this.getIPs({});
            },
            error: (err: any) => {
              console.log(err);
            },
          });
        }
      },
      error: (err: any) => {
        console.log(err);
      },
    });
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();

    if (this.dataSource.paginator) {
      this.dataSource.paginator.firstPage();
    }
  }
}
