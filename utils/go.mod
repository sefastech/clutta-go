module github.com/sefastech/clutta/libraries/golang/utils

go 1.22.1

require (
	github.com/google/uuid v1.6.0
	github.com/labstack/echo/v4 v4.12.0
	github.com/rs/zerolog v1.33.0
	github.com/sefastech/clutta/libraries/golang/coretypes v0.0.0-20240815210507-d22f8934628d
	github.com/sefastech/clutta/libraries/golang/grpc/clutta-sync v0.0.0-20240815210507-d22f8934628d
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/sefastech/clutta/libraries/golang/utils => ../../libraries/golang/utils
