import { defineConfig } from '@tarojs/cli';

export default defineConfig({
  projectName: 'railway-mini-program',
  date: '2026-6-17',
  designWidth: 750,
  deviceRatio: { 640: 2.34 / 2, 750: 1, 828: 1.81 / 2 },
  sourceRoot: 'src',
  outputRoot: 'dist',
  plugins: [],
  framework: 'react',
  compiler: 'webpack5',
  cache: { enable: false },
  mini: {
    postcss: {
      pxtransform: { enable: true },
    },
  },
  h5: {},
});
