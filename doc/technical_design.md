# 个人网站技术设计文档

## 1. 架构概述
系统采用标准的前后端分离架构，并使用 Docker 进行容器化，以确保统一部署和弹性扩展。

- **客户端：** Web 浏览器（Vue 3 单页应用 SPA）
- **反向代理/Web服务器：** Nginx（处理静态文件并代理后端 API 请求）
- **应用服务器：** Go（使用 Gin 或 Fiber Web 框架）
- **数据层：** MySQL（持久化存储） & Redis（缓存及消息/会话代理）

## 2. 技术栈

### 2.1 前端
- **框架：** Vue 3 (Composition API) + Vite (极速构建)
- **UI & 样式：** Tailwind CSS + Headless UI (或 Element Plus)，主打“简洁但不朴素”的美学风格和优雅的导航栏设计。
- **路由与状态管理：** Vue Router & Pinia。
- **HTTP 客户端：** Axios。

### 2.2 后端
- **语言：** Go (Golang) - 选择其极高的并发能力与执行性能。
- **框架：** Gin 或 Echo (轻量级，极易扩展的路由)。
- **ORM 数据管理：** GORM (用于与 MySQL 交互)。
- **架构模式：** 简洁架构 / MVC (Controller -> Service -> Repository) 确保高可扩展性。

### 2.3 基础设施与存储
- **数据库：** MySQL 8.x (关系型数据：用户、博客文章、标签、工具元数据)。
- **缓存：** Redis (缓存热点文章、会话管理、限流)。
- **部署：** Docker & Docker Compose (轻松编排 Go 应用、MySQL 以及 Redis 容器)。

## 3. 系统设计细节

### 3.1 高扩展性 (后端)
- **接口驱动设计：** 针对 Services 和 Repositories 定义 Go Interfaces，从而能够轻松替换具体实现。
- **面向工具的插件/模块化系统：** 工具模块将被设计成 Go 语言的独立服务模块，允许在不影响核心博客逻辑的情况下随时挂载、即插即用各类工具（Block）。

### 3.2 高并发策略
- **Goroutines：** 充分利用 Go 的原生 goroutines 机制来处理大量并发的 HTTP 请求。
- **缓存层 (Redis)：**
  - 将频繁被访问的博文缓存在 Redis 中，极大减少 MySQL 读写压力。
  - 使用 Redis 落地限流策略（例如：令牌桶算法/Token Bucket）以保护 API 免受过度访问。
- **数据库连接池：** 在 GORM 中精细化配置 MySQL 连接池参数。

### 3.3 安全措施
- **身份验证：** 使用 JWT (JSON Web Tokens)，将其存放在 HttpOnly Cookie 或 Authorization headers 中保障安全性。
- **密码安全：** 采用 bcrypt 算法对管理员/用户密码进行不可逆哈希计算。
- **输入校验：** 启用 Go 端的 validator 库对提交的所有 JSON/表单数据进行严格过滤和验证，防注入。
- **CORS与安全标头：** 全局中间件设定严格的 CORS 跨域规则和安全的 HTTP 标头预防 XSS 及 CSRF。

## 4. 部署拓扑 (Docker)
根目录下的 `docker-compose.yml` 将定义以下服务编排：
1. `frontend`：基于 Nginx Alpine 的镜像，挂载 Vue 的 `dist` 构建打包产物。
2. `backend`：基于 Alpine 的自定义 Go 镜像，运行编译后的二进制文件。
3. `mysql`：官方 MySQL 镜像，配置卷持久化 (Volume Mapping)。
4. `redis`：官方 Redis 镜像。

---
