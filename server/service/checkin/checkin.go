package checkin

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gps"
	service "github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"github.com/robfig/cron/v3"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

var monitor *CheckinMonitor

const (
	cronJobSpec = "0 0 22 * * ?"
)

type checkinMeta struct {
	lng float64
	lat float64
	t   time.Time
}

type CheckinService struct {
}

type checkin interface {
	setCheckinTimeout(sec time.Duration) error
	setCheckinDistance(dis int) error
	verifyCheckin(src checkinMeta, dst checkinMeta) (bool, error)
}

func getDistance(lat1, lat2, lng1, lng2 float64) (float64, error) {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius, nil
}

type CheckinMonitor struct {
	cron   *cron.Cron
	verify *verification
}

// xxx
func Init() {
	global.GVA_LOG.Info("checkin monitor initialize...")
	monitor = NewMonitor()
	monitor.Run()
}

// NewMonitor xxx
func NewMonitor() *CheckinMonitor {

	global.GVA_LOG.Info("Run checkin monitor...")
	verify := NewVerification()

	cron := cron.New(cron.WithSeconds())

	return &CheckinMonitor{
		cron:   cron,
		verify: verify,
	}

}

// func (monitor *CheckinMonitor) Run() {
// 	global.GVA_LOG.Info("Create checkin monitor...")

// 	list, cnt, err := monitor.loadClassses()
// 	if err != nil {
// 		global.GVA_LOG.Error("load classes failure", zap.Error(err))
// 		return
// 	}
// 	global.GVA_LOG.Info("classes is loaded...", zap.Int64("total cnt", cnt))

// 	classes, ok := list.([]autocode.ClassesInfo)
// 	if !ok {
// 		global.GVA_LOG.Error("convert to ClassesInfo failure")
// 		return
// 	}

// 	for _, classInfo := range classes {
// 		monitor.monitorClasses(&classInfo)
// 	}

// 	if err != nil {
// 		fmt.Println("GPSLister Start crontab failed: ", err)
// 		return
// 	}
// 	//monitor.cron.Start()

// }

// Run xxx
func (monitor *CheckinMonitor) Run() {
	global.GVA_LOG.Info("Create checkin monitor...")
	_, err := monitor.cron.AddFunc(cronJobSpec, func() {
		list, cnt, err := monitor.loadClassses()
		if err != nil {
			global.GVA_LOG.Error("load classes failure", zap.Error(err))
			return
		}
		global.GVA_LOG.Info("classes is loaded...", zap.Int64("total cnt", cnt))

		classes, ok := list.([]autocode.ClassesInfo)
		if !ok {
			global.GVA_LOG.Error("convert to ClassesInfo failure")
			return
		}

		for _, classInfo := range classes {
			monitor.monitorClasses(&classInfo)
		}

	})
	if err != nil {
		fmt.Println("GPSLister Start crontab failed: ", err)
		return
	}
	monitor.cron.Start()

}

func (monitor *CheckinMonitor) monitorClasses(classes *autocode.ClassesInfo) {
	global.GVA_LOG.Debug("monitor classes start")

	gpsSN, err := monitor.loadGpsSn(*classes.BusId)
	if err != nil {
		global.GVA_LOG.Warn("monitor classes..., can not load gps", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.Error(err))
		return
	}

	global.GVA_LOG.Info("monitor classes...", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.String("gps_sn", gpsSN))
	list, err := monitor.loadArrivals(uint64(classes.ID))
	if err != nil {
		global.GVA_LOG.Info("monitor classes..., can not load arrival", zap.Uint("classes_id", classes.ID), zap.Error(err))
		return
	}

	arrivals, ok := list.([]autocode.ArrivalInfo)
	if !ok {
		global.GVA_LOG.Error("convert to arrival failure")
		return
	}

	global.GVA_LOG.Info("monitor classes...", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.String("gps_sn", gpsSN), zap.Int("arrival_cnt", len(arrivals)))

	for _, arrival := range arrivals {
		fmt.Println(arrival)
		monitor.monitorArrival(gpsSN, &arrival)
	}
}

func (monitor *CheckinMonitor) monitorArrival(gpsSN string, arrival *autocode.ArrivalInfo) {

	global.GVA_LOG.Info("monitor arrival...", zap.Uint("arrival_id", arrival.ID), zap.Uint("location_id", *arrival.LocationId))

	now := time.Now().Add(-24 * time.Hour)
	//now := time.Now()
	date := now.Format("2006-01-02")

	if arrival.Location == nil {
		global.GVA_LOG.Warn("monitor arrival, unexcepted loaction")
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "Unexcepted Error", date)
		return
	}

	t := arrival.ArrivalTime
	t = strings.Join([]string{date, t}, " ")
	global.GVA_LOG.Debug("arrival time..", zap.String("arrival time", t))

	st, _ := time.Parse("2006-01-02 15:04:05", t)

	var bSuccess bool
	var gpsVar gps.GpsDetail
	var gpsTime time.Time
	bSuccess = false

	for i := 0; i < 20; i++ {
		start := st.Add(time.Duration(-1*(i+1)) * time.Minute)
		end := st.Add(time.Duration(-1*i) * time.Minute)
		list, err := monitor.loadGpsDetails(gpsSN, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"), false)

		if err != nil {
			global.GVA_LOG.Error("load gps details failure", zap.Error(err))
			_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "Unexcepted Error", date)
			return
		}

		gpsList, ok := list.([]gps.GpsDetail)
		if !ok {
			global.GVA_LOG.Error("convert gps failure")
			_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "Unexcepted Error", date)
			return
		}

		global.GVA_LOG.Info("monitor arrival...", zap.Uint("arrival_id", arrival.ID), zap.String("arrival_time", t), zap.Int("gps_detail_cnt", len(gpsList)))

		bSuccess = false
		for _, gpsVar = range gpsList {
			curLng, _ := strconv.ParseFloat(arrival.Location.Longtitude, 64)
			curLat, _ := strconv.ParseFloat(arrival.Location.Latitude, 64)

			stringSlice := strings.Split(gpsVar.GpsTime, "T")
			gpsTime, _ = time.Parse("2006-01-02 15:04:05", stringSlice[0]+" "+strings.Split(stringSlice[1], "+")[0])

			src := checkinMeta{
				lng: curLng,
				lat: curLat,
				t:   st,
			}

			dst := checkinMeta{
				lng: gpsVar.Lng,
				lat: gpsVar.Lat,
				t:   gpsTime,
			}

			bSuccess = monitor.verify.verifyCheckin(&src, &dst)
			if bSuccess {
				if i < 5 {
					global.GVA_LOG.Info("checkin success")
					_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, gpsTime.Format("15:04:05"), "", date)
				} else {
					global.GVA_LOG.Info("checkin in over 5 minutes", zap.Int("minute", -i))
					_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", gpsTime.Format("15:04:05"), date)
				}
				return
			}
		}

		start = st.Add(time.Duration(1*i) * time.Minute)
		end = st.Add(time.Duration(1*(i+1)) * time.Minute)
		list, err = monitor.loadGpsDetails(gpsSN, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"), true)

		if err != nil {
			global.GVA_LOG.Error("load gps details failure", zap.Error(err))
			_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "Unexcepted Error", date)
			return
		}

		gpsList, ok = list.([]gps.GpsDetail)
		if !ok {
			global.GVA_LOG.Error("convert gps failure")
			_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "Unexcepted Error", date)
			return
		}

		global.GVA_LOG.Info("monitor arrival...", zap.Uint("arrival_id", arrival.ID), zap.String("arrival_time", t), zap.Int("gps_detail_cnt", len(gpsList)))

		bSuccess = false
		for _, gpsVar = range gpsList {
			curLng, _ := strconv.ParseFloat(arrival.Location.Longtitude, 64)
			curLat, _ := strconv.ParseFloat(arrival.Location.Latitude, 64)

			stringSlice := strings.Split(gpsVar.GpsTime, "T")
			gpsTime, _ = time.Parse("2006-01-02 15:04:05", stringSlice[0]+" "+strings.Split(stringSlice[1], "+")[0])

			src := checkinMeta{
				lng: curLng,
				lat: curLat,
				t:   st,
			}

			dst := checkinMeta{
				lng: gpsVar.Lng,
				lat: gpsVar.Lat,
				t:   gpsTime,
			}

			bSuccess = monitor.verify.verifyCheckin(&src, &dst)
			if bSuccess {
				if i < 5 {
					global.GVA_LOG.Info("checkin success")
					_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, gpsTime.Format("15:04:05"), "", date)
				} else {
					global.GVA_LOG.Info("checkin in over 5 minutes", zap.Int("minute", i))
					_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", gpsTime.Format("15:04:05"), date)
				}
				return
			}
		}

	}

	_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", "no checkin -/+ 20min", date)

}

func (monitor *CheckinMonitor) loadGpsSn(busId uint) (string, error) {
	var gps autocode.GpsInfo
	result := global.GVA_DB.Where("bus_id = ?", busId).Find(&gps)

	var cnt int64
	_ = result.Count(&cnt)
	if cnt == 0 {
		return "", errors.New("can not get gps")
	}

	return gps.GpsSn, result.Error

}

func (monitor *CheckinMonitor) loadGpsDetails(gpsSn string, startTime string, endTime string, order bool) (list interface{}, err error) {
	global.GVA_LOG.Info("load gps details", zap.String("start", startTime), zap.String("end", endTime), zap.Bool("order", order))
	orderStr := "gps_time "
	if order {
		orderStr += "asc"
	} else {
		orderStr += "desc"
	}

	db := global.GVA_DB.Model(&gps.GpsDetail{})
	var gdList []gps.GpsDetail
	db = db.Where("`gps_sn` = ?", gpsSn)
	db = db.Where("`gps_time` >= ?", startTime)
	db = db.Where("`gps_time` <= ?", endTime)
	db = db.Order(orderStr)
	err = db.Find(&gdList).Error
	return gdList, err
}

func (monitor *CheckinMonitor) loadClassses() (list interface{}, cnt int64, err error) {

	db := global.GVA_DB.Model(&autocode.ClassesInfo{})
	err = db.Count(&cnt).Error
	global.GVA_LOG.Info("load classes...", zap.Int64("total cnt", cnt))
	var classesInfos []autocode.ClassesInfo
	err = db.Find(&classesInfos).Error

	return classesInfos, cnt, err
}

func (monitor *CheckinMonitor) loadArrivals(id uint64) (list interface{}, err error) {
	var arrivals []autocode.ArrivalInfo
	result := global.GVA_DB.Where("classes_id = ?", id).Preload("Location").Find(&arrivals)
	var cnt int64
	_ = result.Count(&cnt)
	global.GVA_LOG.Info("load arrivals ...", zap.Int64("cnt", cnt))
	err = result.Error
	return arrivals, err
}

func (monitor *CheckinMonitor) updateCheckin(arrival uint, classes uint, checkinTime string, reason string, checkinDate string) error {

	var checkin autocode.CheckinInfo
	t := time.Now()
	checkin.CreatedAt = t
	checkin.UpdatedAt = t
	checkin.ArrivalId = &arrival
	checkin.ClassesId = &classes
	checkin.CheckinDate = checkinDate
	checkin.CheckinTime = checkinTime
	checkin.Reason = reason
	global.GVA_LOG.Info("update checkin...", zap.Uint("classes_id", classes), zap.Uint("arrival_id", arrival), zap.String("date", checkinDate), zap.String("time", checkinTime), zap.String("reason", reason))

	return global.GVA_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&checkin).Error
}

func (monitor *CheckinMonitor) loadBusIDByPlate(plate string) (uint, error) {
	var bus autocode.BusInfo
	result := global.GVA_DB.Where("bus_plate = ?", plate).Find(&bus)

	var cnt int64
	_ = result.Count(&cnt)
	if cnt == 0 {
		return 0, errors.New("can not get busid")
	}
	return bus.ID, result.Error
}

func (monitor *CheckinMonitor) loadClassByBusID(busID uint) ([]autocode.ClassesInfo, error) {
	var classes []autocode.ClassesInfo
	result := global.GVA_DB.Where("bus_id = ?", busID).Find(&classes)

	var cnt int64
	_ = result.Count(&cnt)
	if cnt == 0 {
		return nil, errors.New("can not get class")
	}
	return classes, result.Error
}

func (monitor *CheckinMonitor) loadRouteByID(id uint) (autocode.RouteInfo, error) {
	var route autocode.RouteInfo
	result := global.GVA_DB.Where("id = ?", id).Find(&route)

	var cnt int64
	_ = result.Count(&cnt)
	if cnt == 0 {
		return autocode.RouteInfo{}, errors.New("can not get route")
	}
	return route, result.Error
}

func (monitor *CheckinMonitor) loadCheckinInfos(ids []uint, start string, end string) ([]autocode.CheckinInfo, error) {
	var checkinData []autocode.CheckinInfo
	err := global.GVA_DB.Where("classes_id in ?", ids).Where("checkin_date >= ?", start).Where("checkin_date <= ?", end).Order("checkin_date").Order("classes_id").Find(&checkinData).Error

	return checkinData, err
}

// ExportExcel ...
func (monitor *CheckinMonitor) ExportExcel(plate string, start string, end string) (string, error) {
	busID, err := monitor.loadBusIDByPlate(plate)
	if err != nil {
		return "", err
	}
	classes, err := monitor.loadClassByBusID(busID)
	if err != nil {
		return "", err
	}

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 1), "日期")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 2), "车牌号")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 3), "路线ID")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 4), "路线名")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 5), "班次ID")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 6), "班次")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 7), "发车时间")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 8), "站点名字")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 9), "站点时间")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 10), "打卡时间")
	f.SetCellValue("Sheet1", service.FormatCoord(1, 11), "原因")

	var classesIDs []uint
	classesMap := make(map[uint]autocode.ClassesInfo)
	for _, class := range classes {
		classesIDs = append(classesIDs, class.ID)
		classesMap[class.ID] = class
	}

	checkins, err := monitor.loadCheckinInfos(classesIDs, start, end)
	if err != nil {
		return "", err
	}
	row := 2
	for _, checkinItem := range checkins {
		classItem := classesMap[*checkinItem.ClassesId]
		route, err := monitor.loadRouteByID(*classItem.RouteId)
		if err != nil {
			continue
		}
		stringSlice := strings.Split(checkinItem.CheckinDate, "T")
		checkinData := stringSlice[0]

		f.SetCellValue("Sheet1", service.FormatCoord(row, 1), checkinData)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 2), plate)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 3), route.ID)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 4), route.RouteName)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 5), classItem.ID)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 6), classItem.Remark)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 7), classItem.ClassesTime)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 8), checkinItem.Arrival.Location.LocationName)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 9), checkinItem.Arrival.ArrivalTime)

		var checkinTime string = checkinItem.CheckinTime
		if checkinTime == "00:00:00" {
			checkinTime = "打卡异常"
		}
		f.SetCellValue("Sheet1", service.FormatCoord(row, 10), checkinTime)
		f.SetCellValue("Sheet1", service.FormatCoord(row, 11), checkinItem.Reason)
		row++
	}

	f.SetActiveSheet(index)
	path, _ := os.Getwd()
	fileName := fmt.Sprintf("%s/%s_checkin_%d.xlsx", path, plate, time.Now().Unix())
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}
	return fileName, nil
}
