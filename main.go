package main

/********************************************************************************
* Temancode Main Package                       		  	                        *
*                                                                               *
* Version: 1.0.0                                                                *
* Date:    2023-01-05                                                           *
* Author:  Waluyo Ade Prasetio                                                  *
* Github:  https://github.com/abdullahPrasetio                                  *
********************************************************************************/

import (
	"context"
	"flag"
	"fmt"
	//"github.com/abdullahPrasetio/base-go/scheduler"
	//"github.com/abdullahPrasetio/base-go/workers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abdullahPrasetio/base-go/configs"
	"github.com/abdullahPrasetio/base-go/constants"
	"github.com/abdullahPrasetio/base-go/utils/log"

	"github.com/abdullahPrasetio/base-go/routers"
	"github.com/gin-gonic/gin"
)

var Port string

func init() {
	config, _ := configs.LoadConfig(".")
	gin.SetMode(config.GIN_MODE)

	log.LoadLogger()
	flag.StringVar(&Port, "port", "80", "help message for port")
}

func main() {
	log := log.Logger
	router := routers.SetupRouter()
	flag.Parse()
	port := Port
	if port == "80" {
		port = constants.ServerPort
	}
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
		Handler:      router,
	}

	go func() {
		log.Info("about to start the application...")

		fmt.Println("Server run in url : http://localhost:" + port)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	// Worker scheduller
	//go workers.ProcessTimeScheduller()

	// Run scheduller
	//go scheduler.NewPrintTimeScheduller()

	// ====== Mematikan server secara gracefull ======

	// Membuat channel untuk menerima sinyal shutdown
	quit := make(chan os.Signal)

	// Menerima sinyal untuk mematikan server dari os
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("shutting down server...")

	// Memberikan waktu timeout 5 detik untuk mematikan server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Jika lebih dari 5 detik matikan paksa
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}

	log.Info("server shutdown gracefully")
}
