import { describe, expect, it } from 'vitest';
import { appConfig } from './config';

describe('appConfig', () => {
  it('exposes APP_URL and BACKEND_URL', () => {
    expect(appConfig.appUrl).toMatch(/^https?:\/\//);
    expect(appConfig.backendUrl).toMatch(/^https?:\/\//);
  });
});
