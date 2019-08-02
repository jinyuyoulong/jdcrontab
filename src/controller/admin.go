package controller

import (
	"github.com/jinyuyoulong/jdcrontab/src/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// AdminController admin
type AdminController struct {
	Ctx iris.Context
}

const (
	adminTitle string = "管理后台"
)

// Get 根路径
// uri: /admin
func (c *AdminController) Get() mvc.Result {
	datalist := models.JdcronTasks{}.GetAll()

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    adminTitle,
			"Datalist": datalist,
		},
		Layout: "admin/layout.html",
	}
}

