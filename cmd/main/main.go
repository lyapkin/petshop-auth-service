package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/lyapkin/shop/auth/config"
	"github.com/lyapkin/shop/auth/internal/app/usecases/auth"
	"github.com/lyapkin/shop/auth/internal/infrastructure/repositories/pgrole"
	"github.com/lyapkin/shop/auth/internal/infrastructure/repositories/pguser"
	"github.com/lyapkin/shop/auth/internal/infrastructure/repositories/redistoken"
	"github.com/lyapkin/shop/auth/internal/infrastructure/services/argon2pass"
	"github.com/lyapkin/shop/auth/internal/infrastructure/services/jwttoken"
	"github.com/lyapkin/shop/auth/internal/presentation/rest"
	"github.com/lyapkin/shop/auth/internal/storage/postgres"
	"github.com/lyapkin/shop/auth/internal/storage/redis"
	"github.com/lyapkin/shop/auth/pkg/logger"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	cfg := config.MustLoad()
	if cfg.Env != config.EnvProduction {
		fmt.Printf("start app in %s environment", cfg.Env)
	}
	fmt.Printf("start app in %s environment", cfg.Env)

	db, err := postgres.New(cfg.DB)
	if err != nil {
		log.Fatal("faild to setup db: %w", err)
	}

	redisCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	redisDB, err := redis.New(redisCtx, cfg.Redis)

	// repositories initialization
	userRepo := pguser.New(db)
	roleRepo := pgrole.New(db, cfg.InMemoryCacheTTL)
	tokenRepo := redistoken.New(redisDB)

	// services initialization
	passwordService := argon2pass.New()
	tokenService := jwttoken.New(&cfg.JWTToken)

	// logger initialization
	logger := logger.New(cfg.Env)
	slog.SetDefault(logger)

	// usecases init
	authUsecase := auth.New(
		logger,
		userRepo,
		roleRepo,
		passwordService,
		tokenService,
		tokenRepo,
	)

	handler := rest.New(authUsecase)

	http := runHTTPServer(&cfg.HTTPServer, handler)

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(ctx, cfg.ShutdownTimeout)
	defer cancel()

	if err := http.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("shutdown server: %v", err)
	}

	if err := db.Close(); err != nil {
		fmt.Printf("close db: %v", err)
	}

	if err := redisDB.Close(); err != nil {
		fmt.Printf("close redis: %v", err)
	}
}

func runHTTPServer(cfg *config.HTTPServer, handler http.Handler) *http.Server {
	server := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server stoped: %v", err)
		}
	}()

	return server
}
