# railway

> 仅作为学习用途

## 项目结构

```
railway
├── .github
│   └── workflows
├── apps                # 所有客户端应用
│   ├── admin           # 管理后台
│   ├── landing         # 官网 / 落地页（Astro）
│   ├── mini-program    # 小程序（Taro）
│   ├── mobile          # 移动端（React Native）
│   └── web             # 主 Web 应用
├── docker              # 本地开发和部署相关 docker 文件
├── docs
│   └── api             # 包含 OpenAPI 和 AsyncAPI
├── libs                # Go 共享库（包含proto定义）
├── packages            # TS 共享库
├── scripts
├── services
├── go.work
├── nx.json
├── package.json
├── pnpm-workspace.yaml
└── README.md
```

## 微服务划分

1. 接入与网关层 (Edge & Gateway)

- BFF/API网关服务 (API Gateway)：负责路由转发、协议转换、SSL卸载、统一鉴权。这是所有流量的入口。

2. 核心业务层 (Core Domain) —— 系统的骨架

- 车次客运服务 (Train Schedule Service)：负责管理全国铁路网、车站信息、列车编组、列车时刻表、停站信息。
- 查票服务 (Ticket Query Service)：独立于核心交易，专门负责余票查询、票价查询。
- 库存/座位服务 (Inventory & Seat Service)：整个系统最复杂、最核心的服务。负责车间段座位的分配、锁定、释放、真实库存扣减。
- 订单服务 (Order Service)：负责订单生命周期流转（创单、待支付、支付成功、取消等）。
- 票务服务 (Ticketing Service)：负责真正意义上的“出票”、改签、退票、电子客票生成、乘车码生成。

3. 支撑业务层 (Support Domain)

- 用户服务 (User Service)：负责账号管理、实名认证（对接公安系统）、常用联系人管理。
- 支付服务 (Payment Service)：聚合支付网关，对接支付宝、微信、银联及各大银行，处理支付回调（幂等性处理极重要）。
- 通知中心 (Notification Service)：负责短信、APP Push、邮件的异步发送。
