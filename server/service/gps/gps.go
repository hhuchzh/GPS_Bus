package gps

import (
    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

type GpsService struct {
}

func (gs *GpsService) ListLocation() ([]openapi.Location, error) {
	return GpsLister.ListLocation()
}

func (gs *GpsService) GetLocation(imei string) (openapi.Location, error) {
	return GpsLister.GetLocation(imei)
}

func (gs *GpsService) ListDevice() ([]openapi.Device, error) {
	return GpsLister.ListDevice()
}

func (gs *GpsService) ListTrack(imei, beginTime, endTime string) ([]openapi.Track, error) {
	return GpsLister.ListTrack(imei, beginTime, endTime)
}
