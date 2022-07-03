package routers

import (
	"github.com/Clark2002/gin-layout/config"
	"github.com/Clark2002/gin-layout/internal/middleware"
	"github.com/Clark2002/gin-layout/internal/pkg/error_code"
	response2 "github.com/Clark2002/gin-layout/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if config.Config.Debug == false {
		// 生产模式
		r = ReleaseRouter()
		r.Use(
			middleware.RequestCostHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		// 开发调试模式
		r = gin.New()
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}

	// 设置 API 路由
	setApiRoute(r)

	r.NoRoute(func(c *gin.Context) {
		response2.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, error_code.ServerError, "资源不存在")
	})

	return r
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()

	return engine
}
