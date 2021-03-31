编译命令：
`protoc -I=. --go_out=. ./type_enum.proto`

grpc编译命令：

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    grpc_req_res.proto
```
