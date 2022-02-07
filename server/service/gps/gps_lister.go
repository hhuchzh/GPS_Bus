package gps

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

const (
	cronJobSpec = "*/11 * 7,8,9,10,11,12,13,14,15,16,17,18,19,20,21 * * *"
)

var (
	GpsLister = NewGPSLister()
)

func init() {
	GpsLister.Start()
}

type GPSLister struct {
	client  *openapi.Client
	lc      *locationCache
	stopCh  chan struct{}
	cron    *cron.Cron
	backend Backend
}

func NewGPSLister() *GPSLister {
	return &GPSLister{
		client:  openapi.NewClient(),
		lc:      NewLocationCache(),
		stopCh:  make(chan struct{}),
		backend: NewBackend(),
	}
}

func (g *GPSLister) GetLocation(imei string) (openapi.Location, error) {
	return g.lc.get(imei)
}

func (g *GPSLister) ListLocation() ([]openapi.Location, error) {
	return g.lc.list()
}

func (g *GPSLister) ListDevice() ([]openapi.Device, error) {
	device, err := g.client.ListDevice()
	if err != nil {
		return nil, err
	}

	dList := make([]openapi.Device, 0, len(device))
	for _, d := range device {
		dList = append(dList, *d)
	}
	return dList, nil
}

func (g *GPSLister) ListTrack(imei, beginTime, endTime string) (interface{}, error) {
	return g.backend.ListTrack(imei, beginTime, endTime)
}

func (g *GPSLister) Start() {
	fmt.Println("****** GPSLister Start")

	g.cron = cron.New(cron.WithSeconds())
	_, err := g.cron.AddFunc(cronJobSpec, func() {
		location, err := g.client.ListLocation()
		if err != nil {
			global.GVA_LOG.Error("GPSLister client.ListLocation failed: ", zap.Error(err))
		} else {
			diff := g.lc.replace(location)
			g.backend.Save(diff)
		}
	})
	if err != nil {
		global.GVA_LOG.Error("GPSLister Start crontab failed: ", zap.Error(err))
		return
	}
	g.cron.Start()
	g.client.Start()
}

func (g *GPSLister) Stop() {
	g.client.Stop()
	g.cron.Stop()
}
