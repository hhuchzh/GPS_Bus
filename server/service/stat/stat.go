package stat

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/stat"
	"go.uber.org/zap"
)

type StatService struct {
}

type NResult struct {
	N float32 //or int ,or some else
}
type RouteGroup struct {
	Identifier string
}

func (ss *StatService) GetShuttleStatistics() (*model.StatInfo, error) {
	var groupField = "route_name_ex"
	var cnt int64
	var stat model.StatInfo

	// get route infos
	var routeGroups []RouteGroup
	db := global.GVA_DB.Model(&autocode.RouteInfo{})
	result := db.Select(groupField + " as identifier").Where(groupField + " is not null").Group(groupField).Find(&routeGroups)
	_ = result.Count(&cnt)

	stat.RouteNum = int(cnt)
	global.GVA_LOG.Info("statistics:", zap.Int("Route Num", stat.RouteNum))

	// get bus infos
	db = global.GVA_DB.Model(&autocode.BusInfo{})
	db.Count(&cnt)
	stat.ShuttleNum = int(cnt)
	global.GVA_LOG.Info("statistics:", zap.Int("Shuttle Num", stat.ShuttleNum))

	// get location info
	db = global.GVA_DB.Model(&autocode.LocationInfo{})
	db.Count(&cnt)
	stat.StationNum = int(cnt)
	global.GVA_LOG.Info("statistics:", zap.Int("Station Num", stat.StationNum))

	// get total miles
	var n NResult
	now := time.Now().Add(-24 * time.Hour * 30)
	date := now.Format("2006-01-02")
	db = global.GVA_DB.Model(&autocode.MilesInfo{})
	db.Select("sum(distance) as n").Where("cal_date > ?", date).Scan(&n)
	stat.TotalMileage = int(n.N) / 1000
	global.GVA_LOG.Info("statistics:", zap.Int("Mileage(30 days)", stat.TotalMileage))

	// get detail route
	for _, routeGroup := range routeGroups {
		var routeInfos []autocode.RouteInfo
		var mapBus = make(map[uint]string)
		db = global.GVA_DB.Model(&autocode.RouteInfo{})
		db.Where("route_name_ex = ?", routeGroup.Identifier).Find(&routeInfos)

		for _, routeInfo := range routeInfos {
			var classInfos []autocode.ClassesInfo
			db = global.GVA_DB.Model(autocode.ClassesInfo{})

			db.Where("route_id = ?", routeInfo.ID).Find(&classInfos)

			for _, classInfo := range classInfos {
				mapBus[*classInfo.BusId] = "ok"
			}
		}

		shuttline := model.ShuttleLine{
			RouteName:  routeGroup.Identifier,
			ShuttleNum: len(mapBus),
		}
		stat.ShuttleLineList = append(stat.ShuttleLineList, shuttline)
	}

	fmt.Printf("\n\nxxxxxxxxxxxx\n%v\n\n", stat)
	//fmt.Println(stat.ShuttleLineList)

	return &stat, nil
}

func Init() {
	global.GVA_LOG.Info("statistic...")
	var ss StatService
	p := &ss
	p.GetShuttleStatistics()
}
