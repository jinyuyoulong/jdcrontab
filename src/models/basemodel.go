package models

import (
	"github.com/jinyuyoulong/jdcrontab/src/bootstrap/service"

	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func init() {
	if engine == nil {
		container := service.GetDi().Container
		container.Invoke(func(db *xorm.Engine) {
			engine = db
		})
	}
}
