.PHONY: generate tidy build test run clean

# 从 proto 生成 Go 代码
generate:
	buf generate

# 整理依赖
tidy:
	go mod tidy

# 编译
build: generate tidy
	go build -o bin/server ./cmd/server
	go build -o bin/client ./cmd/client

# 跑测试
test:
	go test ./...

# 启动服务器
run: build
	./bin/server

# 清理
clean:
	rm -rf bin/
