FROM oraclelinux:latest

 ENV LANG=C.UTF-8 \
     http_proxy=http://www-proxy.us.oracle.com:80 \
     https_proxy=http://www-proxy.us.oracle.com:80 \
     no_proxy="localhost,127.0.0.1,.oraclecorp.com,.grungy.us,docker"

 #RUN yum install -y gcc wget openssl-devel libffi-devel
 #RUN yum install -y openssl-devel redhat-lsb yum-utils rpmdevtools

ADD dist/bin/endpoint-monitor-bot  /endpoint-monitor-bot
EXPOSE 8080
ENTRYPOINT ["/endpoint-monnitor-bot"]
