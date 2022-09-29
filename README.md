# PublicFileServer

使用 gin 实现临时文件共享服务后端

## 前端仓库

[PublicFile-Web](https://github.com/sxz799/PublicFile-Web)

## 编译命令

#### mac/linux
`go build -ldflags="-s -w" -o fileServer main.go`

#### windows
`go build -ldflags="-s -w" -o fileServer.exe main.go`