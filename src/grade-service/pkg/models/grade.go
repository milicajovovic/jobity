package models

type Grade struct {
	ID            uint
	EmployerID    uint
	EmployeeID    uint
	Grade         int
	Comment       string
	Inappropriate bool
}
