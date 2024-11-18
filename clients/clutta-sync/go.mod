module github.com/sefastech/clutta-go/clients/clutta-sync

go 1.22.7

toolchain go1.22.9

require github.com/google/uuid v1.6.0

require github.com/sirupsen/logrus v1.9.3 // indirect

require (
	github.com/sefastech/clutta-go/grpc/clutta-sync v0.0.0-20241118153553-d29f7d1b1edb
	github.com/sefastech/clutta/libraries/golang/logging v0.0.0-20241114044337-6d9951a7d298
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241113202542-65e8d215514f // indirect
	google.golang.org/grpc v1.68.0 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
