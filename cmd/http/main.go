package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	authHttp "userman/internal/interface/http/auth"
	"userman/internal/interface/http/common/task"
	"userman/internal/interface/http/container"
	userHttp "userman/internal/interface/http/user"

	"github.com/gin-contrib/graceful"
)

func main() {

	c := container.NewContainer()
	if err := c.BuildContainer(); err != nil {
		log.Fatal("Failed to build container:", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router, err := graceful.Default()
	if err != nil {
		panic(err)
	}
	defer router.Close()

	api := router.Group("/api")
	userGroup := api.Group("/users")
	authGroup := api.Group("/auth")

	err = c.Invoke(func(rs container.RouterSetup) {
		router.Use(rs.LoggingMiddleware.Log())

		task.SeedingAdminTask(ctx, rs.UserRepository)

		userGroup.Use(rs.AuthMiddleware.Validate())

		authHttp.AuthRouter(authGroup, rs.AuthHandler)
		userHttp.UserRouter(userGroup, rs.UserHandler)

		go task.CountingUserTask(ctx, rs.UserRepository)

		if err := router.RunWithContext(ctx); err != nil && err != context.Canceled {
			panic(err)
		}
	})

	if err != nil {
		log.Fatal("Failed to setup routes: ", err)
	}
}
