//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mazrean/todoList/common"
	v1Handler "github.com/mazrean/todoList/handler/v1"
	"github.com/mazrean/todoList/repository"
	"github.com/mazrean/todoList/repository/sqlx"
	"github.com/mazrean/todoList/service"
	v1Service "github.com/mazrean/todoList/service/v1"
)

type Config struct {
	Addr          common.Addr
	SessionKey    common.SessionKey
	SessionSecret common.SessionSecret
}

var (
	addrField          = wire.FieldsOf(new(*Config), "Addr")
	sessionKeyField    = wire.FieldsOf(new(*Config), "SessionKey")
	sessionSecretField = wire.FieldsOf(new(*Config), "SessionSecret")

	dbBind                   = wire.Bind(new(repository.DB), new(*sqlx.DB))
	userRepositoryBind       = wire.Bind(new(repository.User), new(*sqlx.User))
	dashboardRepositoryBind  = wire.Bind(new(repository.Dashboard), new(*sqlx.Dashboard))
	taskStatusRepositoryBind = wire.Bind(new(repository.TaskStatus), new(*sqlx.TaskStatus))
	taskRepositoryBind       = wire.Bind(new(repository.Task), new(*sqlx.Task))

	authorizationServiceBind = wire.Bind(new(service.Authorization), new(*v1Service.Authorization))
	dashboardServiceBind     = wire.Bind(new(service.Dashboard), new(*v1Service.Dashboard))
	taskStatusServiceBind    = wire.Bind(new(service.TaskStatus), new(*v1Service.TaskStatus))
	taskServiceBind          = wire.Bind(new(service.Task), new(*v1Service.Task))
)

func InjectAPI(config *Config) (*v1Handler.API, error) {
	wire.Build(
		addrField,
		sessionKeyField,
		sessionSecretField,

		dbBind,
		userRepositoryBind,
		dashboardRepositoryBind,
		taskStatusRepositoryBind,
		taskRepositoryBind,

		authorizationServiceBind,
		dashboardServiceBind,
		taskStatusServiceBind,
		taskServiceBind,

		v1Handler.NewAPI,
		v1Handler.NewContext,
		v1Handler.NewSession,
		v1Handler.NewMiddleware,
		v1Handler.NewUser,
		v1Handler.NewDashboard,
		v1Handler.NewTaskStatus,
		v1Handler.NewTask,

		sqlx.NewDB,
		sqlx.NewUser,
		sqlx.NewDashboard,
		sqlx.NewTaskStatus,
		sqlx.NewTask,

		v1Service.NewAuthorization,
		v1Service.NewDashboard,
		v1Service.NewTaskStatus,
		v1Service.NewTask,
	)

	return nil, nil
}
