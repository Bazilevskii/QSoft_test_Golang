package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"qsoft-go-test/app/api"
	"syscall"
	"time"
)

// App - is a function with server settings, running, logging, gracefully shutdown, handlers connection and comments :)
func App() {
	gin.SetMode(gin.ReleaseMode)

	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Logging connection error")
		}
	}(logger) // flushes buffer, if any
	sugar := logger.Sugar()

	err := godotenv.Load(".env")
	if err != nil {
		sugar.Fatalf("Config loading error: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		sugar.Fatal("Port not found")
	}

	// server settings
	srv := &http.Server{
		Addr:    port,
		Handler: api.InitHandlers(),
	}

	// server running
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalf("listen: %s\n", err)
		}
	}()

	// simple logging about server started
	sugar.Infof(`Server is running on the port %s`, port)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sugar.Info("Disconnection from server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Fatal("Server shut down because of error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		sugar.Fatal("5 seconds delay")
	}
	sugar.Fatal("Server shutdown")
}
