package models

type Grade struct {
	ID            uint
	EmployerID    uint
	Grade         int
	Comment       string
	Inappropriate bool
}
