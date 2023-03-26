package dto

import "fmt"

type ProjectData struct {
	PmId         string
	ProjectId    string
	ProjectName  string
	Daily        float64 // 日常
	Travel       float64 // 差旅
	Bonds        float64 // 奖金
	Insurance    float64 // 保险
	Salary       float64 // 工资
	Rent         float64 // 房租
	Depreciation float64 // 折旧
}

type ProjectDataForDepartment struct {
	PmId         string
	ProjectId    string
	ProjectName  string
	Daily        float64 // 日常
	Travel       float64 // 差旅
	Bonds        float64 // 奖金
	Insurance    float64 // 保险
	Salary       float64 // 工资
	Department   string  // 部门
	Rent         float64 // 房租
	Depreciation float64 // 折旧
}

func (p *ProjectDataForDepartment) Print() {
	fmt.Printf("pmId : %s , projectId: %s,projectName: %s, daily : %f,travel : %f,bonds : %f,insurance : %f,salary : %f, deparment: %s\n",
		p.PmId, p.ProjectId, p.ProjectName, p.Daily, p.Travel, p.Bonds, p.Insurance, p.Salary, p.Department)
}
