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

func (lc *locationCache) replace(location []*openapi.Location) {
    lc.mu.Lock()
    defer lc.mu.Unlock()

    lc.location = location
    lc.index = make(map[string]int, len(location))
    for i, loc := range location {
        lc.index[loc.Imei] = i
    }
}
