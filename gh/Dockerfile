FROM golang:1.20-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

# 安装ca-certificates，发送HTTPS请求，否则会报错"x509: certificate signed by unknown authority"
RUN apk update && apk upgrade &&\
	apk add --no-cache ca-certificates && update-ca-certificates &&\
	#apk add --no-cache build-base=0.5-r2
	apk add --no-cache build-base


# 编译项目
COPY go.mod .
COPY go.sum .
#COPY public .
#COPY config.toml .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLE=1 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o rss2 .

# 二阶段编译
FROM alpine AS releaser
COPY --from=builder /build/rss2 /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /build/public /public
#COPY --from=builder /build/config.toml /config.toml

RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/shanghai" > /etc/timezone;

EXPOSE 8090
CMD ["/rss2"]
