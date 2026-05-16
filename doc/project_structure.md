# PersonalSite 项目代码规范与目录结构说明

本文档基于生成的基础脚手架代码，约定项目的整体目录结构设计。

## 整体架构

```text
PersonalSite/
├── backend/                  # Go 后端服务层
│   ├── cmd/                  # 包含服务启动的 main 函数 (例如：cmd/api/main.go)
│   ├── internal/             # 私有业务逻辑层 (Clean Architecture)
│   │   ├── handler/          # HTTP 控制器 (解析请求，调用 service，返回响应)
│   │   ├── service/          # 核心业务逻辑实现层
│   │   ├── repository/       # 数据访问层 (GORM 数据库交互)
│   │   ├── model/            # 数据库结构体和各种实体类映射 (对应 database_schema)
│   │   ├── middleware/       # Gin 中间件 (Auth、CORS、限流等)
│   ├── pkg/                  # 公共库 (如 JWT 生成、Redis 客户端封装、工具函数)
│   ├── go.mod                # Go 模块配置文件
│   ├── main.go               # 快速预览版入口文件 (后续完善将移入 cmd 目录)
│   └── Dockerfile            # 后端镜像构建脚本
├── frontend/                 # Vue3 前端 UI 层
│   ├── public/               # 公开静态资源
│   ├── src/                  # 前端源代码
│   │   ├── api/              # Axios 请求封装集合
│   │   ├── assets/           # CSS 样式和本地图片静态资源
│   │   ├── components/       # 通用的 Vue 组件 (如 Navigation Bar)
│   │   ├── views/            # 路由视图文件 (首页、博文列表、工具等)
│   │   ├── router/           # Vue Router 配置
│   │   ├── store/            # Pinia 状态管理
│   │   ├── App.vue           # 根组件
│   │   └── main.js           # Vue 挂载入口
│   ├── package.json          # Node 依赖与脚本
│   ├── vite.config.js        # Vite 构建配置
│   ├── tailwind.config.js    # Tailwind 配置文件
│   └── Dockerfile            # 前端 Nginx 构建脚本
├── doc/                      # 项目设计与说明文档
└── docker-compose.yml        # Docker 容器编排入口
```
