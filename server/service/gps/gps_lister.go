package gps

import (
    "fmt"

    "github.com/robfig/cron/v3"

    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

const (
    cronJobSpec = "*/15 * 7,8,9,15,16,17,18 * * *"
)

var (
    GpsLister = NewGPSLister()
)

type GPSLister struct {
    client *openapi.Client
    lc locationCache
    stopCh chan struct{}
    cron *cron.Cron
}

func NewGPSLister() *GPSLister {
    return &GPSLister{
        client: openapi.NewClient(),
        stopCh: make(chan struct{}),
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

func (g *GPSLister) Start() {
    fmt.Println("****** GPSLister Start")
    g.cron = cron.New(cron.WithSeconds())
    _, err := g.cron.AddFunc(cronJobSpec, func() {
        location, err := g.client.ListLocation()
        if err != nil {
            fmt.Println("GPSLister client.ListLocation failed: ", err)
        } else {
            g.lc.replace(location)
        }
    })
    if err != nil {
        fmt.Println("GPSLister Start crontab failed: ", err)
        return
    }
    g.cron.Start()
    g.client.Start()
}

func (g *GPSLister) Stop() {
    g.client.Stop()
    g.cron.Stop()
}
