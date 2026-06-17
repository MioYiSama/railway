import { describe, expect, it } from 'vitest';
import { resolveVariant } from './button.js';

describe('resolveVariant', () => {
  it('falls back to primary', () => {
    expect(resolveVariant()).toBe('primary');
  });

  it('keeps the explicit variant', () => {
    expect(resolveVariant('ghost')).toBe('ghost');
  });
});
