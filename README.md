# :guardsman: market_monitor

<img src="./public/gopher.png" alt="gopher" width="260" height="278"/>

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

#### 开发相关
##### 1. 生成 swagger 文档：
需要先安装：
* github.com/swaggo/swag/cmd/swag
* github.com/swaggo/gin-swagger
* github.com/swaggo/gin-swagger/swaggerFiles

`go get`无法下载，可以考虑使用[gopm](https://gopm.io/)进行下载；
初始化生成文档
``` bash
$ swag init
```

