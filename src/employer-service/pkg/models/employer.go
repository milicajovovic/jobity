package models

type Employer struct {
	ID             int
	Email          string
	Password       string
	Name           string
	Address        string
	ProfilePicture string
	Deleted        bool
}

func (employer *Employer) EmployerToDTO() EmployerDTO {
	return EmployerDTO{
		ID:             employer.ID,
		Email:          employer.Email,
		Name:           employer.Name,
		Address:        employer.Address,
		ProfilePicture: employer.ProfilePicture,
		Deleted:        employer.Deleted,
	}
}
