//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/buffge/gobyexample/third/kratos/internal/biz"
	"github.com/buffge/gobyexample/third/kratos/internal/conf"
	"github.com/buffge/gobyexample/third/kratos/internal/data"
	"github.com/buffge/gobyexample/third/kratos/internal/server"
	"github.com/buffge/gobyexample/third/kratos/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
