FROM golang:buster

ENV GOPROXY="https://goproxy.cn" 

RUN apt-get update && \
    apt-get install nano iputils-ping telnet net-tools ifstat -y

RUN cp  /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  && \
    echo 'Asia/Shanghai'  > /etc/timezone

RUN go install github.com/yomorun/cli/yomo@v0.0.5; exit 0

WORKDIR $GOPATH/src/github.com/yomorun/yomo-flow-noise-example
COPY app.go go.mod ./

CMD ["sh", "-c", "yomo run app.go -m ./go.mod -u noise-zipper:9999 -n NoiseServerless"]
