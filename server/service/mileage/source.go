package mileage

import (
    "errors"

    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/gps"
	"go.uber.org/zap"
)

type Source struct {
}

func (s *Source) LoadClasses() (list interface{}, cnt int64, err error) {
    db := global.GVA_DB.Model(&autocode.ClassesInfo{})
    err = db.Count(&cnt).Error
    global.GVA_LOG.Info("load classes...", zap.Int64("total cnt", cnt))
    var classesInfos []autocode.ClassesInfo
    err = db.Find(&classesInfos).Error
    return classesInfos, cnt, err
}

func (s *Source) LoadGpsSn(busId uint) (string, error) {
    var gps autocode.GpsInfo
    result := global.GVA_DB.Where("bus_id = ?", busId).Find(&gps)

    var cnt int64
    _ = result.Count(&cnt)
    if cnt == 0 {
        return "", errors.New("can not get gps")
    }
    return gps.GpsSn, result.Error
}

func (s *Source) LoadGpsDetails(gpsSn, startTime, endTime string) (list interface{}, err error) {
    global.GVA_LOG.Info("load gps details", zap.String("start", startTime), zap.String("end", endTime))
    db := global.GVA_DB.Model(&gps.GpsDetail{})
    var gdList []gps.GpsDetail
    db = db.Where("`gps_sn` = ?", gpsSn)
    db = db.Where("`gps_time` >= ?", startTime)
    db = db.Where("`gps_time` <= ?", endTime)
    err = db.Find(&gdList).Error
    return gdList, err
}

func (s *Source) LoadArrivals(id uint64) (list interface{}, err error) {
    var arrivals []autocode.ArrivalInfo
    result := global.GVA_DB.Where("classes_id = ?", id).Find(&arrivals)
    var cnt int64
    _ = result.Count(&cnt)
    global.GVA_LOG.Info("load arrivals ...", zap.Int64("cnt", cnt))
    if cnt == 0 {
        return nil, errors.New("can not get arrivals")
    }
    err = result.Error
    return arrivals, err
}
