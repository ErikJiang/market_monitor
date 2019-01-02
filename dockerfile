# 指定 GO 版本号
ARG GO_VERSION=1.11.1

# 指定构建环境
FROM golang:${GO_VERSION}-alpine AS builder

# 为 go build 设置环境变量:
# * CGO_ENABLED=0 表示构建一个静态链接的可执行程序
# * GOFLAGS=-mod=vendor 在执行 `go build` 强制查看 `/vendor` 目录.
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

# 设置工作目录
WORKDIR /src

# 拷贝所有源码文件
COPY ./ ./

# 构建可执行文件
RUN go build -installsuffix 'static' -o /go/bin/monitor


# 构建最小镜像
FROM alpine AS final

# 设置系统语言
ENV LANG en_US.UTF-8

# 将构建的可执行文件复制到新镜像中
RUN mkdir -p /app/config/
COPY --from=builder /src/config/config.yaml /app/config/config.yaml
COPY --from=builder /go/bin/monitor /app/monitor

# 端口申明
EXPOSE 8000

# 运行
ENTRYPOINT [ "/app/monitor" ]
