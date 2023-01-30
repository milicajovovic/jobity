package models

type Employer struct {
	ID       int
	Email    string
	Password string `json:"-"`
	Name     string
	Address  string
	Deleted  bool
}
