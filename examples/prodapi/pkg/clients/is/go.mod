module apppathway.com/examples/prodapi/pkg/clients/is

replace apppathway.com/examples/prodapi/pkg/orgs/intentsys => /workspaces/devspirit/examples/prodapi/pkg/orgs/intentsys

replace apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb => /workspaces/devspirit/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb

replace apppathway.com/examples/prodapi/pkg/plugins/intent/api/intentpb => /workspaces/devspirit/examples/prodapi/pkg/plugins/intent/api/intentpb

replace apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb => /workspaces/devspirit/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb

replace apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb => /workspaces/devspirit/examples/prodapi/pkg/plugins/sms/api/smspb

replace apppathway.com/pkg/user/auth => /workspaces/devspirit/pkg/user/auth

replace apppathway.com/pkg/net => /workspaces/devspirit/pkg/net

replace apppathway.com/pkg/errors => /workspaces/devspirit/pkg/errors

replace apppathway.com/pkg/debug => /workspaces/devspirit/pkg/debug

go 1.18

require (
	apppathway.com/examples/prodapi/pkg/orgs/intentsys v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.4.0
	github.com/spf13/viper v1.10.1
)

require (
	apppathway.com/examples/prodapi/pkg/orgs/intentsys/api/intentsyspb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/examples/prodapi/pkg/plugins/intent/api/intentpb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/examples/prodapi/pkg/plugins/scheduler/api/schedulerpb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/examples/prodapi/pkg/plugins/sms/api/smspb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582 // indirect
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
