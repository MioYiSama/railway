import { describe, expect, it } from 'vitest';
import { appLink } from './site';

describe('appLink', () => {
  it('joins path onto the app url', () => {
    expect(appLink('/login')).toMatch(/\/login$/);
  });
});
