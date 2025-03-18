# 本脚本记录在开发项目时所用到的代码生成等相关命令


# 用户服务
## api

## model
goctl model mysql ddl -src="user.sql" -dir="./" --cache=true --prefix="user:"

## rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -home ../../../tamplate
