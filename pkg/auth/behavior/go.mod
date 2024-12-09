module apppathway.com/pkg/user/behavior

replace apppathway.com/pkg/user/auth => /workspaces/devspirit/pkg/user/auth

replace apppathway.com/pkg/net => /workspaces/devspirit/pkg/net

replace apppathway.com/pkg/errors => /workspaces/devspirit/pkg/errors

replace apppathway.com/pkg/debug => /workspaces/devspirit/pkg/debug

replace apppathway.com/pkg/user/behavior/api/behaviorpb => /workspaces/devspirit/pkg/user/behavior/api/behaviorpb

go 1.18

require (
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/user/behavior/api/behaviorpb v0.0.0-00010101000000-000000000000
	github.com/dgraph-io/dgo/v200 v200.0.0-20210401091508-95bfd74de60e
	github.com/dgraph-io/ristretto v0.1.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211222154725-9823f7ba7562 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
