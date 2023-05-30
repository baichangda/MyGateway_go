# MyGateway_go

- 跨平台编译  
  CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target