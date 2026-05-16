# 个人网站数据库设计 (MySQL)

本文档定义了个人网站系统的关系型数据库结构。为了配合 Go 语言的 GORM 框架，表设计通常会包含 `created_at`, `updated_at` 以及用于软删除的 `deleted_at` 字段。

## 1. 实体关系模型 (ER 简述)
- **用户 (users)**：目前主要指后台管理员。
- **文章 (articles)**：每篇文章归属于一个分类 (Category)，并由一个用户 (User) 创建。
- **分类 (categories)**：一对多关联文章。
- **标签 (tags)**：多对多关联文章 (通过关联表 `article_tags`)。
- **工具 (tools)**：存放挂载在线工具的元数据。

---

## 2. DDL 建表语句

```sql
-- 创建数据库
CREATE DATABASE IF NOT EXISTS `personal_site` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `personal_site`;

-- 1. 用户表 (Admin 账户)
CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(64) NOT NULL COMMENT '用户名',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '密码的 bcrypt 哈希值',
  `email` VARCHAR(128) DEFAULT NULL COMMENT '邮箱',
  `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 2. 分类表
CREATE TABLE `categories` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) NOT NULL COMMENT '分类名称',
  `description` VARCHAR(255) DEFAULT '' COMMENT '分类描述',
  `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章分类表';

-- 3. 标签表
CREATE TABLE `tags` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(64) NOT NULL COMMENT '标签名称',
  `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章标签表';

-- 4. 文章表
CREATE TABLE `articles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '作者ID',
  `category_id` BIGINT UNSIGNED NOT NULL COMMENT '分类ID',
  `title` VARCHAR(255) NOT NULL COMMENT '文章标题',
  `summary` VARCHAR(512) DEFAULT '' COMMENT '文章摘要',
  `content` LONGTEXT NOT NULL COMMENT 'Markdown 格式正文',
  `view_count` INT UNSIGNED DEFAULT 0 COMMENT '阅读量',
  `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `deleted_at` DATETIME(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_category_id` (`category_id`),
  KEY `idx_user_id` (`user_id`),
  -- GORM 通常可以不设强制物理外键以提升性能，这里保留逻辑关联的外键约束
  CONSTRAINT `fk_article_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT `fk_article_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

-- 5. 文章-标签 多对多关联表
CREATE TABLE `article_tags` (
  `article_id` BIGINT UNSIGNED NOT NULL,
  `tag_id` BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (`article_id`, `tag_id`),
  CONSTRAINT `fk_at_article` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_at_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章与标签关联表';

-- 6. 工具元数据表 (用于动态渲染前端工具集)
CREATE TABLE `tools` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `tool_key` VARCHAR(64) NOT NULL COMMENT '工具的唯一英文标识 (如 json_formatter)',
  `name` VARCHAR(64) NOT NULL COMMENT '工具展示名称',
  `description` VARCHAR(255) DEFAULT '' COMMENT '工具描述',
  `icon` VARCHAR(128) DEFAULT '' COMMENT '前端图标名称或路径',
  `route` VARCHAR(128) NOT NULL COMMENT '前端路由路径',
  `is_active` TINYINT(1) DEFAULT 1 COMMENT '是否启用: 1是，0否',
  `created_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_at` DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tool_key` (`tool_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='在线工具元数据表';
```
