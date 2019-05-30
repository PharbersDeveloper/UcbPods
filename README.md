#NewTMIST

> 新铁马 业务逻辑 代码


# 1. 环境搭建
## 1.1 配置好本机的GOROOT和GOPATH
如我的机器:
```angular2html
GOROOT=/usr/local/Cellar/go/1.10.3/libexec
GOPATH=$HOME/goWorkSpace
```

## 1.2 配置编译器的GOROOT和GOPATH
不同编译器略有差别, 自行去Google百度


# 2. 运行项目
使用命令`go get`, 依赖将自动安装到`$GOPATH`环境下  

## 2.1 本项目依次安装以下依赖:
```angular2html
go get github.com/alfredyang1986/blackmirror
go get github.com/alfredyang1986/BmServiceDef
// 上面语句如果报异常`can't load package`, 可以忽略, 
// 也可以进入下载的$GO_PATH/src对应的目录下, 新建一个main.go, 运行`go install`

// go get gopkg.in/mgo.v2
// go get gopkg.in/yaml.v2
// go get github.com/manyminds/api2go
// go get github.com/hashicorp/go-uuid
// go get github.com/aliyun/aliyun-oss-go-sdk/oss

```

## 2.2 下载配置文件
```
go get github.com/PharbersDeveloper/UcbServiceDeploy
```

## 2.3 配置环境变量NTM_HOME
```
export NTM_HOME=$GOPATH/src/github.com/PharbersDeveloper/NTMServiceDeploy/dev-config
```

## 2.4 执行main.go
```
go run main.go
```

# 3. 发布
## 3.1 下载docker文件
```
go get github.com/PharbersDeveloper/UcbServiceDeploy
```

## 3.2 编译docker file
```
cd $GOPATH
docker build -t ntm:1.0 -f src/github.com/PharbersDeveloper/UcbServiceDeploy/Dockerfile . 
```

## 3.3 docker 运行
```angular2html
docker run -p 31415:31415 -d ntm:1.0
```

## 3.4 docker 发布