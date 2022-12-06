package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
