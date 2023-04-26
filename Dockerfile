# 由 Dockerpacks 自动生成
# 本 Dockerfile 可能不能完全覆盖您的项目需求，若遇到问题请根据实际情况修改或询问客服

# 使用基于 alpine 的 golang 官方镜像
FROM golang:1-alpine

# 设置容器内的当前目录
WORKDIR /app

# 将包括源文件在内的所有文件拷贝到容器中
COPY . .

# 使用速度更快的国内镜像
ENV GOPROXY=https://mirrors.tencent.com/go

RUN go mod download

RUN go build -o wxcloudrun-files/main

CMD ["wxcloudrun-files/main"]

# 服务暴露的端口
EXPOSE 8080