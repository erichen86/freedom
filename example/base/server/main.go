//Package main code generated by 'freedom new-project base'
package main

import (
	"time"

	"github.com/8treenet/freedom/example/base/server/conf"

	"github.com/8treenet/freedom"
	_ "github.com/8treenet/freedom/example/base/adapter/controller"
	"github.com/8treenet/freedom/infra/requests"
	"github.com/8treenet/freedom/middleware"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	app := freedom.NewApplication()
	/*
		installDatabase(app)
		installRedis(app)

		http2 h2c service
		h2caddrRunner := app.CreateH2CRunner(conf.Get().App.Other["listen_addr"].(string))
	*/

	installMiddleware(app)
	addrRunner := app.CreateRunner(conf.Get().App.Other["listen_addr"].(string))
	//app.InstallParty("/base")
	liveness(app)
	app.Run(addrRunner, *conf.Get().App)
}

func installMiddleware(app freedom.Application) {
	//Recover中间件
	app.InstallMiddleware(middleware.NewRecover())
	//Trace链路中间件
	app.InstallMiddleware(middleware.NewTrace("x-request-id"))

	//自定义请求日志配置文件
	//1.打印UA
	//2.修改请求日志前缀
	loggerConfig := middleware.DefaultLoggerConfig()
	loggerConfig.MessageHeaderKeys = append(loggerConfig.MessageHeaderKeys, "User-Agent")
	loggerConfig.Title = "base-access"

	//日志中间件，每个请求一个logger
	app.InstallMiddleware(middleware.NewRequestLogger("x-request-id", loggerConfig))

	//logRow中间件，每一行日志都会触发回调。如果返回true，将停止中间件遍历回调。
	app.Logger().Handle(middleware.DefaultLogRowHandle)
	//HttpClient 普罗米修斯中间件，监控下游的API请求。
	requests.InstallPrometheus(conf.Get().App.Other["service_name"].(string), freedom.Prometheus())
	//总线中间件，处理上下游透传的Header
	app.InstallBusMiddleware(middleware.NewBusFilter())
}

func installDatabase(app freedom.Application) {
	app.InstallDB(func() interface{} {
		conf := conf.Get().DB
		db, e := gorm.Open("mysql", conf.Addr)
		if e != nil {
			freedom.Logger().Fatal(e.Error())
		}

		db.DB().SetMaxIdleConns(conf.MaxIdleConns)
		db.DB().SetMaxOpenConns(conf.MaxOpenConns)
		db.DB().SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeTime) * time.Second)
		return db
	})
}

func installRedis(app freedom.Application) {
	app.InstallRedis(func() (client redis.Cmdable) {
		cfg := conf.Get().Redis
		opt := &redis.Options{
			Addr:               cfg.Addr,
			Password:           cfg.Password,
			DB:                 cfg.DB,
			MaxRetries:         cfg.MaxRetries,
			PoolSize:           cfg.PoolSize,
			ReadTimeout:        time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout:       time.Duration(cfg.WriteTimeout) * time.Second,
			IdleTimeout:        time.Duration(cfg.IdleTimeout) * time.Second,
			IdleCheckFrequency: time.Duration(cfg.IdleCheckFrequency) * time.Second,
			MaxConnAge:         time.Duration(cfg.MaxConnAge) * time.Second,
			PoolTimeout:        time.Duration(cfg.PoolTimeout) * time.Second,
		}
		redisClient := redis.NewClient(opt)
		if e := redisClient.Ping().Err(); e != nil {
			freedom.Logger().Fatal(e.Error())
		}
		client = redisClient
		return
	})
}

func liveness(app freedom.Application) {
	app.Iris().Get("/ping", func(ctx freedom.Context) {
		ctx.WriteString("pong")
	})
}
