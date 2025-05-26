package container

import (
	userApp "userman/internal/application/user"
	"userman/internal/domain/user"
	"userman/internal/infrastructure/database"
	"userman/internal/infrastructure/http/handler"
	middlewareInfra "userman/internal/infrastructure/http/middleware"
	"userman/internal/infrastructure/jwt"
	mongoRepository "userman/internal/infrastructure/repository/mongo"
	authHttp "userman/internal/interface/http/auth"
	middlewareInterface "userman/internal/interface/http/common/middleware"
	"userman/internal/interface/http/config"
	userHttp "userman/internal/interface/http/user"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

type RouterSetup struct {
	dig.In
	UserHandler       userHttp.UserHandler
	AuthHandler       authHttp.AuthHandler
	Config            *config.Config
	UserRepository    user.UserRepository
	AuthMiddleware    middlewareInterface.AuthMiddleware
	LoggingMiddleware middlewareInterface.LoggingMiddleware
}

func NewContainer() *Container {
	return &Container{
		container: dig.New(),
	}
}

func (c *Container) BuildContainer() error {

	if err := c.container.Provide(config.NewConfig); err != nil {
		return err
	}

	if err := c.container.Provide(func(cfg *config.Config) (*mongo.Database, error) {
		return database.NewMongoDatabase(cfg.MongoDb.Uri, cfg.MongoDb.Database)
	}); err != nil {
		return err
	}

	if err := c.container.Provide(func(cfg *config.Config) jwt.JWTService {
		return jwt.NewJWTService(cfg.Jwt.SecretKey, cfg.Jwt.ExpireSec)
	}); err != nil {
		return err
	}

	if err := c.container.Provide(mongoRepository.NewUserRepository); err != nil {
		return err
	}

	if err := c.container.Provide(userApp.NewUserService); err != nil {
		return err
	}

	if err := c.container.Provide(middlewareInfra.NewLoggingMiddleware); err != nil {
		return err
	}

	if err := c.container.Provide(middlewareInfra.NewAuthMiddleware); err != nil {
		return err
	}

	if err := c.container.Provide(handler.NewUserHandler); err != nil {
		return err
	}

	if err := c.container.Provide(handler.NewAuthHandler); err != nil {
		return err
	}

	return nil
}

func (c *Container) Invoke(fn any, opts ...dig.InvokeOption) error {
	return c.container.Invoke(fn, opts...)
}
