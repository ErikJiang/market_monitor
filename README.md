# :guardsman: market_monitor

![gopher](./public/gopher.png)

* web framework: [gin](https://github.com/gin-gonic/gin)
* redis: [redigo](https://github.com/gomodule/redigo)
* mysql: [gorm](https://github.com/jinzhu/gorm)
* logger: [zerolog](https://github.com/rs/zerolog)
* scheduler: [cron](https://github.com/robfig/cron)
* config: [viper](https://github.com/spf13/viper)
* json web token: [jwt-go](https://github.com/dgrijalva/jwt-go)
* swagger docs: [swaggo](https://github.com/swaggo/gin-swagger)


#### 1. 克隆项目
``` shell
$ git clone https://github.com/JiangInk/market_monitor.git
```

#### 2. 创建数据库
``` sql
CREATE DATABASE db_monitor DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
```

#### 3. 运行项目
``` shell
$ cd market_monitor/
$ go run main.go
```
运行过程若出现下载 module 失败，或 build 缓慢，可尝试设置`GOPROXY`环境变量：
``` shell
$ export GOPROXY=https://goproxy.io
```


---

> 简明教程请见:[WIKI](https://github.com/JiangInk/market_monitor/wiki)

