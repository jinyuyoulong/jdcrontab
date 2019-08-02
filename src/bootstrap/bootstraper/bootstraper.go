package bootstraper

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/view"
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/service"
	"time"
)

type Configurator func(*Bootstraper)

type Bootstraper struct {
	*iris.Application
	AppName string
	AppOwner string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstraper {
	b := &Bootstraper{
		AppName:appName,
		AppOwner:appOwner,
		AppSpawnDate:time.Now(),
		Application: iris.New(),
	}

	for _,cfg := range cfgs{
		cfg(b)
	}
	return b
}

func (b *Bootstraper)SetupErrorHandlers()  {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	StaticAssets  = "../public/"
	Favicon = "favicon.ico"
)

func (b *Bootstraper)Configure(cs ...Configurator)  {
	for _,c := range cs{
		c(b)
	}
}

func (b *Bootstraper) Bootstrap() *Bootstraper  {
	service.GetDi().Container.Invoke(func(viewEngine *view.HTMLEngine) {
		b.RegisterView(viewEngine)
	})
	service.GetDi().Container.Invoke(func(session *sessions.Sessions) {
		b.Sessions = session
	})

	b.SetupErrorHandlers()
	b.Favicon(StaticAssets+Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1],StaticAssets)

	return b
}



func (b *Bootstraper)Listen(addr string,cfgs ...iris.Configurator)  {
	b.Run(iris.Addr(addr),cfgs...)
}