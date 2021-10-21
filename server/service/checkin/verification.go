package checkin

import (
	"time"
)

const (
	defaultTimeout  = 5 * time.Minute
	defaultDistance = 100.0
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
		return false
	}

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
