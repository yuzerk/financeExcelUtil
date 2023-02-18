package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func main() {
	cardRecordFile := "/Users/yuzekai/Desktop/card.xlsx"
	paymentFile := "/Users/yuzekai/Desktop/payment.xlsx"

	emWorkTimeMap, projectIdRecordList := getEmployeeRecord(cardRecordFile)

	emPaymentMap := getEmployeePayment(paymentFile)

	projectData := calculateProjectPrice(emWorkTimeMap, projectIdRecordList, emPaymentMap)

	fmt.Println(len(projectData))
}

func calculateProjectPrice(emWorkTimeMap map[string]float64,
	projectIdRecordList map[string][]*Record,
	emPaymentMap map[string]*Payment) []*ProjectData {

	projectData := make([]*ProjectData, len(projectIdRecordList))

	for projectId, recordList := range projectIdRecordList {
		pData := new(ProjectData)
		daily := 0.0
		travel := 0.0
		bonds := 0.0
		insurance := 0.0
		salary := 0.0
		//fmt.Println("calculateProjectPrice: projectId: ", projectId, "   ", len(recordList))
		for _, record := range recordList {
			// 基本数据
			pData.pmId = record.pmId
			pData.projectId = record.projectId
			pData.projectName = record.projectName

			emId := record.employeeId
			workTime := record.workSpendTime

			totalWorkTime := emWorkTimeMap[emId]
			// 每个员工在这个项目中的工时比例
			rate := workTime / totalWorkTime
			//fmt.Println("recordList: ", emId)
			if emPaymentMap[emId] == nil {
				continue
			}
			daily += rate * emPaymentMap[emId].GetDaily()
			travel += rate * emPaymentMap[emId].GetTravel()
			bonds += rate * emPaymentMap[emId].GetBonds()
			insurance += rate * emPaymentMap[emId].GetInsurance()
			salary += rate * emPaymentMap[emId].GetSalary()
		}
		pData.daily = daily
		pData.travel = travel
		pData.bonds = bonds
		pData.insurance = insurance
		pData.salary = salary
		fmt.Printf("%s 项目中 日常：%f，差旅： %f, 奖金：%f, 五险一金: %f, 工资: %f \n", projectId, daily, travel, bonds, insurance, salary)
		projectData = append(projectData, pData)
	}
	return projectData
}

func getEmployeeRecord(filepath string) (map[string]float64, map[string][]*Record) {
	f, err := excelize.OpenFile(filepath)
	if err != err {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}

	// 工号map 总工时
	emWorkTimeMap := make(map[string]float64)
	// 项目号map recordList
	projectIdRecordList := make(map[string][]*Record)
	for index, row := range rows {
		if index == 0 {
			continue
		}
		record := new(Record)
		for index, col := range row {
			switch index {
			case 0:
				record.pmId = col
			case 1:
				record.projectId = col
			case 2:
				record.projectName = col
			case 5:
				record.employee = col
			case 6:
				record.employeeId = col
			case 8:
				record.workSpendTime, _ = strconv.ParseFloat(col, 64)
			}
			emWorkTimeMap[record.employeeId] = emWorkTimeMap[record.employeeId] + record.workSpendTime
		}
		childList := projectIdRecordList[record.projectId]
		childList = append(childList, record)
		projectIdRecordList[record.projectId] = childList
	}
	fmt.Println("getEmployeeRecord", len(emWorkTimeMap), len(projectIdRecordList))
	return emWorkTimeMap, projectIdRecordList
}

func getEmployeePayment(filepath string) map[string]*Payment {
	f, err := excelize.OpenFile(filepath)
	if err != err {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}
	eIdPaymentMap := make(map[string]*Payment)
	for index, row := range rows {
		if index <= 1 {
			continue
		}
		payment := new(Payment)
		payment.employeeId = row[1]
		payment.daily = row[48]
		payment.travel = row[49]
		payment.bonds = row[53]
		payment.insurance = row[54]
		payment.salary = row[55]
		eIdPaymentMap[payment.employeeId] = payment
	}
	fmt.Println("getEmployeePayment", len(eIdPaymentMap))
	return eIdPaymentMap
}
