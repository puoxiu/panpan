# 📦 PanPan Cloud - 微服务网盘系统

## 🚀 项目简介

> PanPan Cloud 是一个支持多存储后端的分布式网盘系统，提供文件存储、秒传、分片上传 / 下载、异步任务处理等核心功能。支持本地磁盘、腾讯云 COS、Minio 集群三种存储方式，通过批处理队列（Kafka）实现数据库操作的异步化处理，保障高并发场景下的系统稳定性。

## 🌟 核心功能特性

🔥 文件存储（三模支持）
 <br>
✅ 多存储后端：
* 本地磁盘（服务器文件系统）
* 腾讯云 COS（对象存储）
* Minio 集群（私有云对象存储）

✅ 上传能力：
* 🧩 分片上传（支持断点续传）
* ✨ 秒传（基于 SHA-1 指纹快速校验）
* 📏 自定义分片大小（开发环境 100MB，可调）

✅ 下载能力：
* ⚡ 分块并行下载（Goroutine 池并发控制）
* 🛡️ 文件完整性校验（SHA-1 哈希比对）
* 🚀 流式传输（零拷贝优化）


## 🔄 关键流程
### 1. 分片上传流程：

```mermaid
sequenceDiagram
  客户端->>API: 初始化分片（SHA-1, 大小）
  API->>存储层: 生成UploadID（Redis）
  API-->>客户端: 返回UploadID, 分片信息
  客户端->>API: 分片上传（含Index）
  API->>存储层: 校验分片完整性
  客户端->>API: 完成上传（合并请求）
  API->>存储层: 合并分片+SHA-1校验
  API->>批处理: 记录存储元数据（Kafka）
```
<br>
<br>
<br>


### 2. 批处理流程

```mermaid
graph LR
  业务逻辑 -->|数据库操作| 环形缓冲区
  环形缓冲区 -->|阈值/时间触发| Kafka生产者
  Kafka -->|消费者组| 数据库Worker
  数据库Worker -->|异步写入| MySQL/PostgreSQL

```

## 🛠️ 技术栈
Golang、Go-Zero、gRPC、MySQL、Redis、etcd、Docker、Kafka、Jaeger、MinIO、COS

## 📥 安装与部署
> 依据docker-compose文件


## 📌 todo
1. 用户文件分享、连接分享
2. 文件删除、更改
3. 系统监控系统，微服务监控
4. 限速功能
5. 。。。


