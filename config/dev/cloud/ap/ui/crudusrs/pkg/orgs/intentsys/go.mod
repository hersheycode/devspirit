module apppathway.com/examples/prodapi/pkg/orgs/intentsys

// replace apppathway.com/pkg/user/auth => /workspaces/devspirit/pkg/user/auth

replace apppathway.com/pkg/net => /workspaces/devspirit/pkg/net

replace apppathway.com/pkg/errors => /workspaces/devspirit/pkg/errors

replace apppathway.com/pkg/debug => /workspaces/devspirit/pkg/debug

replace apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb => /workspaces/devspirit/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb

replace codestore.localhost/crudusrs/crud_basic/api/crudbasic/api/intentpb => /workspaces/devspirit/examples/prodapi/pkg/plugins/intent/api/intentpb

replace apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb => /workspaces/devspirit/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb

replace apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb => /workspaces/devspirit/examples/prodapi/pkg/plugins/sms/api/smspb

go 1.18

require (
	apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb v0.0.0-00010101000000-000000000000
	codestore.localhost/crudusrs/crud_basic/api/crudbasic/api/intentpb v0.0.0-00010101000000-000000000000
	apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb v0.0.0-00010101000000-000000000000
	apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000
	github.com/dgraph-io/ristretto v0.1.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
