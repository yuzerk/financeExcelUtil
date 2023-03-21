package dto

import "fmt"

type ProjectData struct {
	pmId        string
	projectId   string
	projectName string
	daily       float64 // 日常
	travel      float64 // 差旅
	bonds       float64 // 奖金
	insurance   float64 // 保险
	salary      float64 // 工资
}

type ProjectDataForDepartment struct {
	pmId        string
	projectId   string
	projectName string
	daily       float64 // 日常
	travel      float64 // 差旅
	bonds       float64 // 奖金
	insurance   float64 // 保险
	salary      float64 // 工资
	department  string  // 部门
}

func (p *ProjectDataForDepartment) Print() {
	fmt.Printf("pmId : %s , projectId: %s,projectName: %s, daily : %f,travel : %f,bonds : %f,insurance : %f,salary : %f, deparment: %s\n",
		p.pmId, p.projectId, p.projectName, p.daily, p.travel, p.bonds, p.insurance, p.salary, p.department)
}
