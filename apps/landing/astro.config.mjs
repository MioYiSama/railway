import { defineConfig } from 'astro/config';

// 静态站点（默认 output: 'static'），构建产物挂载到 Caddy。
export default defineConfig({
  site: 'https://railway.local',
});
