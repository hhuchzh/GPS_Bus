package gps

import (
    //"time"
)

type GpsDetail struct {
    GpsSn  string `json:"gpsSn" form:"gpsSn" gorm:"column:gps_sn;comment:GPS序列号;type:varchar(100);"`
    GpsStatus int `json:"gpsStatus" form:"gpsStatus" gorm:"column:gps_status;comment:;type:tinyint"`
    Lng float64 `json:"lng" form:"lng" gorm:"column:lng;comment:;type:decimal"`
    Lat float64 `json:"lat" form:"lat" gorm:"column:lat;comment:;type:decimal"`
    PosType int `json:"posType" form:"posType" gorm:"column:pos_type;comment:;type:tinyint"`
    GpsTime string `json:"gpsTime" form:"gpsTime" gorm:"column:gps_time;comment:;type:datetime"`
    Speed int `json:"speed" form:"speed" gorm:"column:speed;comment:;type:smallint"`
    Dir int `json:"dir" form:"dir" gorm:"column:dir;comment:;type:smallint"`
    AccStatus int `json:"accStatus" form:"accStatus" gorm:"column:acc_status;comment:;type:smallint"`
}

func (GpsDetail) TableName() string {
  return "gps_detail"
}
