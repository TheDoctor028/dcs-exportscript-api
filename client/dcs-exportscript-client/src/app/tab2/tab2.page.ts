import { Component, OnInit } from '@angular/core';
import { DcsService } from '../../services/dcs.service';

@Component({
  selector: 'app-tab2',
  templateUrl: 'tab2.page.html',
  styleUrls: ['tab2.page.scss']
})
export class Tab2Page implements OnInit{

  public dataVal = '000000';

  constructor(
    protected readonly dcsService: DcsService,
  ) {}

  public sendButtonPress(val: number) {
    this.dcsService.connectWebsocket().next(`C12,302${val},1`);
    this.dcsService.connectWebsocket().next(`C12,302${val},0`);
  }

  ngOnInit() {
    this.dcsService.connectWebsocket().pipe().subscribe((res) => {
      console.log(res);
      this.dataVal = res.data;
    });
  }

}
