import { TestBed } from '@angular/core/testing';

import { DcsService } from './dcs.service';

describe('DcsService', () => {
  let service: DcsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DcsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
