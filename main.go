package main

import (
	"github.com/gin-gonic/gin"
	"my/db"
	"my/dto"
	"my/excel"
	"my/syn"
	"net/http"
)

const cardRecordFile = "/Users/yuzekai/Desktop/baobei/worktime/card.xlsx"
const paymentFile = "/Users/yuzekai/Desktop/baobei/worktime/payment.xlsx"

func main() {
	//doExcel()
	//doCardSave()
	//doCostSave()
	//doSelectCard("PM20210701102533")
	//router := gin.Default()
	//router.GET("/getCardsByProjectId", func(context *gin.Context) {
	//	projectId := context.Query("projectId")
	//	records := doSelectCard(projectId)
	//	context.JSON(http.StatusOK, records)
	//})
	//router.Use(CrosHandler())
	//router.Run(":8088")
	syn.Exe()
}

// 跨域访问：cross  origin resource share
func CrosHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,token,openid,opentoken")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")
		context.Header("Access-Control-Allow-Credentials", "false")
		context.Set("content-type", "application/json")

		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}

		//处理请求
		context.Next()
	}
}

func doSelectCard(projectId string) []*dto.Record {
	records := db.SelectCardsByProjectId(projectId, 1, 10)
	for _, record := range records {
		record.Print()
	}
	return records
}

func doCardSave() {
	records := excel.GetEmployeeRecordForDb(cardRecordFile)
	db.InsertToCard(records, "2023-02-01")
}

func doCostSave() {
	payments := excel.GetEmployeePaymentForDB(paymentFile)
	db.InsertToPayment(payments, "2023-02-01")
}

func doExcel() {

	// 一维表
	//emWorkTimeMap, projectIdRecordList := getEmployeeRecord(cardRecordFile)

	// 二维表
	emWorkTimeMap, projectIdRecordList := excel.GetEmployeeRecordForTwoDimension(cardRecordFile)

	emPaymentMap := excel.GetEmployeePayment(paymentFile)

	projectData := excel.CalculateProjectPrice(emWorkTimeMap, projectIdRecordList, emPaymentMap)
	excel.GenExcelOutput(projectData)

	projectDepartmentData := excel.CalculateProjectPriceByDepartment(emWorkTimeMap, projectIdRecordList, emPaymentMap)
	////for _, p := range projectDepartmentData {
	////	if p.pmId == "PM5930" {
	////		p.Print()
	////	}
	////}
	excel.GenExcelOutput2(projectDepartmentData)
}
