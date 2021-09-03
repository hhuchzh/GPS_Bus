package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RouteInfoRouter struct {
}

// InitRouteInfoRouter 初始化 RouteInfo 路由信息
func (s *RouteInfoRouter) InitRouteInfoRouter(Router *gin.RouterGroup) {
	routeInfoRouter := Router.Group("routeInfo").Use(middleware.OperationRecord())
	var routeInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.RouteInfoApi
	{
		routeInfoRouter.POST("createRouteInfo", routeInfoApi.CreateRouteInfo)   // 新建RouteInfo
		routeInfoRouter.DELETE("deleteRouteInfo", routeInfoApi.DeleteRouteInfo) // 删除RouteInfo
		routeInfoRouter.DELETE("deleteRouteInfoByIds", routeInfoApi.DeleteRouteInfoByIds) // 批量删除RouteInfo
		routeInfoRouter.PUT("updateRouteInfo", routeInfoApi.UpdateRouteInfo)    // 更新RouteInfo
		routeInfoRouter.GET("findRouteInfo", routeInfoApi.FindRouteInfo)        // 根据ID获取RouteInfo
		routeInfoRouter.GET("getRouteInfoList", routeInfoApi.GetRouteInfoList)  // 获取RouteInfo列表
	}
}
