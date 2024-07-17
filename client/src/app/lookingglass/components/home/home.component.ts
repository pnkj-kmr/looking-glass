import {
  AfterViewInit,
  Component,
  ElementRef,
  Inject,
  OnDestroy,
  OnInit,
  Renderer2,
  ViewChild,
} from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Observable } from 'rxjs';
import { map, startWith } from 'rxjs/operators';
import { HomeService } from '../../services/home.service';
import { CoreService } from '../../services/core.service';
import { WebsocketService } from '../../services/websocket.service';
import { Protocol, SrcHost, Result } from '../../models';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss'],
  providers: [WebsocketService],
})
export class HomeComponent implements OnInit, AfterViewInit, OnDestroy {
  @ViewChild('query_canvas') canvasRef!: ElementRef;

  queryForm: FormGroup;
  captchaStr!: string;
  public captchaCheck: Boolean = false;
  public status: string = '';
  public result: string[] = [];
  public resultError: boolean = false;
  private connectionErr: boolean = false;
  public protocols: Protocol[] = [];
  public processing: boolean = false;
  public srcHosts: SrcHost[] = [];
  public srcHostsOptions?: Observable<SrcHost[]>;

  constructor(
    private _fb: FormBuilder,
    private _home: HomeService,
    private _core: CoreService,
    @Inject(DOCUMENT) private document: Document,
    private renderer: Renderer2,
    private _ws: WebsocketService
  ) {
    this.queryForm = this._fb.group({
      src: '',
      proto: '',
      dst: '',
      captcha: '',
    });
  }

  ngOnInit(): void {
    this.messageSubscription();
    this._home.protocols({}).subscribe({
      next: (res: Protocol[]) => {
        this.protocols = res;
      },
      error: (err: any) => {
        console.log(err);
      },
    });
    this._home.srchosts({}).subscribe({
      next: (res: SrcHost[]) => {
        this.srcHosts = res;
        // added for auto complete
        this.srcHostsOptions = this.queryForm.get('src')!.valueChanges.pipe(
          startWith(''),
          map((value) => {
            return this._filter(value || '');
          })
        );
      },
      error: (err: any) => {
        console.log(err);
      },
    });
    // added for auto complete
    this.srcHostsOptions = this.queryForm.get('src')!.valueChanges.pipe(
      startWith(''),
      map((value) => {
        return this._filter(value || '');
      })
    );
  }

  ngAfterViewInit(): void {
    const child = this.getCaptcha();
    this.renderer.appendChild(this.canvasRef.nativeElement, child);
    this.captchaCheck = false;
  }

  ngOnDestroy(): void {
    this._ws.message.complete();
  }

  displayFn(src: SrcHost): string {
    return src && src?.host ? src.host : '';
  }

  private _filter(src: string | SrcHost): SrcHost[] {
    let fiterValue = '';
    if (typeof src === 'string') {
      fiterValue = src.toLowerCase();
    } else if (typeof src === 'object') {
      fiterValue = src.host.toLowerCase();
    }
    return this.srcHosts.filter((src) =>
      src.host.toLowerCase().includes(fiterValue)
    );
  }

  onFormReset() {
    this.captchaCheck = false;
    this.status = '';
    this.result = [];
    this.refreshCaptcha();

    this._ws.message.unsubscribe();
    this.messageSubscription();
  }

  messageSubscription(retry: boolean = true) {
    this._ws.message.subscribe({
      next: (msg: Result) => {
        // console.log('message received ---- ' + msg);
        if (msg.completed) {
          this.processing = false;
          this.resultError = msg.is_error;
          this.status = '';
          this.result.push(msg.message);
          this.refreshCaptcha();
        } else {
          this.resultError = false;
          this.status = 'processing...';
          this.result.push(msg.message);
        }
      },
      error: (err) => {
        console.log('error....', err);
        this._ws.message.unsubscribe();
        if (retry) {
          this.messageSubscription(false);
          this.connectionErr = true;
        }
      },
      complete: () => {
        // console.log('complete');
        this.processing = false;
      },
    });
  }

  messagePush(message: any) {
    if (this.connectionErr) {
      this.connectionErr = false;
      this.messageSubscription();
    }
    this._ws.message.next(message);
  }

  onFormSubmit() {
    this.resultError = false;
    if (this.queryForm.valid) {
      let data = this.queryForm.value;
      // vaildate and pre-check for src input
      if (
        this.srcHosts?.find(
          (val) => val.ip === data?.src?.ip && val.host === data?.src?.host
        )
      ) {
        data = { ...data, src: data.src.ip };
      } else {
        this.queryForm.patchValue({ src: '' });
        return;
      }
      // captcha validation
      if (this._core.compareString(this.captchaStr, data.captcha)) {
        this.captchaCheck = false;
        this.status = 'processing...';
        this.result = [];
        this.messagePush(data);
        this.processing = true;
      } else {
        this.captchaCheck = true;
        this.queryForm.patchValue({ captcha: '' });
      }
    }
  }

  refreshCaptcha() {
    const el = this.document.getElementById('captcha_canvas');
    this.renderer.removeChild(this.canvasRef.nativeElement, el);
    const child = this.getCaptcha();
    this.renderer.appendChild(this.canvasRef.nativeElement, child);
    this.captchaCheck = false;
    this.processing = false;
    this.resultError = false;
  }

  getCaptcha() {
    this.captchaStr = this._core.getRandomString();
    let canv = this.document.createElement('canvas');
    canv.height = 20;
    canv.width = 120;
    canv.id = 'captcha_canvas';
    let ctx = canv.getContext('2d');
    ctx!.font = '15px Georgia';
    ctx!.strokeText(this.captchaStr, 15, 15);
    ctx!.lineWidth = 2;
    ctx!.moveTo(0, 0);
    ctx!.stroke();
    return canv;
  }
}
