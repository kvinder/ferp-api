package model

// User type
type User struct {
	ID         int
	EmployeeID string
	Username   string
	Password   string
	Name       string
	Sex        string
	Department string
	Email      string
	Telephone  string
	Status     string
	Roles      []string
	CreateDate string
	UpdateDate string
}