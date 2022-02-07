package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Maryam-Yumna/FCL-Website-Backend/src/routes"
)

func main() {
	fmt.Println("Hello")
	router := routes.Router()
	router.Path("/")
	http.Handle("/", router)
	log.Println("Server Started localhost :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

	// handleRequests()

	// result := fcllib.NewFCLWrapper().GetLogLDALResult("C:\\Users\\User\\Desktop\\QueryTest\\Defs.txt", `$RESULT.SetCustomString(object)
	// $RESULT.SetRValue(4)

	// //define X, Y axis
	// $RESULT.AddNode:=AXIS
	// $AXIS.SetCustomString(object)
	// $AXIS.SetRValue(4)
	// $AXIS.SetLValue(axis)
	// $AXIS.AddNode:=NAMETAG
	// $NAMETAG.SetCustomString(normal)
	// $NAMETAG.SetRValue(4)
	// $NAMETAG.SetLValue(Xaxis)
	// $NAMETAG.SetValue(District)
	// $AXIS.AddNode:=NAMETAG
	// $NAMETAG.SetCustomString(normal)
	// $NAMETAG.SetRValue(4)
	// $NAMETAG.SetLValue(Yaxis)
	// $NAMETAG.SetValue(People count)

	// $RESULT.AddNode:=DATA
	// 	$DATA.SetCustomString(array)
	// 	$DATA.SetRValue(4)
	// 	$DATA.SetLValue(data)
	// 	$DATA.AddNode:=GAMPAHA
	// //Gampaha Data
	// $GAMPAHA.SetCustomString(object)
	// $GAMPAHA.SetRValue(4)
	// $GAMPAHA.AddNode:=NAMETAG
	// $NAMETAG.SetCustomString(normal)
	// $NAMETAG.SetLValue(name)
	// $NAMETAG.SetValue(Gampaha)
	// $GAMPAHA.AddNode:=PEOPLECOUNT
	// $PEOPLECOUNT.SetCustomString(normal)
	// $PEOPLECOUNT.SetLValue(value)
	// $X.FilterSubtree($Item.GetValue.IsStringEqualTo(Gampaha)):=GampahaList
	// $GampahaList.GetItemCount.ToString:=STR
	// $PEOPLECOUNT.SetValue($STR)
	// //Colombo Data
	// $DATA.AddNode:=COLOMBO
	// $COLOMBO.SetCustomString(object)
	// $COLOMBO.SetRValue(4)
	// $COLOMBO.AddNode:=NAMETAG
	// $NAMETAG.SetCustomString(normal)
	// $NAMETAG.SetLValue(name)
	// $NAMETAG.SetValue(Colombo)
	// $COLOMBO.AddNode:=PEOPLECOUNT
	// $PEOPLECOUNT.SetCustomString(normal)
	// $PEOPLECOUNT.SetLValue(value)
	// $X.FilterSubtree($Item.GetValue.IsStringEqualTo(Colombo)):=ColomboList
	// $ColomboList.GetItemCount.ToString:=STR
	// $PEOPLECOUNT.SetValue($STR)`, `[{"ROOT":{
	// 	"peoples":[
	// 	{	"name":  "Tharindu","city":
	// 	"Gampaha"},{"name":"rajeen","city":"Gampaha"},{"name":"dias","city":"Colombo"}]}}]`)

	// fmt.Println(result)
}
