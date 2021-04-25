FROM golang:buster

RUN apt-get update && \
    apt-get install nano iputils-ping telnet net-tools ifstat -y

RUN cp  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  && \
    echo 'Asia/Shanghai'  > /etc/timezone

RUN GO111MODULE=off go get github.com/yomorun/yomo; exit 0
RUN cd $GOPATH/src/github.com/yomorun/yomo && make install

WORKDIR $GOPATH/src/github.com/yomorun/yomo-flow-noise-example
COPY app.go go.mod ./
RUN echo '\r\n replace github.com/yomorun/yomo v1.0.0 => ../yomo' >> go.mod
RUN go get -d -v ./...

EXPOSE 4242/udp

CMD ["sh", "-c", "yomo run app.go -p 4242"]