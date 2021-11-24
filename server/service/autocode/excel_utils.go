package autocode

import (
    "fmt"
)

func FormatCoord(row, col int) string {
    colName := NumberToColumnName(col)
    coord := fmt.Sprintf("%s%d", colName, row)
    return coord
}

func NumberToColumnName(col int) string {
    var colName string
    for col > 0 {
        col--
        s := fmt.Sprintf("%c", col%26+'A')
        colName = s + colName
        col = (col - col%26) / 26
    }
    return colName
}
