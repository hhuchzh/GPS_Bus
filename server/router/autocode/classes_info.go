package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClassesInfoRouter struct {
}

// InitClassesInfoRouter 初始化 ClassesInfo 路由信息
func (s *ClassesInfoRouter) InitClassesInfoRouter(Router *gin.RouterGroup) {
	classesInfoRouter := Router.Group("classesInfo").Use(middleware.OperationRecord())
	var classesInfoApi = v1.ApiGroupApp.AutoCodeApiGroup.ClassesInfoApi
	{
		classesInfoRouter.POST("createClassesInfo", classesInfoApi.CreateClassesInfo)   // 新建ClassesInfo
		classesInfoRouter.DELETE("deleteClassesInfo", classesInfoApi.DeleteClassesInfo) // 删除ClassesInfo
		classesInfoRouter.DELETE("deleteClassesInfoByIds", classesInfoApi.DeleteClassesInfoByIds) // 批量删除ClassesInfo
		classesInfoRouter.PUT("updateClassesInfo", classesInfoApi.UpdateClassesInfo)    // 更新ClassesInfo
		classesInfoRouter.GET("findClassesInfo", classesInfoApi.FindClassesInfo)        // 根据ID获取ClassesInfo
		classesInfoRouter.GET("getClassesInfoList", classesInfoApi.GetClassesInfoList)  // 获取ClassesInfo列表
	}
}
