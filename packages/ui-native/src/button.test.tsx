import { fireEvent, render, screen } from '@testing-library/react';
import { describe, expect, it, vi } from 'vitest';
import { Button } from './button.js';

describe('Button (native)', () => {
  it('renders the label', () => {
    render(<Button label="退票" />);
    expect(screen.getByText('退票')).toBeInTheDocument();
  });

  it('fires onPress when pressed', () => {
    const onPress = vi.fn();
    render(<Button label="退票" onPress={onPress} />);
    fireEvent.click(screen.getByRole('button'));
    expect(onPress).toHaveBeenCalledOnce();
  });
});
