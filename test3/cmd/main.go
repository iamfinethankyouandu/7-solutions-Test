package main

import (
	"7solutionstest3/config"
	"7solutionstest3/internal/handlers/rest"
	restadaptor "7solutionstest3/internal/repositories/adapters/rest"
	"7solutionstest3/internal/services"
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	cfg := config.InitConfig(ctx)

	// adaptor
	beefAPIAdaptor := restadaptor.NewBeefAPI(cfg.Adaptor.BeefAPI.URL)

	// service
	beefService := services.NewBeefSummaryService(beefAPIAdaptor)

	// handler
	beefHandler := rest.NewBeefSummaryHandler(beefService)

	// routing
	router := rest.InitRouter(beefHandler)

	// Start server
	srv := &http.Server{Addr: cfg.App.Port, Handler: router}
	go run(ctx, srv)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func run(ctx context.Context, srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(ctx, err.Error())
	}
}
