# Railway

> 一个仿 12306 的铁路购票系统，仅供学习参考。

[![CI](https://github.com/MioYiSama/railway/actions/workflows/ci.yml/badge.svg)](https://github.com/MioYiSama/railway/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Node](https://img.shields.io/badge/Node-%E2%89%A526-5FA04E?logo=node.js&logoColor=white)](https://nodejs.org/)
[![pnpm](https://img.shields.io/badge/pnpm-11-F69220?logo=pnpm&logoColor=white)](https://pnpm.io/)
[![TypeScript](https://img.shields.io/badge/TypeScript-6-3178C6?logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Nx](https://img.shields.io/badge/Nx-23-143055?logo=nx&logoColor=white)](https://nx.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

---

## 技术栈

| 层次 | 技术 |
|------|------|
| 微服务 | Go 1.26 · Fiber v3 · gRPC · Protocol Buffers |
| 可观测性 | OpenTelemetry · Grafana · Tempo · Prometheus · Loki |
| Web 前端 | React 19 · Vite 8 · TanStack Router / Query · TypeScript 6 |
| 落地页 | Astro |
| 移动端 | React Native |
| 小程序 | Taro |
| 工程化 | Nx 23 · pnpm 11 · React Compiler · Oxlint · Vitest · Playwright |
| 基础设施 | Docker Compose · Caddy |

---

## 项目结构

```
railway/
├── apps/
│   ├── admin           # 管理后台（React）
│   ├── landing         # 官网落地页（Astro）
│   ├── mini-program    # 微信小程序（Taro）
│   ├── mobile          # 移动端（React Native）
│   └── web             # 主 Web 应用（React + TanStack）
├── services/
│   ├── api-gateway     # 接入网关（Fiber v3）
│   ├── train-schedule  # 车次与时刻表
│   ├── ticket-query    # 余票与票价查询
│   ├── inventory-seat  # 库存与座位分配
│   ├── order           # 订单生命周期
│   ├── ticketing       # 出票 / 改签 / 退票
│   ├── user            # 用户与实名认证
│   ├── payment         # 聚合支付
│   └── notification    # 通知中心（短信 / Push / 邮件）
├── libs/               # Go 共享库
│   ├── proto           # Protobuf 定义与生成代码
│   ├── config          # 配置加载
│   ├── httpserver      # HTTP 服务器封装
│   ├── grpcserver      # gRPC 服务器封装
│   ├── observability   # OTel 初始化
│   └── gen             # 代码生成产物
├── packages/           # TypeScript 共享库
│   ├── ui-core         # 基础 UI 原语
│   ├── ui-web          # Web 组件库
│   ├── ui-native       # React Native 组件库
│   └── ui-taro         # Taro 组件库
├── docker/             # 本地开发 Docker 配置
│   ├── docker-compose.yml
│   ├── caddy/          # 反向代理
│   ├── grafana/        # 可视化面板
│   ├── prometheus/     # 指标采集
│   ├── loki/           # 日志聚合
│   ├── tempo/          # 链路追踪
│   └── otel-collector/ # OTel 收集器
└── docs/
    └── api/            # OpenAPI & AsyncAPI 文档
```

---

## 微服务架构

```
                          客户端（Web / Mobile / 小程序）
                                    │
                              API Gateway
                         （路由 · 鉴权 · 协议转换）
                                    │
              ┌─────────┬───────────┼───────────┬─────────┐
              │         │           │           │         │
       Train-Schedule  Ticket    Inventory   Order   Ticketing
        车次时刻表     Query      & Seat      订单     出票/退改
                       余票查询   座位库存
              └─────────┴───────────┴───────────┴─────────┘
                                    │
                    ┌───────────────┼───────────────┐
                    │               │               │
                   User          Payment       Notification
                  用户认证        聚合支付       通知中心
```

服务间通过 **gRPC** 通信，所有 Protobuf 定义统一维护在 `libs/proto`。

---

## 快速开始

### 前置依赖

- Go ≥ 1.26
- Node ≥ 26 + pnpm 11
- Docker & Docker Compose

### 安装依赖

```bash
pnpm install
```

### 启动本地基础设施

```bash
docker compose -f docker/docker-compose.yml up -d
```

### 生成 Protobuf 代码

```bash
pnpm nx run libs-proto:codegen
# 或
pnpm codegen
```

---

## 常用命令

| 命令 | 说明 |
|------|------|
| `pnpm build` | 构建所有项目 |
| `pnpm test` | 运行所有测试 |
| `pnpm lint` | 检查所有项目 |
| `pnpm fmt` | 格式化代码 |
| `pnpm nx run <project>:<target>` | 执行单个项目任务 |
| `pnpm nx affected -t test` | 只测试受影响的项目 |
| `pnpm nx graph` | 可视化项目依赖图 |
