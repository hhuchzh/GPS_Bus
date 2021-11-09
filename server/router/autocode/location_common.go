package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LocationCommonRouter struct {
}

// InitLocationCommonRouter 初始化 LocationCommon 路由信息
func (s *LocationCommonRouter) InitLocationCommonRouter(Router *gin.RouterGroup) {
	locationCommonRouter := Router.Group("locationCommon").Use(middleware.OperationRecord())
	var locationCommonApi = v1.ApiGroupApp.AutoCodeApiGroup.LocationCommonApi
	{
		locationCommonRouter.POST("createLocationCommon", locationCommonApi.CreateLocationCommon)   // 新建LocationCommon
		locationCommonRouter.DELETE("deleteLocationCommon", locationCommonApi.DeleteLocationCommon) // 删除LocationCommon
		locationCommonRouter.DELETE("deleteLocationCommonByIds", locationCommonApi.DeleteLocationCommonByIds) // 批量删除LocationCommon
		locationCommonRouter.PUT("updateLocationCommon", locationCommonApi.UpdateLocationCommon)    // 更新LocationCommon
		locationCommonRouter.GET("findLocationCommon", locationCommonApi.FindLocationCommon)        // 根据ID获取LocationCommon
		locationCommonRouter.GET("getLocationCommonList", locationCommonApi.GetLocationCommonList)  // 获取LocationCommon列表
	}
}
