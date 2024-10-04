package services

import (
	"delivery-system/app/domain/entities"
	"errors"
	"time"
)

type DeliveryService struct{}

func NewDeliveryService() *DeliveryService {
	return &DeliveryService{}
}

func (ds *DeliveryService) CalculateDeliveryCost(pkg *entities.Package, distance float64) (float64, error) {
	if distance <= 0 || pkg.Weight <= 0 {
		return 0, errors.New("invalid distance or weight")
	}

	baseCost := 50.0
	weightCost := pkg.Weight * 10
	distanceCost := distance * 5

	return baseCost + weightCost + distanceCost, nil
}

func (ds *DeliveryService) ScheduleDeliver(delivery *entities.Delivery, scheduledAt time.Time) error {
	if scheduledAt.Before(time.Now()) {
		return errors.New("cannot schedule a delivery in the past")
	}

	delivery.Status = "scheduled"
	return nil
}
