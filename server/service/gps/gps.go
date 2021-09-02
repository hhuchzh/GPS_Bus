package gps

import (
    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

type GpsService struct {
}

func (gs *GpsService) ListLocation() ([]openapi.Location, error) {
	return GpsLister.ListLocation()
}
