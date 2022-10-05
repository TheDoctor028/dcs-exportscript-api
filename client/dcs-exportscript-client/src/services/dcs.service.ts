import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class DcsService {
  constructor(private readonly httpClient: HttpClient) { }

  /**
   * Queries the hello end point.
   * If it will return a 'hello' string.
   */
  public hello(): Observable<string> {
    return this.httpClient.get();
  }
}
