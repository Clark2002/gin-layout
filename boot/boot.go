package boot

import (
	"github.com/Clark2002/gin-layout/data"
	"github.com/Clark2002/gin-layout/internal/pkg/logger"
	"github.com/Clark2002/gin-layout/internal/validator"
)

func init() {
	// 1、初始化zap日志
	logger.InitLogger()

	// 2、初始化数据库
	data.InitData()

	// 3、初始化验证器
	if err := validator.InitValidatorTrans("zh"); err != nil {
		panic("验证器初始化失败")
	}
}
