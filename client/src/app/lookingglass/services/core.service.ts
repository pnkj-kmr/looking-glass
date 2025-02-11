import { Injectable } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';

@Injectable({
  providedIn: 'root',
})
export class CoreService {
  constructor(private _snackBar: MatSnackBar) {}

  messageAlert(data: string, action: string = 'OK') {
    this._snackBar.open(data, action, {
      duration: 3000,
      verticalPosition: 'bottom',
    });
  }

  getRandomString(): string {
    let alpha = new Array(
      'A',
      'B',
      'C',
      'D',
      'E',
      'F',
      'G',
      'H',
      'I',
      'J',
      'K',
      'L',
      'M',
      'N',
      'O',
      'P',
      'Q',
      'R',
      'S',
      'T',
      'U',
      'V',
      'W',
      'X',
      'Y',
      'Z',
      'a',
      'b',
      'c',
      'd',
      'e',
      'f',
      'g',
      'h',
      'i',
      'j',
      'k',
      'l',
      'm',
      'n',
      'o',
      'p',
      'q',
      'r',
      's',
      't',
      'u',
      'v',
      'w',
      'x',
      'y',
      'z',
      '1',
      '2',
      '3',
      '4',
      '5',
      '6',
      '7',
      '8',
      '9',
      '0'
    );
    let a = alpha[Math.floor(Math.random() * alpha.length)];
    let b = alpha[Math.floor(Math.random() * alpha.length)];
    let c = alpha[Math.floor(Math.random() * alpha.length)];
    let d = alpha[Math.floor(Math.random() * alpha.length)];
    let e = alpha[Math.floor(Math.random() * alpha.length)];
    let f = alpha[Math.floor(Math.random() * alpha.length)];
    let g = alpha[Math.floor(Math.random() * alpha.length)];

    let code =
      a + ' ' + b + ' ' + ' ' + c + ' ' + d + ' ' + e + ' ' + f + ' ' + g;

    return code;
  }

  private removeSpaces(val: string): string {
    return val.split(' ').join('');
  }

  compareString(str1: string, str2: string): Boolean {
    let s1 = this.removeSpaces(str1);
    let s2 = this.removeSpaces(str2);
    return s1 === s2;
  }
}
