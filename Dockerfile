FROM golang:alpine as builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#tzdata 在 builder 镜像安装，并在最终镜像只拷贝了需要的时区
#自动设置了本地时区，这样我们在日志里看到的是北京时间了
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -ldflags="-s -w " -gcflags "-N -l" -installsuffix cgo -o qcxx-admin-server main.go

FROM alpine as prod

# 在build阶段复制时区到
RUN apk update --no-cache && apk add --no-cache ca-certificates
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
# 在build阶段复制可执行的go二进制文件app
COPY --from=builder /app/qcxx-admin-server /app
# 在build阶段复制配置文件
COPY --from=builder /app/config /app/config

EXPOSE 8081

ENTRYPOINT ["./qcxx-admin-server","-f","dev.yml"]