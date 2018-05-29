#ENV LANG=C.UTF-8 \
#    http_proxy=http://www-proxy.us.oracle.com:80 \
#    https_proxy=http://www-proxy.us.oracle.com:80 \
#    no_proxy="localhost,127.0.0.1,.oraclecorp.com,.grungy.us,docker"


# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ADD . /go/src/github.com/madhukirans/endpoint-monitor-bot

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/madhukirans/endpoint-monitor-bot

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/endpoint-monitor-bot

# Document that the service listens on port 8080.
EXPOSE 8080
