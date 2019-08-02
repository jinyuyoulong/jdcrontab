package page

import "github.com/jinyuyoulong/jdcrontab/src/library/page/pagemodel"

//计算offset 和 limit
func GetDbNum(condition pagemodel.PageQueryCondition) (dbnum model.DbNum) {
	if condition.PageNum <= 1 {
		condition.PageNum = 1
	}
	if condition.PageSize <= 1 {
		condition.PageSize = 10
	}
	dbnum.Limit = condition.PageSize
	dbnum.Offset = (condition.PageNum - 1) * condition.PageSize
	return
}
