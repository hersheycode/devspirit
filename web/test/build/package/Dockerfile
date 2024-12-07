FROM golang:1.18rc1-bullseye

# COPY secrets/crt/auth /etc/ssl/certs
# COPY secrets/crt/dgraph /etc/ssl/certs/dgraph
# # netrc enables 'go get' to get private mod through https
# COPY secrets/netrc /root/.netrc 
# RUN go env -w GOPRIVATE=apppathway.com/pkg/*

WORKDIR /app
ENV PROJECT_DIR=web
COPY $PROJECT_DIR/test ./test

WORKDIR /app/test
RUN go mod tidy -compat=1.18

RUN go get github.com/githubnemo/CompileDaemon 
RUN go build -o /go/bin/CompileDaemon /go/pkg/mod/github.com/githubnemo/!compile!daemon@v1.4.0 
ENTRYPOINT CompileDaemon -graceful-kill=true -command="make reset" 
