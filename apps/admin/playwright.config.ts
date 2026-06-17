import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './e2e',
  use: { baseURL: 'http://localhost:5273' },
  webServer: {
    command: 'pnpm dev',
    port: 5273,
    reuseExistingServer: !process.env.CI,
  },
});
