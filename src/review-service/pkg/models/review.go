package models

type Review struct {
	ID            int
	EmployerID    int
	EmployeeID    int
	Grade         int
	Comment       string
	Inappropriate bool
	Deleted       bool
}
