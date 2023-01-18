package models

type Employer struct {
	ID       uint
	Email    string
	Password string
	Name     string
	Address  string
}

func (employer *Employer) EmployerToDTO() EmployerDTO {
	return EmployerDTO{
		ID:      employer.ID,
		Email:   employer.Email,
		Name:    employer.Name,
		Address: employer.Address,
	}
}
