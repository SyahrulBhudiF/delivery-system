package repositories

import (
	"delivery-system/app/domain/entities"
	"errors"
)

type DeliveryRepository struct {
	deliveries []entities.Delivery
}

func NewDeliveryRepository() *DeliveryRepository {
	return &DeliveryRepository{
		deliveries: []entities.Delivery{},
	}
}

func (r *DeliveryRepository) Save(delivery *entities.Delivery) error {
	delivery.ID = len(r.deliveries) + 1
	r.deliveries = append(r.deliveries, *delivery)
	return nil
}

func (r *DeliveryRepository) FindByID(id int) (*entities.Delivery, error) {
	for _, delivery := range r.deliveries {
		if delivery.ID == id {
			return &delivery, nil
		}
	}
	return nil, errors.New("Delivery not found")
}
