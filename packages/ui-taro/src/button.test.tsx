import { fireEvent, render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import { Button } from './button.js';

describe('Button (taro)', () => {
  it('renders the label', () => {
    render(<Button label="改签" />);
    expect(screen.getByRole('button', { name: '改签' })).toBeInTheDocument();
  });

  it('fires onPress when clicked', () => {
    const onPress = vi.fn();
    render(<Button label="改签" onPress={onPress} />);
    fireEvent.click(screen.getByRole('button'));
    expect(onPress).toHaveBeenCalledOnce();
  });
});
