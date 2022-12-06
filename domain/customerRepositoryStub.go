package domain

// adapter
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Tamique", "Ulaanbaatar", "111111", "01.01.2000", "1"},
		{"1002", "Alex", "NY", "222222", "03.04.1999", "2"},
		{"1003", "Klark", "Delhi", "3333333", "31.12.2006", "3"},
		{"1004", "Pragna", "Beijing", "444444", "22.11.1960", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
