package main

import (
	"context"
	"fmt"
	"github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"golibs/utils"
	"go-demo-project/server/routers"
	"golibs/conf"
)

const LOGDIR = "/data/log/demo-server/"

func main() {
	var err error

	// 初始化 log
	err = utils.EnsurePath(LOGDIR)
	if err != nil {
		fmt.Println(err)
		return
	}
	log4go.LoadConfiguration("conf/log.xml")
	defer log4go.Close()

	// 初始化 ini
	err = conf.Init("conf/conf.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	// http server
	f, _ := os.Create(LOGDIR + "gin.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	g := gin.Default()

	// routers
	err = routers.Init(g)
	if err != nil {
		fmt.Println(err)
		return
	}

	// server
	httpConf, err := utils.GetHttpConf()
	if err != nil {
		fmt.Println(err)
		return
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", httpConf.Port),
		Handler:      g,
		ReadTimeout:  time.Duration(httpConf.ReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(httpConf.WriteTimeout) * time.Millisecond,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log4go.Error("listen fail: %v", err)
		}
	}()

	// gracefully shutdown when signal
	chQuit := make(chan os.Signal)
	signal.Notify(chQuit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-chQuit
	log4go.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(httpConf.ShutdownTimeout)*time.Millisecond)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log4go.Error("Server Shutdown: %v", err)
	}
	log4go.Info("Server exiting")

	return
}
