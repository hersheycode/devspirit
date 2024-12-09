FROM golang:1.18

# paths
ENV SHARED=pkg
ENV DEV=examples/prodapi/config/dev
ENV PKG=examples/prodapi/pkg

# certs
COPY $DEV/secrets/crt/auth /etc/ssl/certs/auth
COPY $DEV/secrets/crt/dgraph /etc/ssl/certs/dgraph

# netrc enables 'go get' to get private mod through https
COPY $DEV/secrets/netrc /root/.netrc  
RUN go env -w GOPRIVATE=apppathway.com/pkg/*

# prodapi modules: this needs to match go.mod replace directives
WORKDIR /workspaces/devspirit/examples/prodapi/pkg
COPY $PKG/plugins/scheduler ./plugins/scheduler

# apppathway shared modules: this needs to match go.mod replace directives
WORKDIR /workspaces/devspirit/pkg
COPY $SHARED/errors ./errors
COPY $SHARED/debug ./debug

# build 
WORKDIR /workspaces/devspirit/examples/prodapi/pkg/plugins/scheduler
RUN go mod tidy -compat=1.18

# runs with -race flag 
COPY $PKG/plugins/scheduler/cmd/schedulerd/devstart/run .

# reload run script when files are edited
RUN go get github.com/githubnemo/CompileDaemon 
RUN go build -o /go/bin/CompileDaemon /go/pkg/mod/github.com/githubnemo/!compile!daemon@v1.4.0 
ENTRYPOINT CompileDaemon -graceful-kill=true -command="./run" 

