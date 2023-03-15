package db

import (
	"fmt"
	"my/dto"
)

const INSERT_INTO_CARD = "insert into punch_card(`employ_id`, `pm_id`, `project_id`, `project_name`, `punch_month`, `work_time`) value (?,?,?,?,?,?)"

const INSERT_INTO_PAYMENT = "insert into employ_cost(`employ_id`, `employ_name`, `daily_fee`, `travel_fee`, `bonds_fee`, `insurance_fee`, `salary_fee`, `punch_month`, `department_for_finance`) value (?,?,?,?,?,?,?,?,?)"

func InsertToCard(records []*dto.Record, dateString string) {
	tx, err := Db.Begin()
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
	for _, record := range records {
		_, err = tx.Exec(INSERT_INTO_CARD, record.EmployeeId, record.PmId, record.ProjectId, record.ProjectName, dateString, record.WorkSpendTime)
		if err != nil {
			_ = tx.Rollback()
			fmt.Println("exec error", err)
			return
		}
	}
	tx.Commit()
}

func InsertToPayment(payment []*dto.Payment, dateString string) {
	tx, err := Db.Begin()
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
	for _, record := range payment {
		_, err = tx.Exec(INSERT_INTO_PAYMENT,
			record.GetEmployId(), record.GetName(), record.GetDaily(), record.GetTravel(), record.GetBonds(), record.GetInsurance(), record.GetSalary(), dateString, record.GetDepartment())
		if err != nil {
			_ = tx.Rollback()
			fmt.Println("exec error", err)
			return
		}
	}
	tx.Commit()
}
