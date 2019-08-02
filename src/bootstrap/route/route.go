package route

import (
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/bootstraper"
	"github.com/jinyuyoulong/jdcrontab/src/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// SetRoute 配置路由
func SetRoute(route *iris.Application) {
	IndexRoute(route)
	AdminRoute(route)
}

// IndexRoute 配置index route
func IndexRoute(route *iris.Application) {
	indexC := new(controller.IndexController)
	index := mvc.New(route.Party("/"))
	index.Handle(indexC)
}

// AdminRoute admin route
func AdminRoute(route *iris.Application) {
	admin := mvc.New(route.Party("/admin"))
	admin.Handle(new(controller.AdminController))
}

func Configure(b *bootstraper.Bootstraper)  {
	IndexRoute(b.Application)
	AdminRoute(b.Application)
}