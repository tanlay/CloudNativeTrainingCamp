# Go实现简单web服务器

## 一共实现4个接口

- /header
  - 接收客户端 request，并将 request 中带的 header 写入 response header
- /version
  - 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
- /log
  - Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
- /healthz
  - 当访问 localhost/healthz 时，应返回200

## 用法

### 访问/header

```shell
> curl 172.17.0.2:8099/header -H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36' -I
> curl 172.17.0.2:8099/header
> curl 172.17.0.2:8099/header?q=query
```

### 访问/version

```shell
> curl 172.17.0.2:8099/version
GoVersion: go1.17.1

> curl 172.17.0.2:8099/version -I
HTTP/1.1 200 OK
Goversion: go1.17.1
Date: Wed, 29 Sep 2021 09:19:50 GMT
Content-Length: 19
Content-Type: text/plain; charset=utf-8
```

### 访问/log

```shell
> curl 172.17.0.2:8099/log
```

### 访问/healthz

```shell
> curl 172.17.0.2:8099/healthz
200

> curl 172.17.0.2:8099/healthz -I
HTTP/1.1 200 OK
Date: Wed, 29 Sep 2021 09:21:15 GMT
Content-Length: 3
Content-Type: text/plain; charset=utf-8
```

## 使用Dockerfile打包程序

### 打包
```shell
> docker build -t simplehttpserver:v1 .
```
### 启动docker镜像
```shell
> docker run -itd simplehttpserver:v1
```

### 测试接口
进入docker查看容器IP地址为172.17.0.2
```shell
> curl 172.17.0.2:8099/header -I
> curl 172.17.0.2:8099/version
> curl 172.17.0.2:8099/log
> curl 172.17.0.2:8099/healthz
```
