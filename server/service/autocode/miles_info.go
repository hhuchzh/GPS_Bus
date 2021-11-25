package autocode

import (
    "os"
    "time"
    "fmt"
    "errors"
    "strings"
    "sort"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"

    "github.com/xuri/excelize/v2"
)

type MilesInfoService struct {
}

type exportData struct {
    date string
    classId uint
    routeName string
    classTime string
    distance float64
}

func (s *MilesInfoService) GetMilesInfoList(info autoCodeReq.MilesInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.MilesInfo{})
	var milesInfos []autocode.MilesInfo
	if info.ClassesId != nil {
		db = db.Where("`classes_id` = ?", info.ClassesId)
	}

	if info.CalcDate != "" {
		db = db.Where("`cal_date` = ?", info.CalcDate)
	}
	err = db.Count(&total).Error
	//err = db.Limit(limit).Offset(offset).Preload("Arrival").Find(&checkinInfos).Error
	err = db.Limit(limit).Offset(offset).Find(&milesInfos).Error
	return err, milesInfos, total
}

func (s *MilesInfoService) ExportExcel(export autoCodeReq.MilesInfoExport) (string, error) {
    f := excelize.NewFile()
    index := f.NewSheet("Sheet1")
    f.SetCellValue("Sheet1", FormatCoord(1, 1), "日期")
    f.SetCellValue("Sheet1", FormatCoord(1, 2), "路线ID")
    f.SetCellValue("Sheet1", FormatCoord(1, 3), "路线名")
    f.SetCellValue("Sheet1", FormatCoord(1, 4), "发车时间")
    f.SetCellValue("Sheet1", FormatCoord(1, 5), "里程数(KM)")
    data, err := s.LoadData(export)
    if err == nil {
        row := 2
        for _, d := range data {
            f.SetCellValue("Sheet1", FormatCoord(row, 1), d.date)
            f.SetCellValue("Sheet1", FormatCoord(row, 2), d.classId)
            f.SetCellValue("Sheet1", FormatCoord(row, 3), d.routeName)
            f.SetCellValue("Sheet1", FormatCoord(row, 4), d.classTime)
            f.SetCellValue("Sheet1", FormatCoord(row, 5), d.distance)
            row++
        }
    }
    f.SetActiveSheet(index)
    path, _ := os.Getwd()
    fileName := fmt.Sprintf("%s/%s_miles_%d.xlsx", path, export.Plate, time.Now().Unix())
    if err := f.SaveAs(fileName); err != nil {
        return "", err
    }
    return fileName, nil
}

func (s *MilesInfoService) LoadData(export autoCodeReq.MilesInfoExport) ([]exportData, error) {
    busID, err := s.LoadBusIDByPlate(export.Plate)
    if err != nil {
        return nil, err
    }
    classes, err := s.LoadClassByID(busID)
    if err != nil {
        return nil, err
    }
    miles_data := make(map[string][]autocode.MilesInfo)
    class_data := make(map[uint][]string)
    for _, class := range classes {
        route, err := s.LoadRouteByID(*class.RouteId)
        if err != nil {
            continue
        }
        milesInfos, err := s.LoadMilesByClass(class.ID, export.BeginTime, export.EndTime)
        if err != nil {
            continue
        }

        for _, miles := range milesInfos {
            stringSlice := strings.Split(miles.CalcDate, "T")
            miles_data[stringSlice[0]] = append(miles_data[stringSlice[0]], miles)
        }
        class_data[class.ID] = append(class_data[class.ID], route.RouteName)
        class_data[class.ID] = append(class_data[class.ID], class.ClassesTime)
    }
    var dates []string
    for date, _ := range miles_data {
        dates = append(dates, date)
    }
    sort.Strings(dates)

    var data []exportData
    for _, date := range dates {
        milesInfo, _ := miles_data[date]
        for _, miles := range milesInfo {
            data = append(data, exportData{
                date: date,
                classId: *miles.ClassesId,
                routeName: class_data[*miles.ClassesId][0],
                classTime: class_data[*miles.ClassesId][1],
                distance: miles.Distance/1000,
            })
        }
    }
    return data, nil
}

func (s *MilesInfoService) LoadMilesByClass(class uint, beginTime, endTime string) ([]autocode.MilesInfo, error) {
    db := global.GVA_DB.Model(&autocode.MilesInfo{})
    var milesInfos []autocode.MilesInfo
    db = db.Where("`classes_id` = ?", class)
    db = db.Where("`cal_date` >= ?", beginTime)
    db = db.Where("`cal_date` <= ?", endTime)
    err := db.Find(&milesInfos).Error
    if err != nil {
        return nil, err
    }
    return milesInfos, nil
}

func (s *MilesInfoService) LoadBusIDByPlate(plate string) (uint, error) {
    var bus autocode.BusInfo
    result := global.GVA_DB.Where("bus_plate = ?", plate).Find(&bus)

    var cnt int64
    _ = result.Count(&cnt)
    if cnt == 0 {
        return 0, errors.New("can not get busid")
    }
    return bus.ID, result.Error
}

func (s *MilesInfoService) LoadClassByID(busID uint) ([]autocode.ClassesInfo, error) {
    var classes []autocode.ClassesInfo
    result := global.GVA_DB.Where("bus_id = ?", busID).Find(&classes)

    var cnt int64
    _ = result.Count(&cnt)
    if cnt == 0 {
        return nil, errors.New("can not get class")
    }
    return classes, result.Error
}

func (s *MilesInfoService) LoadRouteByID(id uint) (autocode.RouteInfo, error) {
    var route autocode.RouteInfo
    result := global.GVA_DB.Where("id = ?", id).Find(&route)

    var cnt int64
    _ = result.Count(&cnt)
    if cnt == 0 {
        return autocode.RouteInfo{}, errors.New("can not get class")
    }
    return route, result.Error
}
