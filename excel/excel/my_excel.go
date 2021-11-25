package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"net/http"
	"os"
)


func CreateExcel(w http.ResponseWriter, r *http.Request)  {
	file := excelize.NewFile()

	//每个单元格添加
	file.SetCellValue("Sheet1", "A1", "工号")
	file.SetCellValue("Sheet1", "B1", "姓名")
	file.SetCellValue("Sheet1", "C1", "性别")
	file.SetCellValue("Sheet1", "D1", "部门")
	file.SetCellValue("Sheet1", "E1", "职位")

	//1行添加
	row := []interface{}{123456, "拜登", "男", "有关部门", "美利坚同志"}
	file.SetSheetRow("Sheet1", "A2", &row)

	//暂存在tmp目录下
	if err := file.SaveAs("./tmp/demo.xlsx"); err != nil {
		fmt.Printf("Create excel error: %v\n", err)
		w.Write([]byte("Create excel failed"))
		return
	}
	w.Write(([]byte)("Create excel success"))
	return
}

func DownloadExcel(w http.ResponseWriter, r *http.Request)  {
	file, err := os.Open("./tmp/demo.xlsx")
	if err != nil {
		fmt.Printf("Open excel file error: %v\n", err)
		w.Write([]byte("Open excel failed"))
		return
	}
	defer file.Close()

	//将文件读取出来
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ReadAll excel file error: %v\n", err)
		w.Write([]byte("ReadAll excel failed"))
		return
	}
	w.Header().Set("Content-Disposition", `attachment; filename="test.xlsx"`)
	//w.Header().Set("Content-Disposition", "inline")
	w.Write(data)
	return
}

