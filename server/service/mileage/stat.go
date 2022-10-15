package mileage

import (
	"math"
	"sort"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gps"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

const (
        cronJobSpec = "0 0 23 * * *"
	//cronJobSpec = "*/120 * 8,11 * * *"   //every 120s in 8,11
	//cronJobSpec = "0 0 16,17 * * ?"       // run in 16, 17
	//cronJobSpec = "0 0/15 * * * ?" //run every half hour
)

type MileageStat struct {
	src  *Source
	cron *cron.Cron
}

func NewMileageStat() *MileageStat {
	return &MileageStat{
		src:  &Source{},
		cron: cron.New(cron.WithSeconds()),
	}
}

func Init() {
	global.GVA_LOG.Info("mileage stat initialize...")
	ms := NewMileageStat()
	ms.Start()
}

func (ms *MileageStat) Start() {
	_, err := ms.cron.AddFunc(cronJobSpec, func() {
		list, cnt, err := ms.src.LoadClasses()
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
			ms.statClass(&classInfo)
		}
	})
	if err != nil {
		global.GVA_LOG.Error("MileageStat start crontab failed")
		return
	}
	ms.cron.Start()
}

func (ms *MileageStat) statClass(classes *autocode.ClassesInfo) {
	gpsSN, err := ms.src.LoadGpsSn(*classes.BusId)
	if err != nil {
		global.GVA_LOG.Warn("monitor classes..., can not load gps", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.Error(err))
		return
	}

	list, err := ms.src.LoadArrivals(uint64(classes.ID))
	if err != nil {
		global.GVA_LOG.Info("monitor classes..., can not load arrival", zap.Uint("classes_id", classes.ID), zap.Error(err))
		return
	}

	arrivals, ok := list.([]autocode.ArrivalInfo)
	if !ok {
		global.GVA_LOG.Error("convert to arrival failure")
		return
	}
	var arrivalsTime []string
	for _, a := range arrivals {
		arrivalsTime = append(arrivalsTime, a.ArrivalTime)
	}
	sort.Strings(arrivalsTime)

	nTime := time.Now()
	//yesTime := nTime.AddDate(0, 0, -day)
	date := nTime.Format("2006-01-02")

	checklist, err := ms.src.LoadCheckIn(uint64(classes.ID), date)
	if err != nil {
		global.GVA_LOG.Info("monitor classes..., can not load checkin data", zap.Uint("classes_id", classes.ID), zap.Error(err))
		return
	}
	checkins, ok := checklist.([]autocode.CheckinInfo)
	if !ok {
		global.GVA_LOG.Error("convert to checkin failure")
		return
	}

	global.GVA_LOG.Info("#### start checkin data", zap.Uint("classes_id", classes.ID), zap.String("date", date), zap.Error(err))

	var checkinTimes []string
	for _, a := range checkins {
		r_len := len(a.Reason)
		if r_len == 0 {
			global.GVA_LOG.Info("add time a.CheckinTime ", zap.String("reason", a.CheckinTime))
			checkinTimes = append(checkinTimes, a.CheckinTime)
		} else if r_len > 0 && r_len < 10 {
			global.GVA_LOG.Info("add time a.Reason", zap.String("reason", a.Reason))
			checkinTimes = append(checkinTimes, a.Reason)
		} else {
			global.GVA_LOG.Info("check in failed,not add time")
			continue
		}
	}

	checkinTimes = append(checkinTimes, arrivalsTime[len(arrivalsTime)-1])

	sort.Strings(checkinTimes)

	t_len := len(checkinTimes)
	if t_len <= 1 {
		return
	}

	global.GVA_LOG.Info("monitor classes...", zap.Uint("classes_id", classes.ID), zap.Uint("bus_id", *classes.BusId), zap.String("gps_sn", gpsSN), zap.Int("arrival_cnt", len(checkinTimes)))
	ms.statArrival(classes.ID, gpsSN, checkinTimes[0], checkinTimes[len(checkinTimes)-1])
}

func (ms *MileageStat) statArrival(classesid uint, gpsSN string, beginTime, endTime string) {
	now := time.Now()
	date := now.Format("2006-01-02")

	//nTime := time.Now()
	//yesTime := nTime.AddDate(0, 0, -day)
	//date := yesTime.Format("2006-01-02")

	t1 := beginTime
	t1 = strings.Join([]string{date, t1}, " ")
	t2 := endTime
	t2 = strings.Join([]string{date, t2}, " ")
	global.GVA_LOG.Debug("arrival time..", zap.String("arrival begin time", t1), zap.String("arrival end time", t2))
	st1, _ := time.Parse("2006-01-02 15:04:05", t1)
	st2, _ := time.Parse("2006-01-02 15:04:05", t2)
	m1, _ := time.ParseDuration("-2m")
	start := st1.Add(m1)
	m2, _ := time.ParseDuration("2m")
	end := st2.Add(m2)

	list, err := ms.src.LoadGpsDetails(gpsSN, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))

	if err != nil {
		global.GVA_LOG.Error("load gps details failure", zap.Error(err))
		_ = ms.updateMileage(classesid, 0, date)
	}

	gpsList, ok := list.([]gps.GpsDetail)
	if !ok {
		global.GVA_LOG.Error("convert gps failure")
		_ = ms.updateMileage(classesid, 0, date)
	}
	if len(gpsList) < 2 {
		_ = ms.updateMileage(classesid, 0, date)
		return
	}

	prevLat := gpsList[0].Lat
	prevLng := gpsList[0].Lng
	gpsList = gpsList[1:]
	var totalDist float64
	for _, gps := range gpsList {
		dist := GetDistance(prevLat, prevLng, gps.Lat, gps.Lng)
		prevLat = gps.Lat
		prevLng = gps.Lng
		if math.IsNaN(dist) {
			continue
		}
		totalDist += dist
	}
	_ = ms.updateMileage(classesid, totalDist, date)
}

func (ms *MileageStat) updateMileage(classes uint, dist float64, date string) error {
	var mileage autocode.MilesInfo
	t := time.Now()
	mileage.CreatedAt = t
	mileage.UpdatedAt = t
	mileage.ClassesId = &classes
	if classes >= 142 && classes <= 167 && dist > 2000 {
		dist += 3000
	}
	mileage.CalcDate = date
	mileage.Distance = dist

	global.GVA_LOG.Info("update miles...", zap.Uint("classes_id", classes), zap.String("date", date), zap.Float64("distance", dist))

	return global.GVA_DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&mileage).Error
}
