# 文件下载


## 从云COS服务下载
* 使用腾讯云 COS SDK 的FGetObject（或类似接口）直连下载


## 从MinIO集群下载

* 基于 Minio 的多部分下载 API（GetObjectOptions.PartNumber）
* 分块大小固定 100MB（通过chunkSize常量配置）
* 并发控制：使用sync.WaitGroup管理 goroutine 池

## 文件完整性校验
下载后立即计算 SHA1校验

