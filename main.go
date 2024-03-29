package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"maan/common"
	"maan/config"
	_ "maan/pkg/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	common.InitRouter(r)
	Run(r, config.Conf.ServerConfig.Name, config.Conf.ServerConfig.Addr)
}

func Run(r *gin.Engine, srvName, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// 优雅的启停
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal)
	// SIGINT 用户发送 INTR 字符(Ctrl + C)触发 kill -2
	// SIGTERM 结束程序 (可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutting Down project %s ...\n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s shutdown, cause by : %v\n", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Printf("%s stop success...\n", srvName)
}
