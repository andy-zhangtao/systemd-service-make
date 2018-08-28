FROM vikings/golang-1.10
COPY gpl /go/src/github.com/andy-zhangtao/systemd-service-make/gpl
COPY module /go/src/github.com/andy-zhangtao/systemd-service-make/module
COPY tools /go/src/github.com/andy-zhangtao/systemd-service-make/tools
COPY main.go /go/src/github.com/andy-zhangtao/systemd-service-make/main.go

COPY unit-test.sh /go/src/github.com/andy-zhangtao/systemd-service-make/unit-test.sh
COPY build.sh /go/src/github.com/andy-zhangtao/systemd-service-make/build.sh
ENTRYPOINT ["sh"]
