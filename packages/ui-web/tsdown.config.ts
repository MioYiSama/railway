import { defineConfig } from 'tsdown';

export default defineConfig({
  entry: ['src/index.ts'],
  format: ['esm'],
  dts: { tsgo: true },
  tsconfig: './tsconfig.build.json',
  clean: true,
  deps: { neverBundle: ['react', 'react-dom', 'react/jsx-runtime'] },
});
