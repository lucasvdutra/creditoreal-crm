package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"creditoreal-crm/internal/auth"
	"creditoreal-crm/internal/authorization"
	"creditoreal-crm/internal/config"
	"creditoreal-crm/internal/health"
	apphttp "creditoreal-crm/internal/http/middleware"
	"creditoreal-crm/internal/http/respond"
	"creditoreal-crm/internal/logger"
	"creditoreal-crm/internal/property"
	"creditoreal-crm/internal/propertyaddress"
	"creditoreal-crm/internal/propertytype"
	"creditoreal-crm/internal/rbac"
	"creditoreal-crm/internal/relationship"
	"creditoreal-crm/internal/tenant"
	"creditoreal-crm/internal/user"
	"creditoreal-crm/pkg/database"
	db "creditoreal-crm/pkg/database/queries"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	sqlDB, err := database.Open(cfg.DatabaseURL)
	if err != nil {
		log.Error("database.open_failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer sqlDB.Close()

	if err := database.Ping(context.Background(), sqlDB); err != nil {
		log.Error("database.ping_failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	queries := db.New(sqlDB)
	tokenManager := auth.NewTokenManager(cfg.JWTSecret)
	authorizationService := authorization.NewService(queries)
	tenantHandler := tenant.NewHandler(tenant.NewService(queries))
	userHandler := user.NewHandler(user.NewService(queries))
	rbacHandler := rbac.NewHandler(rbac.NewService(queries))
	relationshipHandler := relationship.NewHandler(relationship.NewService(queries))
	propertyTypeHandler := propertytype.NewHandler(propertytype.NewService(queries))
	propertyHandler := property.NewHandler(property.NewService(queries, authorizationService), tokenManager)
	propertyAddressHandler := propertyaddress.NewHandler(propertyaddress.NewService(queries, authorizationService), tokenManager)
	authHandler := auth.NewHandler(auth.NewService(
		queries,
		tokenManager,
		time.Duration(cfg.AccessTokenMinutes)*time.Minute,
		time.Duration(cfg.RefreshTokenHours)*time.Hour,
	))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", health.Handle)
	tenantHandler.Register(mux)
	userHandler.Register(mux)
	rbacHandler.Register(mux)
	relationshipHandler.Register(mux)
	propertyTypeHandler.Register(mux)
	propertyHandler.Register(mux)
	propertyAddressHandler.Register(mux)
	authHandler.Register(mux)
	mux.Handle("GET /api/auth/me", apphttp.Auth(tokenManager, http.HandlerFunc(me)))

	handler := apphttp.RequestID(mux)
	server := &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Info("api.started", slog.String("addr", cfg.HTTPAddr), slog.String("env", cfg.AppEnv))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("api.failed", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("api.shutdown_failed", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("api.stopped")
}

func me(w http.ResponseWriter, r *http.Request) {
	authCtx, ok := apphttp.AuthFromContext(r.Context())
	if !ok {
		respond.Error(w, http.StatusUnauthorized, "permissao.negada", "Permissao negada.")
		return
	}

	respond.JSON(w, http.StatusOK, map[string]string{
		"user_id":   authCtx.UserID,
		"tenant_id": authCtx.TenantID,
	})
}
