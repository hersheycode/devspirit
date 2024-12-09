FROM golang:1.18

COPY secrets/crt/auth /etc/ssl/certs/auth
COPY secrets/crt/dgraph /etc/ssl/certs/dgraph
# netrc enables 'go get' to get private mod through https
COPY secrets/netrc /root/.netrc 
RUN go env -w GOPRIVATE=apppathway.com/pkg/*

WORKDIR /workspaces/devspirit/pkg
ENV PROJECT_DIR=pkg
COPY $PROJECT_DIR/builder/base ./builder/base
COPY $PROJECT_DIR/user ./user
COPY $PROJECT_DIR/net ./net
COPY $PROJECT_DIR/errors ./errors
COPY $PROJECT_DIR/debug ./debug

WORKDIR /workspaces/devspirit/pkg/builder/base
RUN go mod tidy -compat=1.18

# DETECT RACE
COPY $PROJECT_DIR/builder/base/cmd/cplugind/devstart/run .
RUN go get github.com/githubnemo/CompileDaemon 
RUN go build -o /go/bin/CompileDaemon /go/pkg/mod/github.com/githubnemo/!compile!daemon@v1.4.0 
RUN go install golang.org/x/tools/...@v0.1.10
# BUILD DEFAULT is just go build AND its requires?
ENTRYPOINT CompileDaemon -graceful-kill=true -command="./run" 
