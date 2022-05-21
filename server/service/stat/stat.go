package stat

import (
	model "github.com/flipped-aurora/gin-vue-admin/server/model/stat"
)

type StatService struct {
}

func (ss *StatService) GetShuttleStatistics() ([]model.StatInfo, error) {
	return nil, nil
}
