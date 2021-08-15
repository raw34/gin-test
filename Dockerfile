# Step1 程序编译
# 拉取 Go 语言最新的基础镜像
FROM golang:1.13-alpine AS build

# 镜像加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装依赖
RUN apk add --update --no-cache gcc g++ musl-dev

# 设置当前工作目录
WORKDIR /data/build

# 把文件复制到当前工作目录
COPY . .

# 设置 GOPROXY 环境变量
ENV GOOS=linux
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 下载全部依赖项
RUN go mod download

# 编译项目
RUN go build -o gin-test .

# Step2 程序运行
# 安装运行环境
FROM golang:1.13-alpine AS runtime
#FROM alpine:3.12  AS runtime

# 镜像加速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 设置当前工作目录
WORKDIR /data/app

# 创建日志目录
RUN  mkdir /data/log

# 设置时区
RUN apk add --update --no-cache tzdata \
    && echo "Asia/Shanghai" > /etc/timezone

# 拷贝已编译程序
COPY --from=build /go /go
COPY --from=build /data/build /data/build
COPY --from=build /data/build/gin-test /data/app

# 暴露端口
EXPOSE 28009

# 运行
CMD ["./gin-test", "gin"]
