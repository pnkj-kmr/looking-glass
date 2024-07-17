import { Injectable, Inject } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { environment } from 'src/environments';

@Injectable()
export class WebsocketService {
  private endpoint: string = '';
  private subject?: WebSocketSubject<any>;

  constructor(@Inject(DOCUMENT) private document: Document) {
    if (environment.prod) {
      let host_url = this.document.location.host;
      let protocol = 'ws://';
      if (this.document.location.protocol === 'https:') protocol = 'wss://';
      this.endpoint = `${protocol}${host_url}/ws/query`;
    } else {
      this.endpoint = environment.ws + '/ws/query';
    }
  }

  set message(_) {
    this.subject = undefined;
  }

  get message(): WebSocketSubject<any> {
    if (!this.subject || this.subject.closed) {
      this.subject = webSocket(this.endpoint);
      // console.log('ws connected...', this.subject);
    }
    // console.log('........................', this.subject);
    return this.subject;
  }
}
