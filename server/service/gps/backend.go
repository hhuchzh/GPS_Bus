package gps

import (
    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
    model "github.com/flipped-aurora/gin-vue-admin/server/model/gps"
)

type Backend interface {
    ListTrack(imei, startTime, endTime string) ([]model.GpsDetail, error)
    Save(location []*openapi.Location)
}
