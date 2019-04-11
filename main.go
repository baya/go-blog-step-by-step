package main

import (
	"context"
	"fmt"
	"go-blog-step-by-step/pkg/setting"
	"go-blog-step-by-step/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	//if err := s.ListenAndServe(); err != nil {
	//	log.Printf("Listen: %s\n", err)
	//}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	log.Println("serving...")
	<- quit
	log.Println("try sig again...")
	<- quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}

