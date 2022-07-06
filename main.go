package main

import (
	"SkyWings/dao/mysql"
	"SkyWings/dao/redis"
	"SkyWings/logger"
	"SkyWings/pkg/snowflake"
	"SkyWings/pkg/validate"
	"SkyWings/routers"
	"SkyWings/settings"
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//var confFile string
	//flag.StringVar(&confFile, "conf", "./conf/config.yaml", "配置文件")
	//flag.Parse()
	// 加载配置
	if err := settings.Init(); err != nil {
		log.Printf("[%s] 配置加载失败！错误：%v\n", settings.Conf.App, err)
		return
	}
	log.Printf("[%s] 配置加载成功！", settings.Conf.App)

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		log.Printf("[%s] 日志记录器初始化失败！错误：%v\n", settings.Conf.App, err)
		return
	}
	log.Printf("[%s] 日志记录器初始化成功！", settings.Conf.App)

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		log.Printf("[%s] MySQL初始化失败！错误：%v\n", settings.Conf.App, err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	log.Printf("[%s] MySQL初始化成功！", settings.Conf.App)

	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("[%s] Redis初始化失败！错误：%v\n", settings.Conf.App, err)
		return
	}
	defer redis.Close()
	log.Printf("[%s] Redis初始化成功！", settings.Conf.App)

	if err := snowflake.Init(1); err != nil {
		log.Printf("[%s] 分布式ID生成器初始化失败！错误：%v\n", settings.Conf.App, err)
		return
	}
	log.Printf("[%s] 分布式ID生成器初始化成功！", settings.Conf.App)

	// 注册验证器
	if err := validate.InitValidator(); err != nil {
		fmt.Printf("[%s] 参数校验器初始化失败！错误：%v\n", settings.Conf.App, err)
	}
	log.Printf("[%s] 参数校验器初始化成功！", settings.Conf.App)

	// 注册路由
	r := routers.SetupRouter()
	log.Printf("[%s] 路由设置成功！", settings.Conf.App)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[%s] 服务启动失败！错误：%v\n", settings.Conf.App, err)
		}
	}()
	log.Printf("[%s] 已启动！正在%d端口监听：\n", settings.Conf.App, settings.Conf.Port)
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("[%s] 服务关闭中 ...", settings.Conf.App)
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown:", zap.Error(err))
	}

	// 优雅关机和平滑重启
	//if err := endless.ListenAndServe(fmt.Sprintf(":%d", settings.Conf.Port), r); err != nil {
	//	zap.L().Error("server listen :%v\n", zap.Error(err))
	//}
}
