all:
	@echo "target"
	CGO_ENABLED=0 go build  -a -ldflags '-extldflags "-static"' -o out/ping cmd/main.go
