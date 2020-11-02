package main

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

func DoExcel() {
	file, err := os.Open("./CustomRiverCity.json")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	content_b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	//content_s := string(content_b)
	var result map[string][]map[string]string
	err = json.Unmarshal(content_b, &result)
	if err != nil {
		panic(err)
	}

	xlsx := xlsx.NewFile()
	sheetRiver, err := xlsx.AddSheet("riverCityList")
	if err != nil {
		panic(err)
	}
	row := sheetRiver.AddRow()
	name := row.AddCell()
	name.Value = "name"
	center_coordinate := row.AddCell()
	center_coordinate.Value = "center_coordinate"
	all_coordinate := row.AddCell()
	all_coordinate.Value = "all_coordinate"
	land_id := row.AddCell()
	land_id.Value = "land_id"

	// 渡口
	riverCityList := result["riverCityList"]
	for _, city := range riverCityList {
		//info := city["name"] + "," + city["city1"] + "," + city["tileList"] + "," + "4" + "\n"
		//wFile.WriteString(info)
		row := sheetRiver.AddRow()
		name := row.AddCell()
		name.Value = city["name"]
		center_coordinate := row.AddCell()
		center_coordinate.Value = city["city1"]
		all_coordinate := row.AddCell()
		all_coordinate.Value = strings.Replace(city["tileList"], "|", ";", -1)
		land_id := row.AddCell()
		land_id.Value = "4"
	}

	sheetCity, err := xlsx.AddSheet("cityList")
	if err != nil {
		panic(err)
	}
	row_c := sheetCity.AddRow()
	name_c := row_c.AddCell()
	name_c.Value = "name"
	center_coordinate_c := row_c.AddCell()
	center_coordinate_c.Value = "center_coordinate"
	all_coordinate_c := row_c.AddCell()
	all_coordinate_c.Value = "all_coordinate"
	land_id_c := row_c.AddCell()
	land_id_c.Value = "land_id"
	// 城池
	cityList := result["cityList"]
	for _, city := range cityList {
		row := sheetCity.AddRow()
		name := row.AddCell()
		name.Value = city["name"]
		center_coordinate := row.AddCell()
		center_coordinate.Value = city["key"]
		all_coordinate := row.AddCell()
		all_coordinate.Value = city["key"]
		land_id := row.AddCell()
		land_id.Value = "3"
	}

	xlsx.Save("file.xlsx")

}
