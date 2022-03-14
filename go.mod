module analytics

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/micro/micro/v3 v3.10.0
	github.com/micro/services v0.25.0
	google.golang.org/genproto v0.0.0-20220310185008-1973136f34c6 // indirect
	google.golang.org/protobuf v1.27.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
