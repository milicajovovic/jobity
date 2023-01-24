package models

type EmployerDTO struct {
	ID             int
	Email          string
	Name           string
	Address        string
	ProfilePicture string
	Deleted        bool
}
