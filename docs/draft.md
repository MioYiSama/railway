# Draft

## 前端

- React 19 + React Compiler
- TanStack Router
- TanStack Query/Form/Virtual/Table
- React Native
- Taro v4
- Astro v7 (beta)
- Vite v8 & tsdown & oxlint (+ oxlint-tsgolint `--type-aware`) & oxfmt
- TypeScript 6
- vitest + jsdom + testing library + playwright
- Node 26

## 后端

- go 1.26
- fiber v3
- sqlx
- gRPC
- Kafka

## 数据

- PostgreSQL 18
- Redis 8
- SeaweedFS

## 可观测性

- OTel
- Grafana
- Prometheus
- Tempo
- Loki

## 其他

- Nx
- GitHub Actions
- Docker & Docker Compose
- Caddy
- 前端均为静态SPA，编译成静态文件之后直接挂载到caddy容器内
- 不接入任何第三方服务（支付、短信、邮件、第三方登录等），全部mock，只作为本地演示+学习用途
- 前端至少要有两个环境变量：APP_URL和BACKEND_URL，且区分development和production模式

## UI

- packages/ui-core 定义统一组件 API
- packages/ui-web Web 实现
- packages/ui-taro Taro 实现
- packages/ui-native RN 实现
