# 第二次作业
- 构建本地镜像。
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化（请思考有哪些最佳实践可以引入到 Dockerfile 中来）。
- 将镜像推送至 Docker 官方镜像仓库。
- 通过 Docker 命令本地启动 httpserver。
- 通过 nsenter 进入容器查看 IP 配置

## Dockerfile构建镜像

### 编写Dockerfile

```dockerfile
#打包阶段使用golang:alpine作为基础镜像
FROM golang:alpine as build-env
#指定打包者信息
MAINTAINER tanlay

#创建工作目录
RUN mkdir /data
#切换到到工作目录
WORKDIR /data 

#复制源代码到工作目录
COPY simplehttpserver.go .
#执行go build
RUN go build simplehttpserver.go

#指定运行阶段的基础镜像
FROM alpine
RUN mkdir /data
WORKDIR /data
#将上一阶段的文件复制进来
COPY --from=build-env /data/simplehttpserver . 
#暴露端口
EXPOSE 8099
#执行程序
#CMD ["./simplehttpserver"]
ENTRYPOINT ["./simplehttpserver"]
```

### docker build构建镜像
```shell
$ docker build -t simplehttpserver:v1 . 
```

### 编写Dockerfile的最佳实践
- 精简上下文
	- 每次构建,上下文都会被复制给docker deamon,故需要精简centext
- 尽可能利用构建缓存
	- 按照镜像层的变动频率编写Dockerfile，尽可能把不变的层写至内层
- 减少层
- 使用多级构建
	- 构建使用一个基础镜像，运行使用最小化镜像


### 推送镜像至dockerhub

### 登录dockerhub

```shell
$ docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: xxxxx      
Password: 
WARNING! Your password will be stored unencrypted in /home/chenrui/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```
### tag修改镜像名称

```shell
$ docker tag simplehttpserver:v1 tanlay/simplehttpserver:v1
```

### 推送镜像

```shell
$ docker push tanlay/simplehttpserver:v1
The push refers to repository [docker.io/tanlay/simplehttpserver]
0c9cf67fdfcb: Pushed 
b36ae715e4d2: Pushed 
e2eb06d8af82: Pushed 
v1: digest: sha256:c81d095b92885189b28b434b2efa30ead76cd63e2ec13f064aefc35907944e82 size: 946
```

### 启动docker容器
```shell
$ docker run -itd simplehttpserver:v1
```

## 查看容器IP

### 查看httpserver容器运行在宿主机上PID
```shell
# 查看容器ID
$ docker ps 
CONTAINER ID   IMAGE                 COMMAND                CREATED         STATUS         PORTS      NAMES
848355f32b65   simplehttpserver:v1   "./simplehttpserver"   7 minutes ago   Up 7 minutes   8099/tcp   beautiful_brattain
# 通过容器ID查看在宿主机上的pid
$ docker inspect -f {{.State.Pid}}  848355f32b65
2197
```

### 通过PID查询容器IP
```shell
$ sudo nsenter -n -t 2197 hostname -I
172.17.0.2 
```

### 测试接口
进入docker查看容器IP地址为172.17.0.2
```shell
$ curl 172.17.0.2:8099/header -I
$ curl 172.17.0.2:8099/version
$ curl 172.17.0.2:8099/log
$ curl 172.17.0.2:8099/healthz
```

