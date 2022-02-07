package controller

import (
	fcllib "github.com/Maryam-Yumna/FCL-Website-Backend/LogAnalyzer"
	"github.com/Maryam-Yumna/FCL-Website-Backend/src/types"
)

func ExecuteQuery(query types.Query) string {

	// fmt.Println(query.QueryString)
	result := fcllib.NewFCLWrapper().GetLogLDALResult("C:\\Users\\Bisul\\Downloads\\Defs.txt", query.QueryString, `[{"ROOT":{
		"peoples":[
		{	"name":  "Tharindu","city":
		"Gampaha"},{"name":"rajeen","city":"Gampaha"},{"name":"dias","city":"Colombo"}]}}]`)
	return result

}
