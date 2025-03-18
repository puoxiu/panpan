# 本脚本记录在开发项目时所用到的代码生成等相关命令


# 用户服务
## api
goctl api go -api user.api -dir ./ -home ../../../tamplate
## model
goctl model mysql ddl -src="user.sql" -dir="./" --cache=true --prefix="user:"

## rpc
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -home ../../../tamplate



# 上传服务
## api
goctl api go -api upload.api -dir ./ -home ../../../tamplate

## model
goctl model mysql ddl -src="file.sql" -dir="./" --cache=true --prefix="upload:"

## rpc
goctl rpc protoc upload.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. -home ../../../tamplate