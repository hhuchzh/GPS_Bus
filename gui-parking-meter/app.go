package main

import (
    "fmt"
    "os"
    "strconv"
    //"strings"

    "github.com/xuri/excelize/v2"
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
)

const (
    startName = "开始时间"
    endName = "结束时间"
    priceName = "应收金额"
	vehicleName = "车型"
	defaultVehicleType = "小型车"
	sheetName = "Sheet1"
)

func NumberToColumnName(col int) string {
    var colName string
    for col > 0 {
        col--;
        s := fmt.Sprintf("%c", col%26 + 'A')
        colName = s + colName
        col = (col - col % 26) / 26
    }
    return colName
}

type MeterMainWindow struct {
	*walk.MainWindow
	model *RegionModel
	cb *walk.ComboBox
	edit *walk.TextEdit
	dlg *walk.FileDialog
	progressBar *walk.ProgressBar
}

type App struct {
	mw *MeterMainWindow
}

func NewApp() *App {
	return &App{
		mw: &MeterMainWindow{
			model: NewRegionModel(),
		},
	}
}

func (a *App) Run() {
	err := MainWindow{
		AssignTo: &a.mw.MainWindow,
		Title: "公共停车场收费违规检测",
		MinSize: Size{Width: 600, Height: 400},
		MaxSize: Size{Width: 600, Height: 800},
		Size: Size{600, 400},
		Layout: VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					Label{
						Text: "选择区域",
					},
					ComboBox{
						AssignTo: &a.mw.cb,
						Model:    a.mw.model,
						OnCurrentIndexChanged: a.mw.regionIndexChanged,
					},
					PushButton{
						Text: "选择文件",
						OnClicked: a.mw.selectFile,
					},
					PushButton{
						Text: "开始检测",
						OnClicked: a.mw.calc,
					},
				},
			},
			ScrollView{
				Layout: VBox{},
				Children: []Widget{
					TextEdit{
						AssignTo: &a.mw.edit,
						ReadOnly: true,
						HScroll: true,
						VScroll: true,
						Font:     Font{PointSize: 10},
					},
				},
			},
			ProgressBar{
                AssignTo: &a.mw.progressBar,
				MinValue: 0,
				MaxValue: 100,
            },
		},
	}.Create()
	if err != nil {
		fmt.Println("--> MainWindow Create error ", err)
		os.Exit(1)
	}
	a.mw.Run()
}

type RegionModel struct {
	walk.ListModelBase
	lastIndex int
	items []string
}

func NewRegionModel() *RegionModel {
	return &RegionModel{
		lastIndex: -1,
		items: []string{"核心区域公共停车场",
				"一级区域公共停车场",
				"二级区域公共停车场",
				"三级区域公共停车场",
				"火车站公共停车场"},
	}
}

func (m *RegionModel) ItemCount() int {
	return len(m.items)
}

func (m *RegionModel) Value(index int) interface{} {
	return m.items[index]
}

func (mw *MeterMainWindow) regionIndexChanged() {
	i := mw.cb.CurrentIndex()
	if mw.model.lastIndex == i {
		return
	}
	mw.edit.SetText("")
	mw.model.lastIndex = i
	s := fmt.Sprintf("区域: %s", mw.model.items[i])
	if mw.model.items[i] != "二级区域公共停车场" && mw.model.items[i] != "三级区域公共停车场" {
		s = s + "，该区域暂不支持\r\n"
	} else {
		s = s + "\r\n"
	}
	mw.edit.AppendText(s)
}

func (mw *MeterMainWindow) selectFile() {
	if mw.model.lastIndex == -1 {
		s := fmt.Sprintf("未选择区域，请先选择区域\r\n")
		mw.edit.AppendText(s)
		return
	}
	i := mw.model.lastIndex
	if mw.model.items[i] != "二级区域公共停车场" && mw.model.items[i] != "三级区域公共停车场" {
		s := fmt.Sprintf("所选区域暂不支持\r\n")
		mw.edit.AppendText(s)
		return
	}
	mw.dlg = new(walk.FileDialog)
	mw.dlg.Title = "选择文件"
	mw.dlg.Filter = "xlsx文件 (*.xlsx)|*.xlsx"

	if ok, err := mw.dlg.ShowOpen(mw); err != nil {
		mw.edit.AppendText("Error: File Open\r\n")
		return
	} else if !ok {
		//mw.edit.AppendText("Cancel\r\n")
		return
	}
	s := fmt.Sprintf("文件: %s\r\n", mw.dlg.FilePath)
	mw.edit.AppendText(s)
}

func (mw *MeterMainWindow) calc() {
	if mw.dlg == nil || mw.dlg.FilePath == "" {
		mw.edit.AppendText("未选择文件，请先选择文件\r\n")
		return
	}
	go process(mw)
}

func process(mw *MeterMainWindow) {
	//mw.progressBar.SetMarqueeMode(true)
	_, err := PriceConfig.GetPrice(mw.model.items[mw.model.lastIndex], defaultVehicleType)
	if err != nil {
		s := fmt.Sprintf("检测错误： %v\r\n", err)
		mw.edit.AppendText(s)
		return
	}
	f, err := excelize.OpenFile(mw.dlg.FilePath)
    if err != nil {
        mw.edit.AppendText("打开文件失败\r\n")
        return
    }

	mw.edit.AppendText("开始读入文件...\r\n")
    var startIndex, endIndex, priceIndex, vehicleIndex int
    rows, err := f.GetRows(sheetName)
    if err != nil {
        mw.edit.AppendText("读取文件失败\r\n")
        return
    }
    startIndex = -1
    endIndex = -1
    priceIndex = -1
	vehicleIndex = -1
    for i, col := range rows[0] {
        if col == startName {
            startIndex = i
        } else if col == endName {
            endIndex = i
        } else if col == priceName {
            priceIndex = i
        } else if col == vehicleName {
			vehicleIndex = i
		}
    }
    //fmt.Printf("startIndex %d endIndex %d priceIndex %d\n", startIndex, endIndex, priceIndex)
    if startIndex == -1 || endIndex == -1 || priceIndex == -1 {
        s := fmt.Sprintf("文件中未含有\"%s\"或\"%s\"或\"%s\"，请检查\r\n", startName, endName, priceName)
		mw.edit.AppendText(s)
        return
    }
	mw.edit.AppendText("读入完成，开始检测...\r\n")
	mw.progressBar.SetValue(0)

    calc := &Calculator{}
    numCol := len(rows[0])
	style, err := f.NewStyle(`{"font":{"size":10, "color":"#FF0000"}, "alignment":{"horizontal":"center", "vertical": "center"}}`)
	style2, err := f.NewStyle(`{"font":{"size":10, "color":"#000000"}, "alignment":{"horizontal":"center", "vertical": "center"}}`)
    colName := NumberToColumnName(numCol+1)
    coord := fmt.Sprintf("%s%d", colName, 1)
    f.SetCellValue(sheetName, coord, "理论应收费用")
	_ = f.SetCellStyle(sheetName, coord, coord, style2)
	colName2 := NumberToColumnName(numCol+2)
    coord = fmt.Sprintf("%s%d", colName2, 1)
	f.SetCellValue(sheetName, coord, "多收费用")
	_ = f.SetCellStyle(sheetName, coord, coord, style2)

    rows = rows[1:]
	totalCount := len(rows)
    var count int
	step := float64(1)/float64(totalCount)
	progress := step
	if err != nil {
		fmt.Println(err)
	}
	
	var sum float64
    for i, row := range rows {
        startTime := row[startIndex]
        endTime := row[endIndex]
        price := row[priceIndex]
		vehicleType := defaultVehicleType
		if vehicleIndex != -1 {
			vehicleType = row[vehicleIndex]
		}
		ps, err := PriceConfig.GetPrice(mw.model.items[mw.model.lastIndex], vehicleType)
		if err != nil {
			s := fmt.Sprintf("检测错误： %v\r\n", err)
			mw.edit.AppendText(s)
			continue
		}
        totalPrice, err := calc.Calc(startTime, endTime, ps)
        if err != nil {
            fmt.Println("--> calc.Calc error ", err)
			s := fmt.Sprintf("第%d条数据检测错误： %v\r\n", i+2, err)
			mw.edit.AppendText(s)
            continue
        }
        originPrice, _ := strconv.ParseFloat(price, 64)

		coord := fmt.Sprintf("%s%d", colName, i+2)
		err = f.SetCellValue(sheetName, coord, fmt.Sprintf("%.1f", totalPrice))
		err = f.SetCellStyle(sheetName, coord, coord, style2)
		if err != nil {
            fmt.Println(err)
        }
        coord = fmt.Sprintf("%s%d", colName2, i+2)
        if totalPrice < originPrice {
            v := fmt.Sprintf("+%.1f", originPrice - totalPrice)
			sum = sum + (originPrice - totalPrice)
            err = f.SetCellValue(sheetName, coord, v)
            if err != nil {
                fmt.Println(err)
            }
			err = f.SetCellStyle(sheetName, coord, coord, style)
            count++
        }
		mw.progressBar.SetValue(int(progress*100))
		progress += step
    }
	
		index:=fmt.Sprintf("%s%d",colName2,totalCount+2)
		val:=fmt.Sprintf("total:%.1f",sum)
		err = f.SetCellValue(sheetName, index, val)
        if err != nil {
            fmt.Println(err)
        }
		err = f.SetCellStyle(sheetName, index, index, style)
		 if err != nil {
            fmt.Println(err)
        }
	
    err = f.Save()
    if err != nil {
        s := fmt.Sprintf("写入检测结果，保存文件失败\r\n")
		mw.edit.AppendText(s)
    } else {
		mw.progressBar.SetValue(100)
		s := fmt.Sprintf("检测完成，共%d条数据，其中%d条违规，检测结果已追加到原文件\r\n", totalCount, count)
		mw.edit.AppendText(s)
	}
	//mw.progressBar.SetMarqueeMode(false)
}
