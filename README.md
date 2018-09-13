# market_monitor

#### 1. 安装`govendor`
``` shell
$ go get -u -v github.com/kardianos/govendor

```

#### 2. 同步项目依赖
``` shell
$ cd market_monitor
$ govendor sync
```

#### 3. 创建数据库
``` sql
CREATE DATABASE monitor DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
``` 

#### 4. 运行项目
``` shell
$ go run main.go
```
