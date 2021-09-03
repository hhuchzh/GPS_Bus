package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArrivalInfoRouter struct {
}

// InitArrivalInfoRouter 初始化 ArrivalInfo 路由信息
func (s *ArrivalInfoRouter) InitArrivalInfoRouter(Router *gin.RouterGroup) {
	arrivalInfoRouter := Router.Group("arrivalInfo").Use(middleware.OperationRecord())
	var arrivalInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.ArrivalInfoApi
	{
		arrivalInfoRouter.POST("createArrivalInfo", arrivalInfoApi.CreateArrivalInfo)   // 新建ArrivalInfo
		arrivalInfoRouter.DELETE("deleteArrivalInfo", arrivalInfoApi.DeleteArrivalInfo) // 删除ArrivalInfo
		arrivalInfoRouter.DELETE("deleteArrivalInfoByIds", arrivalInfoApi.DeleteArrivalInfoByIds) // 批量删除ArrivalInfo
		arrivalInfoRouter.PUT("updateArrivalInfo", arrivalInfoApi.UpdateArrivalInfo)    // 更新ArrivalInfo
		arrivalInfoRouter.GET("findArrivalInfo", arrivalInfoApi.FindArrivalInfo)        // 根据ID获取ArrivalInfo
		arrivalInfoRouter.GET("getArrivalInfoList", arrivalInfoApi.GetArrivalInfoList)  // 获取ArrivalInfo列表
	}
}
