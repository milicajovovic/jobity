package models

type Review struct {
	ID            uint
	EmployerID    uint
	EmployeeID    uint
	Grade         int
	Comment       string
	Inappropriate bool
	Deleted       bool
}
