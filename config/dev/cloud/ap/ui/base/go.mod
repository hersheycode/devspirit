module apppathway.com/pkg/builder/base

replace apppathway.com/pkg/user/auth => /workspaces/devspirit/pkg/user/auth

replace apppathway.com/pkg/net => /workspaces/devspirit/pkg/net

replace apppathway.com/pkg/errors => /workspaces/devspirit/pkg/errors

replace apppathway.com/pkg/debug => /workspaces/devspirit/pkg/debug

replace apppathway.com/pkg/builder/base/api/cpluginpb => /workspaces/devspirit/pkg/builder/base/api/cpluginpb

go 1.18

require (
	apppathway.com/pkg/builder/base/api/cpluginpb v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000
	github.com/dgraph-io/dgo/v200 v200.0.0-20210401091508-95bfd74de60e
	github.com/dgraph-io/ristretto v0.1.0
	github.com/google/goterm v0.0.0-20200907032337-555d40f16ae2
	github.com/spf13/cobra v1.4.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211222154725-9823f7ba7562 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
