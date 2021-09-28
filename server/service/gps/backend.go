package gps

import (
    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

type Backend interface {
    ListTrack(imei, startTime, endTime string) (interface{}, error)
    Save(location []*openapi.Location)
}
