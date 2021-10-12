package gps

import (
    "strconv"
    "fmt"

    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    model "github.com/flipped-aurora/gin-vue-admin/server/model/gps"

    "gorm.io/gorm"
)

type dbBackend struct {
}

func NewBackend() Backend {
    return &dbBackend{}
}

func (d *dbBackend) ListTrack(imei, startTime, endTime string) (interface{}, error) {
    db := global.GVA_DB.Model(&model.GpsDetail{})
    var gdList []model.GpsDetail
    db = db.Where("`gps_sn` = ?", imei)
    db = db.Where("`gps_time` >= ?", startTime)
    db = db.Where("`gps_time` <= ?", endTime)
    err := db.Find(&gdList).Error
    return gdList, err
}

func (d *dbBackend) Save(location []*openapi.Location) {
    if global.GVA_DB == nil {
        return
    }

    global.GVA_DB.Transaction(func(tx *gorm.DB) error {
        for _, loc := range location {
            gd := openapi2model(*loc)
            err := tx.Create(&gd).Error
            if err != nil {
                fmt.Println("--> Save gps location error ", err)
            }
        }
        return nil
    })
}

func openapi2model (loc openapi.Location) model.GpsDetail {
    gd := model.GpsDetail{}
    gd.GpsSn = loc.Imei
    gd.GpsStatus, _ = strconv.Atoi(loc.Status)
    gd.Lng = loc.Lng
    gd.Lat = loc.Lat
    gd.PosType, _ = strconv.Atoi(loc.PosType)
    gd.GpsTime = loc.GpsTime
    gd.Speed, _  = strconv.Atoi(loc.Speed)
    gd.Dir, _ = strconv.Atoi(loc.Direction)
    gd.AccStatus, _ = strconv.Atoi(loc.AccStatus)
    return gd
}
