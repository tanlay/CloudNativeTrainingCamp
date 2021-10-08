#打包阶段使用golang:alpine作为基础镜像
FROM golang:alpine as build-env
#制定打包者信息
MAINTAINER tanlay

#创建工作目录
RUN mkdir /data
#切换到到工作目录
WORKDIR /data 

#复制源代码到工作目录
COPY simplehttpserver.go .
#执行go build
RUN go build simplehttpserver.go

#制定运行阶段的基础镜像
FROM alpine
WORKDIR /root
#将上一阶段的文件复制进来
COPY --from=build-env /data/simplehttpserver . 
#暴露端口
EXPOSE 8099
#执行程序
CMD ["./simplehttpserver"]
