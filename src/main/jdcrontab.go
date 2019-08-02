package main

import (
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/bootstraper"
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/middleware"
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/route"
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/service"
)

func main() {
	app := initApplication()
	var port string
	service.GetDi().Container.Invoke(func(config *service.Config) {
		port = config.App.Port
	})

	app.Listen(port)
}

func initApplication() *bootstraper.Bootstraper {
	di := service.GetDi()
	container := di.Container
	var appOwner string
	container.Invoke(func(config *service.Config) {
		appOwner = config.App.AppOwner
	})
	app := bootstraper.New("计划任务",appOwner)
	app.Bootstrap()

	app.Configure(route.Configure)
	middleware.RegistMiddleware(app.Application)

	return app
}
