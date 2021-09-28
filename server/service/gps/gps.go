package gps

import (
    "fmt"

    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
    model "github.com/flipped-aurora/gin-vue-admin/server/model/gps"
)

type GpsService struct {
}

func (gs *GpsService) ListLocation() ([]model.GpsDetail, error) {
    var gdList []model.GpsDetail
    locList, err := GpsLister.ListLocation()
    if err != nil {
        return nil, err
    }
    for _, loc := range locList {
        gd := openapi2model(loc)
        gdList = append(gdList, gd)
    }
    return gdList, nil
}

func (gs *GpsService) GetLocation(imei string) (model.GpsDetail, error) {
    loc, err := GpsLister.GetLocation(imei)
    if err != nil {
        return model.GpsDetail{}, err
    }
    return openapi2model(loc), nil
}

func (gs *GpsService) ListDevice() ([]openapi.Device, error) {
	return GpsLister.ListDevice()
}

func (gs *GpsService) ListTrack(imei, beginTime, endTime string) ([]model.GpsDetail, error) {
    list, err := GpsLister.ListTrack(imei, beginTime, endTime)
    if err != nil {
        return nil, err
    }
    gdList, ok := list.([]model.GpsDetail)
    if !ok {
        return nil, fmt.Errorf("--> ListTrack type error")
    }
    return gdList, nil
}

