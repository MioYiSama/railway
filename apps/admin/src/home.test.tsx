import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { render, screen } from '@testing-library/react';
import { describe, expect, it } from 'vitest';
import { Home } from './home';

describe('Home', () => {
  it('renders the heading', () => {
    const queryClient = new QueryClient();
    render(
      <QueryClientProvider client={queryClient}>
        <Home />
      </QueryClientProvider>,
    );
    expect(screen.getByRole('heading', { name: 'Railway Admin' })).toBeInTheDocument();
  });
});
