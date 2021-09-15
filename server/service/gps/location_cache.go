package gps

import (
    "sync"
    "fmt"

    "github.com/flipped-aurora/gin-vue-admin/server/service/gps/openapi"
)

type locationCache struct {
    index map[string]int
    location []*openapi.Location
    mu sync.RWMutex
}

func NewLocationCache() *locationCache {
    return &locationCache{
        index: make(map[string]int),
    }
}

func (lc *locationCache) get(imei string) (openapi.Location, error) {
    lc.mu.RLock()
    defer lc.mu.RUnlock()

    i, ok := lc.index[imei]
    if !ok {
        return openapi.Location{}, fmt.Errorf("could not get location, %s not exist", imei)
    }
    loc := *(lc.location[i])
    return loc, nil
}

func (lc *locationCache) list() ([]openapi.Location, error) {
    lc.mu.RLock()
    defer lc.mu.RUnlock()

    locList := make([]openapi.Location, 0, len(lc.location))
    for _, l := range lc.location {
        locList = append(locList, *l)
    }
    return locList, nil
}

func (lc *locationCache) replace(location []*openapi.Location) []*openapi.Location {
    lc.mu.Lock()
    defer lc.mu.Unlock()

    var diff []*openapi.Location
    for i, loc := range location {
        idx, ok := lc.index[loc.Imei]
        if !ok {
            diff = append(diff, loc)
        } else {
            if lc.location[idx].GpsTime != loc.GpsTime {
                diff = append(diff, loc)
            }
        }
        lc.index[loc.Imei] = i
    }
    lc.location = location
    return diff
}
