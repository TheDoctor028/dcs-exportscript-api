import { TestBed } from '@angular/core/testing';

import { DcsService } from './dcs.service';
import { HttpClient } from '@angular/common/http';

describe('DcsService', () => {
  let service: DcsService;

  let httpClient: HttpClient;

  beforeEach(() => {
    httpClient = TestBed.inject(HttpClient);

    TestBed.configureTestingModule({
      imports: [],
      providers: [
        DcsService,
        {
          provide: HttpClient,
          useValue: httpClient
        }
    ],
      declarations: []
    });
    service = TestBed.inject(DcsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  describe('hello', () => {
    it('should call http get to the given ip + /hello', () => {
      service.hello();
      expect(httpClient.get).toHaveBeenCalledWith('http://localhost:3333/hello');
    });

    it('should return hello', () => {
      expect(service.hello()).toEqual('hello');
    });
  });
});
