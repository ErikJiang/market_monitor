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
$ git clone https://github.com/JiangInk/market_monitor.git
```

#### 运行项目

``` bash
$ cd market_monitor/app/src/
```

目录结构：
``` bash
app/src/:
    ├─config/       # 配置文件目录
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

启动服务：
``` bash
$ go run main.go
```

运行过程若出现下载 module 失败，或 build 缓慢，可尝试设置`GOPROXY`环境变量：
``` shell
$ export GOPROXY=https://goproxy.io
```

运行后，可访问：[`http://localhost:8000/swagger/index.html`](http://localhost:8000/swagger/index.html)，查看接口文档；

#### Docker 运行构建

主目录结构：
``` bash
market_monitor/
    ├─docker-compose.yaml   # compose 镜像服务构建文件
    ├─app                   # monitor 应用服务
    │  ├─src/               # 源码目录
    │  └─dockerfile         # docker 镜像构建文件
    ├─mysql                 # mysql 数据服务
    │  ├─conf
    │  │  ├─conf.d/         # 服务自定义配置
    │  │  └─init.d/         # 初始化SQL脚本及用户权限设置
    │  ├─data/              # 数据文件挂载目录
    │  └─logs/              # 日志目录
    └─redis                 # redis 缓存服务
        ├─conf/             # 缓存服务自定义配置
        └─data/             # 数据文件挂载目录
```

使用 `docker compose` 以后台启动方式构建容器服务：
``` bash
$ cd market_monitor/
$ docker-compose up -d
```

---

> 简明教程请见: [WIKI](https://github.com/JiangInk/market_monitor/wiki)

