module apppathway.com/pkg/user/auth

replace apppathway.com/pkg/db_api => /workspaces/devspirit/pkg/db_api

replace apppathway.com/pkg/net => /workspaces/devspirit/pkg/net

replace apppathway.com/pkg/errors => /workspaces/devspirit/pkg/errors

replace apppathway.com/pkg/debug => /workspaces/devspirit/pkg/debug

replace apppathway.com/pkg/user => /workspaces/devspirit/pkg/user

go 1.18

require (
	apppathway.com/pkg/db_api v0.0.0-20220311005843-5cfd8d3bb370
	apppathway.com/pkg/errors v0.0.0-20220311014640-17934a361582
	apppathway.com/pkg/net v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gogf/gf v1.16.6
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/dgraph-io/dgo/v200 v200.0.0-20210401091508-95bfd74de60e // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/cobra v1.4.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211222154725-9823f7ba7562 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
