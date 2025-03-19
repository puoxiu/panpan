# PanPan网盘项目





测试运行步骤：
* docker启动redis集群
* docker启动mysql
* 启动kafka
* 启动user-rpc：go run user.go

* docker启动minio集群 & Nginx
* 启动upload-rpc：go run upload.go



todo:
* user-rpc的sms接口申请
* 验证码申请是否需要限流（在rpc层面）