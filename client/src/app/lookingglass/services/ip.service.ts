import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { IPInfo } from '../models';
import { environment } from '../../../environments';

@Injectable({
  providedIn: 'root',
})
export class IpService {
  private endpoint: string = environment.api + '/api/ip';
  private endpointAudit: string = environment.api + '/api/audit';

  constructor(private _http: HttpClient) {}

  add(data: IPInfo): Observable<any> {
    return this._http.post(this.endpoint, data);
  }

  delete(id: string, data: any): Observable<any> {
    return this._http.delete(`${this.endpoint}/${id}`, data);
  }

  edit(id: string, data: IPInfo): Observable<any> {
    return this._http.post(`${this.endpoint}/${id}`, data);
  }

  list(data: any): Observable<any> {
    return this._http.get(this.endpoint, data);
  }

  getVendors(data: any): Observable<any> {
    return this._http.get(`${this.endpoint}/vendor`, data);
  }

  getAudit(data: any): Observable<any> {
    return this._http.get(this.endpointAudit, data);
  }
}
