package main

import (
	"context"
	"crud-go/internal/config"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main () {
	fmt.Println("Hello word")

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student api"))
	})


	server := http.Server {
		Addr: cfg.Addr,
		Handler: router,
	}

	fmt.Println("server started", cfg.HttpServer.Addr)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT)

	go func ()  {
		err := server.ListenAndServe()

	if (err != nil) {
		log.Fatal("failed to start server")
	}
		
	} ()

	<-done

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 3* time.Second)
	defer cancel()
	//err := server.Shutdown(ctx)

	// if (err != nil) {
	// 	slog.Error("failed to shutdown", slog.String("error", err.Error()))
	// }

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successful")

	
}