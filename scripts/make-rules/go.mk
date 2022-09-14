.PHONY: go.build
go.build: go.build.user

.PHONY: go.build.%
go.build.%:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o cmd/$*/$*.out cmd/$*/*.go
