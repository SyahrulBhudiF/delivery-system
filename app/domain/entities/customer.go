package entities

type Customer struct {
	ID      int
	Name    string
	Address string
}

func NewCustomer(id int, name string, address string) *Customer {
	return &Customer{
		ID:      id,
		Name:    name,
		Address: address,
	}
}
