module apppathway.com/examples/prodapi/pkg/plugins/scheduler

replace apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb => /home/nate/code/app-pathway/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb

replace apppathway.com/pkg/errors => /home/nate/code/app-pathway/pkg/errors

replace apppathway.com/pkg/debug => /home/nate/code/app-pathway/pkg/debug

go 1.18

require (
	github.com/dgraph-io/dgo/v200 v200.0.0-20210401091508-95bfd74de60e
	github.com/dgraph-io/ristretto v0.1.0
	google.golang.org/grpc v1.45.0
)

require (
	apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/pkg/errors v0.0.0-00010101000000-000000000000 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
