package dto

import "fmt"

type Record struct {
	PmId          string
	ProjectId     string
	ProjectName   string
	Employee      string
	EmployeeId    string
	WorkSpendTime float64
}

func (r *Record) Print() {
	fmt.Printf("pmid: %s, ProjectId: %s, ProjectName: %s, Employee: %s, EmployeeId: %s, WorkSpendTime: %f\n", r.PmId, r.ProjectId, r.ProjectName, r.Employee, r.EmployeeId, r.WorkSpendTime)
}

func (r *Record) PrintWithPrefix(prefix string) {
	fmt.Println(prefix)
	fmt.Printf("pmid: %s, ProjectId: %s, ProjectName: %s, Employee: %s, EmployeeId: %s, WorkSpendTime: %f\n", r.PmId, r.ProjectId, r.ProjectName, r.Employee, r.EmployeeId, r.WorkSpendTime)
}
