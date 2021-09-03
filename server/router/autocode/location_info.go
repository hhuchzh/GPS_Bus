package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LocationInfoRouter struct {
}

// InitLocationInfoRouter 初始化 LocationInfo 路由信息
func (s *LocationInfoRouter) InitLocationInfoRouter(Router *gin.RouterGroup) {
	locationInfoRouter := Router.Group("locationInfo").Use(middleware.OperationRecord())
	var locationInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.LocationInfoApi
	{
		locationInfoRouter.POST("createLocationInfo", locationInfoApi.CreateLocationInfo)   // 新建LocationInfo
		locationInfoRouter.DELETE("deleteLocationInfo", locationInfoApi.DeleteLocationInfo) // 删除LocationInfo
		locationInfoRouter.DELETE("deleteLocationInfoByIds", locationInfoApi.DeleteLocationInfoByIds) // 批量删除LocationInfo
		locationInfoRouter.PUT("updateLocationInfo", locationInfoApi.UpdateLocationInfo)    // 更新LocationInfo
		locationInfoRouter.GET("findLocationInfo", locationInfoApi.FindLocationInfo)        // 根据ID获取LocationInfo
		locationInfoRouter.GET("getLocationInfoList", locationInfoApi.GetLocationInfoList)  // 获取LocationInfo列表
	}
}
