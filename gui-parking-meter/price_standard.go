package main

import (
    "fmt"
)

type PricePart struct {
    StartTime int // 8
    EndTime int // 20
    Price float64 // 0.5
    Unit int // minute 15
    Max int //minute 6 * 60
}

type PriceStandard struct {
    FreeTime int // minute 15
    Day PricePart
    Night PricePart
}

type priceConfig struct {
	config map[string]map[string]PriceStandard
}

var (
	PriceConfig = &priceConfig{}
)

func (pc *priceConfig) GetPrice(region, vehicleType string) (PriceStandard, error) {
	ps, ok := pc.config[region]
	if !ok {
		return PriceStandard{}, fmt.Errorf("所选区域不支持")
	}
	p, ok := ps[vehicleType]
	if !ok {
		return PriceStandard{}, fmt.Errorf("所选车型不支持")
	}
	return p, nil
}

func (pc *priceConfig) Init() {
	pc.config = make(map[string]map[string]PriceStandard)
	ps := make(map[string]PriceStandard)
	ps["小型车"] = PriceStandard{
		FreeTime: 15,
        Day: PricePart{
            StartTime: 8,
            EndTime: 20,
            Price: 0.5,
            Unit: 15,
            Max: 345,
        },
        Night: PricePart{
            StartTime: 20,
            EndTime: 8,
            Price: 1,
            Unit: 60,
            Max: 345,
        },
	}
	ps["大型车"] = PriceStandard{
		FreeTime: 15,
        Day: PricePart{
            StartTime: 8,
            EndTime: 20,
            Price: 1.5,
            Unit: 15,
            Max: 345,
        },
        Night: PricePart{
            StartTime: 20,
            EndTime: 8,
            Price: 1.5,
            Unit: 60,
            Max: 345,
        },
	}
	pc.config["三级区域公共停车场"] = ps
	
	ps = make(map[string]PriceStandard)
	ps["小型车"] = PriceStandard{
		FreeTime: 15,
        Day: PricePart{
            StartTime: 8,
            EndTime: 20,
            Price: 1,
            Unit: 15,
            Max: 345,
        },
        Night: PricePart{
            StartTime: 20,
            EndTime: 8,
            Price: 1,
            Unit: 60,
            Max: 345,
        },
	}
	ps["大型车"] = PriceStandard{
		FreeTime: 15,
        Day: PricePart{
            StartTime: 8,
            EndTime: 20,
            Price: 2,
            Unit: 15,
            Max: 345,
        },
        Night: PricePart{
            StartTime: 20,
            EndTime: 8,
            Price: 1.5,
            Unit: 60,
            Max: 345,
        },
	}
	pc.config["二级区域公共停车场"] = ps
}
