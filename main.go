package main

import (
	"fmt"
	"my/db"
)

func main() {
	//doExcel()
	tx, err := db.Db.Begin()
	if err != nil {
		if tx != nil {
			_ = tx.Rollback()
		}
		fmt.Printf("begin trans action failed, err:%v\n", err)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			// 打印异常，关闭资源，退出此函数
			fmt.Println(err)
			_ = tx.Rollback()
		}
	}()
	_, err1 := tx.Exec(db.INSERT_INTO_CARD, "111", "p123", "2023-02-01", 10.5)
	if err1 != nil {
		_ = tx.Rollback()
		fmt.Println("exec error", err1)
		return
	}
	tx.Commit()
}

func doExcel() {
	cardRecordFile := "/Users/yuzekai/Desktop/baobei/worktime/card.xlsx"
	paymentFile := "/Users/yuzekai/Desktop/baobei/worktime/payment.xlsx"

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
