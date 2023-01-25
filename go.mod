module github.com/vijaynallagatla/cli-stocks

go 1.19

require github.com/vijaynallagatla/cli-stocks/yapi v0.0.0

replace github.com/vijaynallagatla/cli-stocks/yapi v0.0.0 => ./yapi

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
