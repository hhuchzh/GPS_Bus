package checkin

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gps"
	"github.com/robfig/cron/v3"
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

	//global.GVA_LOG.Info("monitor classes...", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.String("gps_sn", gpsSN))
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

	now := time.Now()
	date := now.Format("2006-01-02")

	if arrival.Location == nil {
		global.GVA_LOG.Warn("monitor arrival, unexcepted loaction")
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", date)
		return
	}

	t := arrival.ArrivalTime
	t = strings.Join([]string{date, t}, " ")
	global.GVA_LOG.Debug("arrival time..", zap.String("arrival time", t))

	st, _ := time.Parse("2006-01-02 15:04:05", t)

	m, _ := time.ParseDuration("-5m")
	start := st.Add(m)
	m, _ = time.ParseDuration("5m")
	end := st.Add(m)

	list, err := monitor.loadGpsDetails(gpsSN, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))

	if err != nil {
		global.GVA_LOG.Error("load gps details failure", zap.Error(err))
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", date)
		return
	}

	gpsList, ok := list.([]gps.GpsDetail)
	if !ok {
		global.GVA_LOG.Error("convert gps failure")
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", date)
		return
	}

	global.GVA_LOG.Info("monitor arrival...", zap.Uint("arrival_id", arrival.ID), zap.String("arrival_time", t), zap.Int("gps_detail_cnt", len(gpsList)))

	var bSuccess bool
	var gps gps.GpsDetail
	var gpsTime time.Time
	for _, gps = range gpsList {
		curLng, _ := strconv.ParseFloat(arrival.Location.Longtitude, 64)
		curLat, _ := strconv.ParseFloat(arrival.Location.Latitude, 64)
		stringSlice := strings.Split(gps.GpsTime, "T")
		gpsTime, _ = time.Parse("2006-01-02 15:04:05", stringSlice[0]+" "+strings.Split(stringSlice[1], "+")[0])

		src := checkinMeta{
			lng: curLng,
			lat: curLat,
			t:   st,
		}

		dst := checkinMeta{
			lng: gps.Lng,
			lat: gps.Lat,
			t:   gpsTime,
		}

		bSuccess = monitor.verify.verifyCheckin(&src, &dst)
		if bSuccess {
			break
		}
	}

	if bSuccess {
		global.GVA_LOG.Info("checkin success")
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, gpsTime.Format("15:04:05"), date)
	} else {
		global.GVA_LOG.Info("can not find checkin")
		_ = monitor.updateCheckin(arrival.ID, *arrival.ClassesId, "", date)
	}

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

func (monitor *CheckinMonitor) loadGpsDetails(gpsSn string, startTime string, endTime string) (list interface{}, err error) {
	global.GVA_LOG.Info("load gps details", zap.String("start", startTime), zap.String("end", endTime))
	db := global.GVA_DB.Model(&gps.GpsDetail{})
	var gdList []gps.GpsDetail
	db = db.Where("`gps_sn` = ?", gpsSn)
	db = db.Where("`gps_time` >= ?", startTime)
	db = db.Where("`gps_time` <= ?", endTime)
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

func (monitor *CheckinMonitor) updateCheckin(arrival uint, classes uint, checkinTime string, checkinDate string) error {

	var checkin autocode.CheckinInfo
	t := time.Now()
	checkin.CreatedAt = t
	checkin.UpdatedAt = t
	checkin.ArrivalId = &arrival
	checkin.ClassesId = &classes
	checkin.CheckinDate = checkinDate
	checkin.CheckinTime = checkinTime
	global.GVA_LOG.Info("update checkin...", zap.Uint("classes_id", classes), zap.Uint("arrival_id", arrival), zap.String("date", checkinDate), zap.String("time", checkinTime))

	return global.GVA_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&checkin).Error
}
