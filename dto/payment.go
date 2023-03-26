package dto

import "strconv"

type Payment struct {
	employeeId   string
	daily        string
	travel       string
	bonds        string
	insurance    string
	salary       string
	department   string
	name         string
	rent         string
	depreciation string
}

func (payment *Payment) GetName() string {
	return payment.name
}

func (payment *Payment) GetEmployId() string {
	return payment.employeeId
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

func (payment *Payment) GetRent() float64 {
	res, err := strconv.ParseFloat(payment.rent, 64)
	if err != nil {
		return 0
	}
	return res
}

func (payment *Payment) GetDepreciation() float64 {
	res, err := strconv.ParseFloat(payment.depreciation, 64)
	if err != nil {
		return 0
	}
	return res
}

// set
func (payment *Payment) SetName(_name string) {
	payment.name = _name
}

func (payment *Payment) SetEmployId(_employeeId string) {
	payment.employeeId = _employeeId
}

func (payment *Payment) SetDaily(_daily string) {
	payment.daily = _daily
}

func (payment *Payment) SetTravel(_travel string) {
	payment.travel = _travel
}

func (payment *Payment) SetBonds(_bonds string) {
	payment.bonds = _bonds
}

func (payment *Payment) SetInsurance(_insurance string) {
	payment.insurance = _insurance
}

func (payment *Payment) SetSalary(_salary string) {
	payment.salary = _salary
}

func (payment *Payment) SetDepartment(_department string) {
	payment.department = _department
}

func (payment *Payment) SetRent(_rent string) {
	payment.rent = _rent
}

func (payment *Payment) SetDepreciation(_depreciation string) {
	payment.depreciation = _depreciation
}
