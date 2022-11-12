# ------------------------------------ 编译开发环境可调试镜像 ---------------------------------------------
FROM golang:latest

ENV GO111MODULE=on \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR /matchAPP

## go module拉取更新
COPY ./ /matchAPP
RUN go mod download
RUN go mod tidy

EXPOSE 8080


ENTRYPOINT go run main.go
