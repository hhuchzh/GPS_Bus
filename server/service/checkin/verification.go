package checkin

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

const (
	defaultTimeout  = 20 * time.Minute
	defaultDistance = 150.0
)

type verification struct {
	timeout  time.Duration
	distance float64
}

func (veri *verification) setCheckinTimeout(sec time.Duration) error {
	veri.timeout = sec
	return nil
}

func (veri *verification) setCheckinDistance(dis float64) error {
	veri.distance = dis
	return nil
}

func (veri *verification) verifyCheckin(src *checkinMeta, dst *checkinMeta) bool {

	deltaT := dst.t.Sub(src.t)

	if deltaT < 0 {
		deltaT = -deltaT
	}

	if deltaT > veri.timeout {
		return false
	}

	deltaS, _ := getDistance(src.lat, dst.lat, src.lng, dst.lng)

	if deltaS > veri.distance {
		global.GVA_LOG.Info("gps diff exceed", zap.Float64("diff", deltaS))
		return false
	}
	global.GVA_LOG.Info("gps diff ok", zap.Float64("diff", deltaS))
	return true
}

type options struct {
	timeout  time.Duration
	distance float64
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithDistance(dis float64) Option {
	return optionFunc(func(o *options) {
		o.distance = dis
	})
}

func NewVerification(opts ...Option) *verification {
	options := options{
		timeout:  defaultTimeout,
		distance: defaultDistance,
	}

	for _, o := range opts {
		o.apply(&options)
	}

	return &verification{
		timeout:  options.timeout,
		distance: options.distance,
	}
}
