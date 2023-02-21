cd $(dirname $0)
kitex -module github.com/joker-star-l/dousheng_idls/message -service MessageService message.thrift
go mod tidy
go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0
go mod tidy
