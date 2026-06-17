import { defineConfig } from 'tsdown';

export default defineConfig({
  entry: ['src/index.ts'],
  format: ['esm'],
  dts: { tsgo: true },
  clean: true,
  deps: { neverBundle: ['react', 'react/jsx-runtime', '@tarojs/components'] },
});
