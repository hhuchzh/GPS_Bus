package autocode

import (
    "os"
    "time"
    "fmt"
    "errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"

    "github.com/xuri/excelize/v2"
)

type MilesInfoService struct {
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
    busID, err := s.LoadBusIDByPlate(export.Plate)
    if err != nil {
        return "", err
    }
    classes, err := s.LoadClassByID(busID)
    if err != nil {
        return "", err
    }

    f := excelize.NewFile()
    index := f.NewSheet("Sheet1")
    f.SetCellValue("Sheet1", FormatCoord(1, 1), "路线ID")
    f.SetCellValue("Sheet1", FormatCoord(1, 2), "路线名")
    f.SetCellValue("Sheet1", FormatCoord(1, 3), "日期")
    f.SetCellValue("Sheet1", FormatCoord(1, 4), "里程(米)")
    row := 2
    for _, class := range classes {
        milesInfos, err := s.LoadMilesByClass(class.ID, export.BeginTime, export.EndTime)
        if err != nil {
            continue
        }
        for _, miles := range milesInfos {
            f.SetCellValue("Sheet1", FormatCoord(row, 1), class.ID)
            row++
            f.SetCellValue("Sheet1", FormatCoord(row, 2), class.Remark)
            row++
            f.SetCellValue("Sheet1", FormatCoord(row, 3), miles.CalcDate)
            row++
            f.SetCellValue("Sheet1", FormatCoord(row, 4), miles.Distance)
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
