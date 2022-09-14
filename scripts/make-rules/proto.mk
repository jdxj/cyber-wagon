.PHONY: proto
proto: proto.gateway proto.user


.PHONY: proto.%
proto.%:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/$*/proto/*.proto
