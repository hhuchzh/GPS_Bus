package gps

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GpsRouter struct {
}

func (g *GpsRouter) InitGpsRouter(Router *gin.RouterGroup) {
	gpsRouter := Router.Group("gps").Use(middleware.OperationRecord())
	var gpsApi = v1.ApiGroupApp.GpsApiGroup.GpsApi
	{
		//gpsRouter.GET("device", gpsApi.ListDevice)
		//gpsRouter.GET("location", gpsApi.ListLocation)
		gpsRouter.GET("location", gpsApi.GetLocation)
		gpsRouter.POST("routelocation", gpsApi.GetLocationEx)
		gpsRouter.GET("locationlist", gpsApi.ListLocation)
		gpsRouter.GET("track", gpsApi.ListTrack)

		gpsRouter.GET("trackbus", gpsApi.ListBusTrack)
	}
}
