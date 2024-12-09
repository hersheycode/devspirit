module apppathway.com/pkg/client

replace apppathway.com/pkg/user/auth => /home/nate/code/app-pathway/pkg/user/auth

replace apppathway.com/pkg/user/behavior => /home/nate/code/app-pathway/pkg/user/behavior

replace apppathway.com/pkg/user/behavior/api/behaviorpb => /home/nate/code/app-pathway/pkg/user/behavior/api/behaviorpb

replace apppathway.com/pkg/builder/gonodes => /home/nate/code/app-pathway/pkg/builder/gonodes

replace apppathway.com/pkg/builder/cplugin => /home/nate/code/app-pathway/pkg/builder/cplugin

replace apppathway.com/pkg/builder/ci => /home/nate/code/app-pathway/pkg/builder/ci

replace apppathway.com/pkg/errors => /home/nate/code/app-pathway/pkg/errors

replace apppathway.com/pkg/debug => /home/nate/code/app-pathway/pkg/debug

replace apppathway.com/pkg/net => /home/nate/code/app-pathway/pkg/net

replace apppathway.com/pkg/builder/gonodes/api/nodespb => /home/nate/code/app-pathway/pkg/builder/gonodes/api/nodespb

replace apppathway.com/pkg/builder/cplugin/api/cpluginpb => /home/nate/code/app-pathway/pkg/builder/cplugin/api/cpluginpb

replace apppathway.com/pkg/builder/ci/api/cipb => /home/nate/code/app-pathway/pkg/builder/ci/api/cipb

replace apppathway.com/pkg/builder/cd/api/templatepb => /home/nate/code/app-pathway/pkg/builder/cd/api/templatepb

replace apppathway.com/pkg/builder/cd => /home/nate/code/app-pathway/pkg/builder/cd

go 1.18

require (
	apppathway.com/pkg/builder/ci v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/builder/cplugin v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/builder/gonodes v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000
	apppathway.com/pkg/user/behavior v0.0.0-00010101000000-000000000000
	github.com/fatih/color v1.13.0
	github.com/google/goterm v0.0.0-20200907032337-555d40f16ae2
	github.com/spf13/cobra v1.4.0
	github.com/zyedidia/highlight v0.0.0-20200217010119-291680feaca1
)

require (
	apppathway.com/pkg/builder/ci/api/cipb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/pkg/builder/cplugin/api/cpluginpb v0.0.0-00010101000000-000000000000 // indirect
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582 // indirect
	apppathway.com/pkg/user/behavior/api/behaviorpb v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20220407224826-aac1ed45d8e3 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.8-0.20211004125949-5bd84dd9b33b // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
