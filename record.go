package main

import "fmt"

type Record struct {
	pmId          string
	projectId     string
	projectName   string
	employee      string
	employeeId    string
	workSpendTime float64
}

func (r *Record) Print() {
	fmt.Printf("pmid: %s, projectId: %s, projectName: %s, employee: %s, employeeId: %s, workSpendTime: %f\n", r.pmId, r.projectId, r.projectName, r.employee, r.employeeId, r.workSpendTime)
}
