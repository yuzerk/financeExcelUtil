package main

import (
	"my/db"
)

const cardRecordFile = "/Users/yuzekai/Desktop/baobei/worktime/card.xlsx"
const paymentFile = "/Users/yuzekai/Desktop/baobei/worktime/payment.xlsx"

func main() {
	//doExcel()
	//doCardSave()
	//doCostSave()
	doSelectCard("PM20210701102533")

}

func doSelectCard(projectId string) {
	records := db.SelectCardsByProjectId(projectId, 1, 10)
	for _, record := range records {
		record.Print()
	}
}

func doCardSave() {
	records := GetEmployeeRecordForDb(cardRecordFile)
	db.InsertToCard(records, "2023-02-01")
}

func doCostSave() {
	payments := GetEmployeePaymentForDB(paymentFile)
	db.InsertToPayment(payments, "2023-02-01")
}

func doExcel() {

	// 一维表
	//emWorkTimeMap, projectIdRecordList := getEmployeeRecord(cardRecordFile)

	// 二维表
	emWorkTimeMap, projectIdRecordList := getEmployeeRecordForTwoDimension(cardRecordFile)

	emPaymentMap := getEmployeePayment(paymentFile)

	projectData := calculateProjectPrice(emWorkTimeMap, projectIdRecordList, emPaymentMap)
	genExcelOutput(projectData)

	projectDepartmentData := calculateProjectPriceByDepartment(emWorkTimeMap, projectIdRecordList, emPaymentMap)
	////for _, p := range projectDepartmentData {
	////	if p.pmId == "PM5930" {
	////		p.Print()
	////	}
	////}
	genExcelOutput2(projectDepartmentData)
}
