package entities

import "time"

type Delivery struct {
	ID        int
	Customer  *Customer
	Package   *Package
	Address   string
	CreatedAt time.Time
	Status    string
}

func NewDelivery(id int, customer *Customer, packageInfo *Package, address string, createdAt time.Time, status string) *Delivery {
	return &Delivery{
		ID:        id,
		Customer:  customer,
		Package:   packageInfo,
		Address:   address,
		CreatedAt: createdAt,
		Status:    status,
	}
}

func (d *Delivery) UpdateStatus(status string) {
	d.Status = status
}
