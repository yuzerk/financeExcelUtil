package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"my/dto"
	"strconv"
)

const proId = "PM20230206030000"
const outputSheetName = "sheet1"

const outputpath = "/Users/yuzekai/Desktop/baobei/worktime/res.xlsx"
const outputpath2 = "/Users/yuzekai/Desktop/baobei/worktime/res2.xlsx"

//const outputpath = "res.xlsx"
//const outputpath2 = "res2.xlsx"

func GenExcelOutput(projects []*dto.ProjectData) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet(outputSheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置表头
	f.SetCellValue(outputSheetName, "A1", "项目号")
	f.SetCellValue(outputSheetName, "B1", "项目编号")
	f.SetCellValue(outputSheetName, "C1", "项目名称")
	f.SetCellValue(outputSheetName, "E1", "日常")
	f.SetCellValue(outputSheetName, "F1", "差旅")
	f.SetCellValue(outputSheetName, "G1", "奖金")
	f.SetCellValue(outputSheetName, "H1", "五险一金")
	f.SetCellValue(outputSheetName, "I1", "工资")
	f.SetCellValue(outputSheetName, "J1", "房租")
	f.SetCellValue(outputSheetName, "K1", "折旧")

	// 表数据
	for i := 0; i < len(projects); i++ {
		value := projects[i]
		if value == nil {
			continue
		}
		f.SetCellValue(outputSheetName, "A"+strconv.Itoa(i+2), value.PmId)
		f.SetCellValue(outputSheetName, "B"+strconv.Itoa(i+2), value.ProjectId)
		f.SetCellValue(outputSheetName, "C"+strconv.Itoa(i+2), value.ProjectName)
		f.SetCellValue(outputSheetName, "E"+strconv.Itoa(i+2), value.Daily)
		f.SetCellValue(outputSheetName, "F"+strconv.Itoa(i+2), value.Travel)
		f.SetCellValue(outputSheetName, "G"+strconv.Itoa(i+2), value.Bonds)
		f.SetCellValue(outputSheetName, "H"+strconv.Itoa(i+2), value.Insurance)
		f.SetCellValue(outputSheetName, "I"+strconv.Itoa(i+2), value.Salary)
		f.SetCellValue(outputSheetName, "J"+strconv.Itoa(i+2), value.Rent)
		f.SetCellValue(outputSheetName, "K"+strconv.Itoa(i+2), value.Depreciation)
	}
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(outputpath); err != nil {
		fmt.Println(err)
	}
}

func GenExcelOutput2(projects []*dto.ProjectDataForDepartment) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet(outputSheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 设置表头
	f.SetCellValue(outputSheetName, "A1", "项目号")
	f.SetCellValue(outputSheetName, "B1", "项目编号")
	f.SetCellValue(outputSheetName, "C1", "项目名称")
	f.SetCellValue(outputSheetName, "D1", "部门")
	f.SetCellValue(outputSheetName, "E1", "日常")
	f.SetCellValue(outputSheetName, "F1", "差旅")
	f.SetCellValue(outputSheetName, "G1", "奖金")
	f.SetCellValue(outputSheetName, "H1", "五险一金")
	f.SetCellValue(outputSheetName, "I1", "工资")
	f.SetCellValue(outputSheetName, "J1", "房租")
	f.SetCellValue(outputSheetName, "K1", "折旧")

	// 表数据
	for i := 0; i < len(projects); i++ {
		value := projects[i]
		if value == nil {
			continue
		}
		f.SetCellValue(outputSheetName, "A"+strconv.Itoa(i+2), value.PmId)
		f.SetCellValue(outputSheetName, "B"+strconv.Itoa(i+2), value.ProjectId)
		f.SetCellValue(outputSheetName, "C"+strconv.Itoa(i+2), value.ProjectName)
		f.SetCellValue(outputSheetName, "D"+strconv.Itoa(i+2), value.Department)
		f.SetCellValue(outputSheetName, "E"+strconv.Itoa(i+2), value.Daily)
		f.SetCellValue(outputSheetName, "F"+strconv.Itoa(i+2), value.Travel)
		f.SetCellValue(outputSheetName, "G"+strconv.Itoa(i+2), value.Bonds)
		f.SetCellValue(outputSheetName, "H"+strconv.Itoa(i+2), value.Insurance)
		f.SetCellValue(outputSheetName, "I"+strconv.Itoa(i+2), value.Salary)
		f.SetCellValue(outputSheetName, "J"+strconv.Itoa(i+2), value.Rent)
		f.SetCellValue(outputSheetName, "K"+strconv.Itoa(i+2), value.Depreciation)
	}
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs(outputpath2); err != nil {
		fmt.Println(err)
	}
}

/**
 */
func CalculateProjectPrice(emWorkTimeMap map[string]float64,
	projectIdRecordList map[string][]*dto.Record,
	emPaymentMap map[string]*dto.Payment) []*dto.ProjectData {

	fmt.Println("len(projectIdRecordList) is ", len(projectIdRecordList))
	projectData := make([]*dto.ProjectData, 0)

	for projectId, recordList := range projectIdRecordList {
		pData := new(dto.ProjectData)
		daily := 0.0
		travel := 0.0
		bonds := 0.0
		insurance := 0.0
		salary := 0.0
		rent := 0.0
		depreciation := 0.0
		if projectId == proId {
			for _, record := range recordList {
				record.Print()
			}
		}
		//fmt.Println("calculateProjectPrice: projectId: ", projectId, "   ", len(recordList))
		for _, record := range recordList {

			// 基本数据
			pData.PmId = record.PmId
			pData.ProjectId = record.ProjectId
			pData.ProjectName = record.ProjectName

			emId := record.EmployeeId
			workTime := record.WorkSpendTime

			totalWorkTime := emWorkTimeMap[emId]
			// 每个员工在这个项目中的工时比例
			rate := workTime / totalWorkTime
			//if projectId == proId {
			//	fmt.Println(projectId, "每个员工在这个项目中的工时比例   ", emId, "  ", totalWorkTime, "   ", rate)
			//}
			//fmt.Println("recordList: ", emId)
			if emPaymentMap[emId] == nil {
				continue
			}
			daily += rate * emPaymentMap[emId].GetDaily()
			travel += rate * emPaymentMap[emId].GetTravel()
			bonds += rate * emPaymentMap[emId].GetBonds()
			insurance += rate * emPaymentMap[emId].GetInsurance()
			salary += rate * emPaymentMap[emId].GetSalary()
			rent += rate * emPaymentMap[emId].GetRent()
			depreciation += rate * emPaymentMap[emId].GetDepreciation()
		}
		pData.Daily = daily
		pData.Travel = travel
		pData.Bonds = bonds
		pData.Insurance = insurance
		pData.Salary = salary
		pData.Rent = rent
		pData.Depreciation = depreciation
		fmt.Printf("%s 项目中 日常：%f，差旅： %f, 奖金：%f, 五险一金: %f, 工资: %f \n", projectId, daily, travel, bonds, insurance, salary)
		projectData = append(projectData, pData)
	}
	return projectData
}

/*
*
根据部门分摊
*/
func CalculateProjectPriceByDepartment(emWorkTimeMap map[string]float64,
	projectIdRecordList map[string][]*dto.Record,
	emPaymentMap map[string]*dto.Payment) []*dto.ProjectDataForDepartment {

	fmt.Println("len(projectIdRecordList) is ", len(projectIdRecordList))
	projectData := make([]*dto.ProjectDataForDepartment, 0)

	for projectId, recordList := range projectIdRecordList {
		if projectId == proId {
			//fmt.Println("pppppppppppppppppppppp")
			//for _, record := range recordList {
			//	record.Print()
			//}
		}
		//fmt.Println("calculateProjectPrice: projectId: ", projectId, "   ", len(recordList))
		projectForDepartment := make(map[string]*dto.ProjectDataForDepartment)
		for _, record := range recordList {

			pData := new(dto.ProjectDataForDepartment)

			// 基本数据
			pData.PmId = record.PmId
			pData.ProjectId = record.ProjectId
			pData.ProjectName = record.ProjectName

			emId := record.EmployeeId
			workTime := record.WorkSpendTime

			totalWorkTime := emWorkTimeMap[emId]
			// 每个员工在这个项目中的工时比例
			rate := workTime / totalWorkTime
			//if projectId == proId {
			//	fmt.Println(projectId, "每个员工在这个项目中的工时比例   ", emId, "  ", totalWorkTime, "   ", rate)
			//}
			//fmt.Println("recordList: ", emId)
			if emPaymentMap[emId] == nil {
				continue
			}
			pData.Department = emPaymentMap[emId].GetDepartment()
			pData.Daily = rate * emPaymentMap[emId].GetDaily()
			pData.Travel = rate * emPaymentMap[emId].GetTravel()
			pData.Bonds = rate * emPaymentMap[emId].GetBonds()
			pData.Insurance = rate * emPaymentMap[emId].GetInsurance()
			pData.Salary = rate * emPaymentMap[emId].GetSalary()
			pData.Rent = rate * emPaymentMap[emId].GetRent()
			pData.Depreciation = rate * emPaymentMap[emId].GetDepreciation()

			pDataStore := projectForDepartment[pData.Department]
			if pDataStore != nil {
				pDataStore.Daily += pData.Daily
				pDataStore.Travel += pData.Travel
				pDataStore.Bonds += pData.Bonds
				pDataStore.Insurance += pData.Insurance
				pDataStore.Salary += pData.Salary
				pDataStore.Rent += pData.Rent
				pDataStore.Depreciation += pData.Depreciation
			} else {
				projectForDepartment[pData.Department] = pData
			}
		}

		for _, projectDetail := range projectForDepartment {
			fmt.Printf("%s 项目中 日常：%f，差旅： %f, 奖金：%f, 五险一金: %f, 工资: %f , 部门: %s\n", projectDetail.ProjectId, projectDetail.Daily, projectDetail.Travel, projectDetail.Bonds, projectDetail.Insurance, projectDetail.Salary, projectDetail.Department)
			projectData = append(projectData, projectDetail)
		}
	}
	return projectData
}

/*
*
for 一维打卡表
*/
func getEmployeeRecord(filepath string) (map[string]float64, map[string][]*dto.Record) {
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
	projectIdRecordList := make(map[string][]*dto.Record)
	for index, row := range rows {
		if index == 0 {
			continue
		}
		record := new(dto.Record)
		for index, col := range row {
			switch index {
			case 0:
				record.PmId = col
			case 1:
				record.ProjectId = col
			case 2:
				record.ProjectName = col
			case 5:
				record.Employee = col
			case 6:
				record.EmployeeId = col
			case 8:
				record.WorkSpendTime, _ = strconv.ParseFloat(col, 64)
			}
			emWorkTimeMap[record.EmployeeId] = emWorkTimeMap[record.EmployeeId] + record.WorkSpendTime
		}
		childList := projectIdRecordList[record.ProjectId]
		childList = append(childList, record)
		projectIdRecordList[record.ProjectId] = childList
	}
	fmt.Println("getEmployeeRecord", len(emWorkTimeMap), len(projectIdRecordList))
	return emWorkTimeMap, projectIdRecordList
}

/*
*
for 二维打卡表
*/
func GetEmployeeRecordForTwoDimension(filepath string) (map[string]float64, map[string][]*dto.Record) {
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
	projectIdRecordList := make(map[string][]*dto.Record)
	for index, row := range rows {
		if index <= 2 {
			continue
		}
		for colIndex, col := range row {
			if len(col) == 0 {
				continue
			}
			if colIndex <= 14 {
				continue
			}
			record := new(dto.Record)
			record.PmId = row[0]
			record.ProjectId = row[1]
			record.ProjectName = row[2]
			colNameString, err := ConvertNumToChar(colIndex + 1)
			if err != nil {
				fmt.Println("ConvertNumToChar error, ", err)
			}
			data, err := f.GetCellValue(sheetName, colNameString+"3")
			if err != nil {
				fmt.Println("GetCellValue error, ", err, "|||||||", colNameString, 1)
			}
			record.EmployeeId = data
			record.WorkSpendTime, err = strconv.ParseFloat(col, 64)
			if record.EmployeeId == "0005720" {
				fmt.Println(colNameString)
				record.PrintWithPrefix("xxxxxxxxxxxxx" + strconv.Itoa(index) + "_" + strconv.Itoa(colIndex))
			}
			if err != nil {
				fmt.Println("strconv.ParseFloat error, ", err, "|||||||", col)
			}
			emWorkTimeMap[record.EmployeeId] = emWorkTimeMap[record.EmployeeId] + record.WorkSpendTime
			childList := projectIdRecordList[record.ProjectId]
			childList = append(childList, record)
			projectIdRecordList[record.ProjectId] = childList
		}
	}
	fmt.Println("getEmployeeRecordForTwoDimension", len(emWorkTimeMap), len(projectIdRecordList))
	return emWorkTimeMap, projectIdRecordList
}

/*
*
工资表处理
*/
func GetEmployeePayment(filepath string) map[string]*dto.Payment {
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
	eIdPaymentMap := make(map[string]*dto.Payment)
	for index, row := range rows {
		if index <= 1 {
			continue
		}
		payment := new(dto.Payment)
		payment.SetEmployId(row[1])
		//payment.SetDaily(row[48-1])
		//payment.SetTravel(row[49-1])
		//payment.SetBonds(row[53-1])
		//payment.SetInsurance(row[54-1])
		//payment.SetSalary(row[55-1])
		//payment.SetDepartment(row[17])

		payment.SetDaily(row[22-1])
		payment.SetTravel(row[23-1])
		payment.SetBonds(row[21-1])
		payment.SetInsurance(row[20-1])
		payment.SetSalary(row[19-1])
		payment.SetDepartment(row[17])
		payment.SetRent(row[24-1])
		payment.SetDepreciation(row[25-1])
		eIdPaymentMap[payment.GetEmployId()] = payment
	}
	//fmt.Println("getEmployeePayment", len(eIdPaymentMap))
	return eIdPaymentMap
}

func GetEmployeeRecordForDb(filepath string) []*dto.Record {
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
	records := make([]*dto.Record, 0)
	for index, row := range rows {
		if index <= 2 {
			continue
		}
		for colIndex, col := range row {
			if len(col) == 0 {
				continue
			}
			if colIndex <= 14 {
				continue
			}
			record := new(dto.Record)
			record.PmId = row[0]
			record.ProjectId = row[1]
			record.ProjectName = row[2]
			colNameString, err := ConvertNumToChar(colIndex + 1)
			if err != nil {
				fmt.Println("ConvertNumToChar error, ", err)
			}
			data, err := f.GetCellValue(sheetName, colNameString+"3")
			if err != nil {
				fmt.Println("GetCellValue error, ", err, "|||||||", colNameString, 1)
			}
			record.EmployeeId = data
			record.WorkSpendTime, err = strconv.ParseFloat(col, 64)
			if record.EmployeeId == "0005720" {
				fmt.Println(colNameString)
				record.PrintWithPrefix("xxxxxxxxxxxxx" + strconv.Itoa(index) + "_" + strconv.Itoa(colIndex))
			}
			if err != nil {
				fmt.Println("strconv.ParseFloat error, ", err, "|||||||", col)
			}
			records = append(records, record)
		}
	}
	return records
}

func GetEmployeePaymentForDB(filepath string) []*dto.Payment {
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
	paymentList := make([]*dto.Payment, 0)
	for index, row := range rows {
		if index <= 1 {
			continue
		}
		payment := new(dto.Payment)
		payment.SetEmployId(row[1])
		payment.SetDaily(row[48-1])
		payment.SetTravel(row[49-1])
		payment.SetBonds(row[53-1])
		payment.SetInsurance(row[54-1])
		payment.SetSalary(row[55-1])
		payment.SetDepartment(row[17])
		payment.SetName(row[2])
		paymentList = append(paymentList, payment)
	}
	//fmt.Println("getEmployeePayment", len(eIdPaymentMap))
	return paymentList
}
