package usecases

import (
	"delivery-system/app/domain/entities"
	"delivery-system/app/domain/services"
	"delivery-system/app/infrastructure/repositories"
	"time"
)

type CreateDeliveryUseCase struct {
	deliveryService *services.DeliveryService
	deliveryRepo    *repositories.DeliveryRepository
}

func NewCreateDeliveryUseCase(db *services.DeliveryService, dr *repositories.DeliveryRepository) *CreateDeliveryUseCase {
	return &CreateDeliveryUseCase{
		deliveryService: db,
		deliveryRepo:    dr,
	}
}

func (uc *CreateDeliveryUseCase) Execute(customer *entities.Customer, pkg *entities.Package, address string) (*entities.Delivery, error) {
	delivery := entities.NewDelivery(1, customer, pkg, address, time.Now(), "pending")

	err := uc.deliveryRepo.Save(delivery)
	if err != nil {
		return nil, err
	}

	return delivery, nil
}
