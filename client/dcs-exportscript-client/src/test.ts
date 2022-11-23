import 'jest-preset-angular/setup-jest';
import { ngMocks } from 'ng-mocks';

Object.defineProperty(window, 'CSS', {value: null});
Object.defineProperty(window, 'getComputedStyle', {
  value: () => ({
      display: 'none',
      appearance: ['-webkit-appearance']
    })
});

Object.defineProperty(document, 'doctype', {
  value: '<!DOCTYPE html>'
});
Object.defineProperty(document.body.style, 'transform', {
  value: () => ({
      enumerable: true,
      configurable: true
    })
});

// All methods in mock declarations and providers
// will be automatically spied on their creation.
// https://ng-mocks.sudo.eu/extra/auto-spy
ngMocks.autoSpy('jest');
