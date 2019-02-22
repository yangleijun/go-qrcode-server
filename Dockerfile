#从空镜像中构建
FROM scratch

#作者
MAINTAINER yangleijun "https://github.com/yangleijun"

#变量
ENV GOPATH=/apps/go

#设置工作目录
WORKDIR /apps/go/src/yangleijun/go-qrcode-server

#复制当前目录的所有文件

ADD . /apps/go/src/yangleijun/go-qrcode-server

#暴露端口
EXPOSE 80

#启动二维码服务
ENTRYPOINT ["/apps/go/src/yangleijun/go-qrcode-server/main"]