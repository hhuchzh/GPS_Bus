package main

import (
    "time"
    _ "time/tzdata"
    "fmt"
)

type Calculator struct {
}

func (c *Calculator) calcPart(start, end, endPart time.Time, p PricePart) (float64, bool) {
    var nextTime time.Time
    var exit bool
    if end.Before(endPart) {
        nextTime = end
        exit = true
    } else {
        nextTime = endPart
    }
    duration := nextTime.Sub(start)
    var sec int
    if duration.Minutes() > float64(p.Max) {
        sec = p.Max * 60
    } else {
        sec = int(duration.Seconds())
    }
    unitS := p.Unit * 60
    count := sec / unitS
    if sec % unitS != 0 {
        count++
    }
    return float64(count) * p.Price, exit
}

func (c *Calculator) Calc(startTime, endTime string, price PriceStandard) (float64, error) {
    start, err := toTime(startTime)
    if err != nil {
        return -1, err
    }
    end, err := toTime(endTime)
    if err != nil {
        return -1, err
    }

    freeUnit := fmt.Sprintf("%dm", price.FreeTime)
    freeDuration, _ := time.ParseDuration(freeUnit)
    t := start.Add(freeDuration)
    if end.Before(t) {
        return 0, nil
    }

    var totalPrice float64
    location, err := time.LoadLocation("Asia/Shanghai")
    if err != nil {
        return -1, err
    }
    for {
        isDay := c.inTimeSection(price.Day.StartTime, price.Day.EndTime, t.Hour())
        if isDay {
            dayEndTime := time.Date(t.Year(), t.Month(), t.Day(), price.Day.EndTime, 0, 0, 0, location)
            p, exit := c.calcPart(t, end, dayEndTime, price.Day)
            totalPrice += p
            if exit {
                break
            }
            t = dayEndTime
        } else {
            day := t.Day()
            if t.Hour() >= price.Night.StartTime {
                day += 1
            }
            nightEndTime := time.Date(t.Year(), t.Month(), day, price.Night.EndTime, 0, 0, 0, location)
            p, exit := c.calcPart(t, end, nightEndTime, price.Night)
            totalPrice += p
            if exit {
                break
            }
            t = nightEndTime
        }
    }

    return totalPrice, nil
}

func (c *Calculator) inTimeSection(start, end int, hour int) bool {
    if (start < end) {
        if hour >= start && hour < end {
            return true
        }
        return false
    } else {
        if hour >= start || hour < end {
            return true
        }
        return false
    }
}

func toTime(v string) (time.Time, error) {
    location, err := time.LoadLocation("Asia/Shanghai")
    if err != nil {
        return time.Time{}, err
    }
    return time.ParseInLocation("2006-01-02 15:04:05", v, location)
}
