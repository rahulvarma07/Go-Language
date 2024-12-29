package models

type EmployeeDetails struct {
	EmployeeID         string `json:"employee_id,omitempty" bson:"employee_id"`
	EmployeeName       string `json:"employee_name,omitempty" bson:"employee_name"`
	EmployeeDepartment string `json:"employee_department,omitempty" bson:"employee_department"`
}
