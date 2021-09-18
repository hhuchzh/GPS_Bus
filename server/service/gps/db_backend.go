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

func (d *dbBackend) ListTrack(imei, startTime, endTime string) ([]model.GpsDetail, error) {
    return nil, nil
}

func (d *dbBackend) Save(location []*openapi.Location) {
    if global.GVA_DB == nil {
        return
    }

    global.GVA_DB.Transaction(func(tx *gorm.DB) error {
        for _, loc := range location {
            gd := model.GpsDetail{}
            gd.GpsSn = loc.Imei
            gd.GpsStatus, _ = strconv.Atoi(loc.Status)
            //gd.Lng = strconv.FormatFloat(loc.Lng, 'G', 50, 64)
            //gd.Lat = strconv.FormatFloat(loc.Lat, 'G', 50, 64)
            gd.Lng = loc.Lng
            gd.Lat = loc.Lat
            gd.PosType, _ = strconv.Atoi(loc.PosType)
            gd.GpsTime = loc.GpsTime
            gd.Speed, _  = strconv.Atoi(loc.Speed)
            gd.Dir, _ = strconv.Atoi(loc.Direction)
            err := tx.Create(&gd).Error
            if err != nil {
                fmt.Println("--> Save gps location error ", err)
            }
        }
        return nil
    })
}
