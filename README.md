# heegproto

## gen proto
```
protoc --proto_path=$GOPATH/src/heegproto:. --micro_out=. --go_out=. common.proto
```

## 添加自定义tag
```
go get -u github.com/favadi/protoc-go-inject-tag
go install github.com/favadi/protoc-go-inject-tag@latest

message question {
    // @inject_tag: form:"timu"
    string                  timu = 1;
    repeated string          images = 2;
    repeated question_option options = 3;
}

执行：
    protoc-go-inject-tag -input=./common.pb.go
```