import { Component, OnInit } from '@angular/core';
import { DcsService } from '../../services/dcs.service';
import { FormControl, Validators } from '@angular/forms';
import { catchError, tap } from 'rxjs/operators';
import { of } from 'rxjs';

@Component({
  selector: 'app-tab1',
  templateUrl: 'tab1.page.html',
  styleUrls: ['tab1.page.scss']
})
export class Tab1Page implements OnInit {

  readonly addressControl: FormControl<string> = new FormControl<string>('', [Validators.required]);

  constructor(
    protected readonly dcsService: DcsService
  ) {}

  ngOnInit() {
  }


  connect() {
    if (this.addressControl.invalid) {return;}

    console.log('Connecting to dcs service...');
    this.dcsService.address = this.addressControl.value;

    this.dcsService.connectWebsocket().pipe(tap(() => {
      console.log('Connected to dcs service.');
    }), catchError((err) => {
      console.log('Could not connect to dcs service.');
      return of(err);
    })).subscribe((msg) => {
      console.log('Received message from dcs service: ', msg);
    });

  }

}
