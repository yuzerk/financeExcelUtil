package main

import "strconv"

type Payment struct {
	employeeId string
	daily      string
	travel     string
	bonds      string
	insurance  string
	salary     string
	department string
}

func (payment *Payment) GetDaily() float64 {
	res, err := strconv.ParseFloat(payment.daily, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetTravel() float64 {
	res, err := strconv.ParseFloat(payment.travel, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetBonds() float64 {
	res, err := strconv.ParseFloat(payment.bonds, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetInsurance() float64 {
	res, err := strconv.ParseFloat(payment.insurance, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetSalary() float64 {
	res, err := strconv.ParseFloat(payment.salary, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetDepartment() string {
	return payment.department
}
