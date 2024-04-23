module kang-blogging/internal/blogging

go 1.20

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/genproto/googleapis/api v0.0.0-20240412170617-26222e5d3d56
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.33.0
	kang-blogging/internal/common v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
)

replace kang-blogging/internal/common => ../common/
