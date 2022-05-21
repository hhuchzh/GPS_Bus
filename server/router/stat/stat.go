package stat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StatRouter struct {
}

func (s *StatRouter) InitStatRouter(Router *gin.RouterGroup) {
	statRouter := Router.Group("stat").Use(middleware.OperationRecord())
	var statApi = v1.ApiGroupApp.StatApiGroup.StatApi
	{

		statRouter.GET("location", statApi.GetShuttleStatistics)
	}
}
