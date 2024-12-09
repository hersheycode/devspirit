FROM golang:1.18rc1-bullseye

COPY secrets/crt/auth /etc/ssl/certs
# netrc enables 'go get' to get private mod through https
COPY secrets/netrc /root/.netrc 
RUN go env -w GOPRIVATE=apppathway.com/pkg/*

WORKDIR /workspaces/devspirit/pkg/user/auth
ENV ROOT=pkg/user
COPY $ROOT/auth/internals ./internals
COPY $ROOT/auth/go.mod .
COPY $ROOT/auth/cmd/server/main.go . 

WORKDIR /workspaces/devspirit/pkg
ENV PROJECT_DIR=pkg
COPY $PROJECT_DIR/db_api ./db_api
COPY $PROJECT_DIR/net ./net
COPY $PROJECT_DIR/errors ./errors
COPY $PROJECT_DIR/debug ./debug

WORKDIR /workspaces/devspirit/pkg/user/auth
RUN go mod tidy -compat=1.18

RUN go get github.com/githubnemo/CompileDaemon 
RUN go build -o /go/bin/CompileDaemon /go/pkg/mod/github.com/githubnemo/!compile!daemon@v1.4.0 

ENTRYPOINT CompileDaemon --build="go build main.go" --command='./main'

