

## Build windows development version
upx: https://github.com/upx/upx/releases
```shell
go build -ldflags="-s -w" -o ya-daemon-go-amd64.exe && upx -9 ya-daemon-go-amd64.exe
```