package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	. "go-test/bootstrap"
	. "go-test/config"
	. "go-test/routes"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	runMode := flag.String("mode", "dev", "run mode")
	flag.Parse()
	InitConfig(*runMode)
	InitDB()

	defer Db.Close()
	gin.SetMode(Config.GetString("app.env"))

	// //日志
	_ = os.MkdirAll(Config.GetString("logging.log_path"), os.ModePerm)
	logFile, _ := os.Create(Config.GetString("logging.log_path") + Config.GetString("logging.log_file"))
	errorLogFile, _ := os.Create(Config.GetString("logging.log_path") + Config.GetString("logging.error_log_file"))

	gin.DefaultWriter = io.MultiWriter(logFile)
	gin.DefaultErrorWriter = io.MultiWriter(errorLogFile)

	//初始化路由
	router := InitRouter()

	httpServer := &http.Server{
		Addr:    ":" + Config.GetString("http.server.port"),
		Handler: router,
	}

	go func() {
		log.Println("Start Http Server ...")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Start Http Server Failed: %s\n", err)
		}
	}()

	//优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
