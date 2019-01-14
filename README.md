# :guardsman: market_monitor

![gopher](./app/src/public/gopher.png)

* web framework: [gin](https://github.com/gin-gonic/gin)
* redis: [redigo](https://github.com/gomodule/redigo)
* mysql: [gorm](https://github.com/jinzhu/gorm)
* logger: [zerolog](https://github.com/rs/zerolog)
* scheduler: [cron](https://github.com/robfig/cron)
* config: [viper](https://github.com/spf13/viper)
* json web token: [jwt-go](https://github.com/dgrijalva/jwt-go)
* swagger docs: [swaggo](https://github.com/swaggo/gin-swagger)

#### 克隆项目
``` shell
$ git clone https://github.com/ErikJiang/market_monitor.git
```

#### 单独运行项目

进入应用服务源码目录：
``` bash
$ cd market_monitor/app/src/
```

目录结构：
``` bash
app/src/:
    ├─config/       # 配置文件目录（运行前需要调整database、redis等参数配置）
    ├─controller/   # 控制器层目录
    ├─docs/         # 接口文档 Swagger 生成目录
    ├─extend/       # 扩展模块目录
    ├─middleware/   # 中间件目录
    ├─models/       # 数据模型层
    ├─public/       # 静态资源目录
    ├─router/       # 路由目录
    ├─schedule/     # 定时任务目录
    ├─service/      # 服务层目录
    ├─templates/    # 视图模板层目录
    ├─go.mod        # 包管理文件
    ├─go.sum        # 依赖包版本哈希文件
    └─main.go       # 运行入口main文件
```

自主创建数据库：
``` sql
CREATE DATABASE db_monitor DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
```

启动服务：
``` bash
$ go run main.go
```

运行过程若出现下载 module 失败，或 build 缓慢，可尝试设置`GOPROXY`环境变量：
``` shell
$ export GOPROXY=https://goproxy.io
```

运行后访问：[`http://localhost:8000/swagger/index.html`](http://localhost:8000/swagger/index.html)，查看接口文档；

#### Docker 运行构建
进入项目根目录：
``` bash
$ cd market_monitor/
```

主目录结构：
``` bash
market_monitor/
    ├─docker-compose.yaml   # compose 镜像服务构建文件
    ├─app                   # monitor 应用服务
    │  ├─src/               # 源码目录
    │  └─dockerfile         # docker 镜像构建文件
    ├─mysql                 # mysql 数据服务
    │  ├─conf
    │  │  ├─conf.d/         # 服务自定义配置（字符编码等）
    │  │  └─init.d/         # 初始化SQL脚本（建库语句）及用户权限设置
    │  ├─data/              # 数据文件挂载目录
    │  └─logs/              # 日志目录
    └─redis                 # redis 缓存服务
        ├─conf/             # 缓存服务自定义配置(密码等)
        └─data/             # 数据文件挂载目录
```

由于 app/ 下的 dockerfile 指定 golang 编译所需依赖从 app/src/vendor/ 目录中读取，故需要提前准备编译所需依赖：
``` bash
$ cd app/src/
$ go mod vendor
```

最后再回到项目根目录，使用 `docker compose` 以后台启动方式构建容器服务：
``` bash
$ docker-compose up -d
```

---

> 简明教程请见: [Wiki](https://github.com/ErikJiang/market_monitor/wiki)

