# 个人网站 API 接口文档

## 1. 基础说明

- **基础路径 (Base URL):** `/api/v1`
- **数据交互格式:** 所有请求和响应的数据格式均为 `application/json`。
- **认证方式:** 部分需授权的接口（如后台管理）需在请求头携带 JWT Token。
  - 格式：`Authorization: Bearer <your_jwt_token>`

### 1.1 通用响应格式

所有接口默认遵循以下 JSON 结构返回数据：

```json
{
  "code": 200,           // 状态码：200成功，40x客户端错误，50x服务端错误
  "message": "success",  // 提示信息
  "data": {}             // 实际数据承载体（对象或数组）
}
```

---

## 2. 认证模块 (Auth)

### 2.1 管理员登录
- **路径:** `POST /api/v1/auth/login`
- **描述:** 管理员登录并获取 JWT Token。
- **请求体 (Body):**
  ```json
  {
    "username": "admin",
    "password": "your_password"
  }
  ```
- **响应体 (Response):**
  ```json
  {
    "code": 200,
    "message": "登录成功",
    "data": {
      "token": "eyJhbGciOiJIUzI1NiIsInR5c... (JWT String)",
      "expire_at": 1700000000
    }
  }
  ```

---

## 3. 博客模块 (Blog)

### 3.1 获取文章列表
- **路径:** `GET /api/v1/articles`
- **描述:** 获取博客文章列表，支持分页和分类/标签过滤。
- **请求参数 (Query):**
  - `page` (int, 默认 1): 当前页码
  - `size` (int, 默认 10): 每页数量
  - `category_id` (int, 可选): 分类ID
  - `tag_id` (int, 可选): 标签ID
- **响应体 (Response):**
  ```json
  {
    "code": 200,
    "message": "success",
    "data": {
      "total": 50,
      "page": 1,
      "size": 10,
      "list": [
        {
          "id": 1,
          "title": "如何使用 Go 构建高并发系统",
          "summary": "文章简介...",
          "category": { "id": 1, "name": "后端开发" },
          "tags": ["Go", "并发"],
          "created_at": "2023-10-01T12:00:00Z",
          "view_count": 120
        }
      ]
    }
  }
  ```

### 3.2 获取文章详情
- **路径:** `GET /api/v1/articles/:id`
- **描述:** 根据文章 ID 获取正文详情内容（Markdown）。
- **响应体 (Data):**
  ```json
  {
    "id": 1,
    "title": "如何使用 Go 构建高并发系统",
    "content": "这里是完整的 Markdown 格式正文...",
    "category_name": "后端开发",
    "tags": ["Go", "并发"],
    "created_at": "2023-10-01T12:00:00Z"
  }
  ```

### 3.3 创建文章 (需要 JWT)
- **路径:** `POST /api/v1/articles`
- **描述:** 发布新文章。
- **请求体 (Body):**
  ```json
  {
    "title": "新文章标题",
    "summary": "这是摘要",
    "content": "这里是 Markdown 内容",
    "category_id": 1,
    "tag_ids": [1, 2]
  }
  ```

### 3.4 修改文章 (需要 JWT)
- **路径:** `PUT /api/v1/articles/:id`
- **描述:** 更新已存在的博客文章。

### 3.5 删除文章 (需要 JWT)
- **路径:** `DELETE /api/v1/articles/:id`
- **描述:** 删除现有文章（或逻辑删除）。

---

## 4. 分类与标签模块 (Categories & Tags)

### 4.1 获取所有分类
- **路径:** `GET /api/v1/categories`
- **响应体 (Data):**
  ```json
  [
    { "id": 1, "name": "后端开发", "article_count": 15 },
    { "id": 2, "name": "生活随笔", "article_count": 5 }
  ]
  ```

### 4.2 获取所有标签
- **路径:** `GET /api/v1/tags`
- **响应体 (Data):**
  ```json
  [
    { "id": 1, "name": "Go", "article_count": 10 },
    { "id": 2, "name": "Vue3", "article_count": 8 }
  ]
  ```

---

## 5. 在线工具模块 (Tools)

### 5.1 获取工具列表
- **路径:** `GET /api/v1/tools`
- **描述:** 获取当前系统支持的在线工具/Block列表，用于填充前端工具中心页面。
- **响应体 (Data):**
  ```json
  [
    {
      "id": "json_formatter",
      "name": "JSON 格式化工具",
      "description": "提供 JSON 字符串的美化、压缩与语法检查",
      "icon": "icon-json",
      "route": "/tools/json-formatter"
    },
    {
      "id": "base64_codec",
      "name": "Base64 编码/解码",
      "description": "支持文本和文件的 Base64 互相转换",
      "icon": "icon-base64",
      "route": "/tools/base64"
    }
  ]
  ```

*(注：具体的工具如果需要后端计算支持，可以在对应的插件模块中单独扩展接口，如 `POST /api/v1/tools/execute/:tool_id`)*
