import { Injectable } from '@angular/core';
import { Observable, Observer, of, Subject } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import { AnonymousSubject } from 'rxjs/internal-compatibility';

@Injectable({
  providedIn: 'root'
})
export class DcsService {
  // The default address of the DCS GO server. (with port but without protocol)
  public address = 'localhost:3333';

  protected exportAPI$: Subject<any>;

  constructor(
    protected readonly httpclient: HttpClient
  ) { }

  /**
   * Queries the hello end point.
   * If it returns a 'hello' string.
   */
  public hello(): Observable<string> {
    return this.httpclient.get<string>('http://' + this.address + '/hello');
  }

  public connectWebsocket(): Subject<any> {
    if(!this.exportAPI$) {
      this.exportAPI$ = this.createWebsocket();
    }
    return this.exportAPI$;
  }

  private createWebsocket(): AnonymousSubject<any> {
    const ws = new WebSocket('ws://' + this.address + '/raw');
    const observable = new Observable((obs: Observer<any>) => {
      ws.onmessage = obs.next.bind(obs);
      ws.onerror = obs.error.bind(obs);
      ws.onclose = obs.complete.bind(obs);
      return ws.close.bind(ws);
    });
    const observer = {
      error: null,
      complete: null,
      next: (data: any) => {
        console.log('Message sent to websocket: ', data);
        if (ws.readyState === WebSocket.OPEN) {
          ws.send(data);
        }
      }
    };
    return new AnonymousSubject<any>(observer, observable);
  };


}
