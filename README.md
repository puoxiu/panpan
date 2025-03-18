# PanPan网盘项目



## user

rpc : 8081
api : 8961



测试运行步骤：
* docker启动redis
* docker启动mysql
* 启动user-rpc：go run user.go

todo:
* user-rpc的sms接口申请
* 验证码申请是否需要限流（在rpc层面）