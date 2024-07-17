import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { environment } from '../../../environments';

export interface Message {
  author: string;
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class HomeService {
  private endpoint: string = environment.api + '/api/lg';

  constructor(private _http: HttpClient) {}

  // submit(data: any): Observable<any> {
  //   return this._http.post(this.endpoint, data);
  // }

  protocols(data: any): Observable<any> {
    return this._http.get(`${this.endpoint}/protocol`, data);
  }

  srchosts(data: any): Observable<any> {
    return this._http.get(`${this.endpoint}/src`, data);
  }
}
