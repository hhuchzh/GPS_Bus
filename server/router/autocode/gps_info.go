package autocode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GpsInfoRouter struct {
}

// InitGpsInfoRouter 初始化 GpsInfo 路由信息
func (s *GpsInfoRouter) InitGpsInfoRouter(Router *gin.RouterGroup) {
	gpsInfoRouter := Router.Group("gpsInfo").Use(middleware.OperationRecord())
	var gpsInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.GpsInfoApi
	{
		gpsInfoRouter.POST("createGpsInfo", gpsInfoApi.CreateGpsInfo)                    // 新建GpsInfo
		gpsInfoRouter.DELETE("deleteGpsInfo", gpsInfoApi.DeleteGpsInfo)                  // 删除GpsInfo
		gpsInfoRouter.DELETE("deleteGpsInfoByIds", gpsInfoApi.DeleteGpsInfoByIds)        // 批量删除GpsInfo
		gpsInfoRouter.PUT("updateGpsInfo", gpsInfoApi.UpdateGpsInfo)                     // 更新GpsInfo
		gpsInfoRouter.GET("findGpsInfo", gpsInfoApi.FindGpsInfo)                         // 根据ID获取GpsInfo
		gpsInfoRouter.GET("getGpsInfoList", gpsInfoApi.GetGpsInfoList)                   // 获取GpsInfo列表
		gpsInfoRouter.GET("getAvailableGpsInfoList", gpsInfoApi.GetAvailableGpsInfoList) // 获取GpsInfo列表
		gpsInfoRouter.GET("getNotAvailableGpsInfoList", gpsInfoApi.GetNotAvailableGpsInfoList) // 获取GpsInfo列表
	}
}
