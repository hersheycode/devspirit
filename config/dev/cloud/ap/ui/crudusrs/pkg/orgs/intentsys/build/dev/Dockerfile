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
WORKDIR /home/nate/code/app-pathway/examples/prodapi/pkg
COPY $PKG/orgs/intentsys ./orgs/intentsys
COPY $PKG/plugins/intent ./plugins/intent
COPY $PKG/plugins/scheduler ./plugins/scheduler
COPY $PKG/plugins/sms ./plugins/sms

# apppathway shared modules: this needs to match go.mod replace directives
WORKDIR /home/nate/code/app-pathway/pkg
COPY $SHARED/net ./net
COPY $SHARED/errors ./errors
COPY $SHARED/debug ./debug

# build 
WORKDIR /home/nate/code/app-pathway/examples/prodapi/pkg/orgs/intentsys
RUN go mod tidy -compat=1.18

# runs with -race flag 
COPY $PKG/orgs/intentsys/cmd/intentsysd/devstart/run .

# reload run script when files are edited
RUN go get github.com/githubnemo/CompileDaemon 
RUN go build -o /go/bin/CompileDaemon /go/pkg/mod/github.com/githubnemo/!compile!daemon@v1.4.0 
ENTRYPOINT CompileDaemon -graceful-kill=true -command="./run" 

