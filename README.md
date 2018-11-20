# market_monitor

#### 1. 克隆项目
``` shell
$ git clone https://github.com/JiangInk/market_monitor.git
```

#### 2. 下载项目依赖
``` shell
$ cd market_monitor
$ go mod download
```

#### 3. 创建数据库
``` sql
CREATE DATABASE db_monitor DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
``` 

#### 4. 运行项目
``` shell
$ go run main.go
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

