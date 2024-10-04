package handlers

import (
	"delivery-system/app/domain/entities"
	"delivery-system/app/domain/services"
	"delivery-system/app/infrastructure/repositories"
	"delivery-system/app/usecases"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

var deliveryRepo = repositories.NewDeliveryRepository()
var deliveryService = services.NewDeliveryService()

// Create a new delivery
func CreateDelivery(c *fiber.Ctx) error {
	type Request struct {
		CustomerName    string  `json:"customer_name"`
		Address         string  `json:"address"`
		PackageWeight   float64 `json:"package_weight"`
		PackageContents string  `json:"package_contents"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	customer := entities.NewCustomer(1, req.CustomerName, req.Address)
	pkg := entities.NewPackage(1, req.PackageWeight, req.PackageContents)

	createDeliveryUseCase := usecases.NewCreateDeliveryUseCase(deliveryService, deliveryRepo)
	delivery, err := createDeliveryUseCase.Execute(customer, pkg, req.Address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create delivery"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"delivery_id": delivery.ID,
		"status":      delivery.Status,
		"created_at":  delivery.CreatedAt,
	})
}

// Get a delivery by ID
func GetDelivery(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid delivery ID"})
	}

	delivery, err := deliveryRepo.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Delivery not found"})
	}

	return c.JSON(fiber.Map{
		"delivery_id":      delivery.ID,
		"customer_name":    delivery.Customer.Name,
		"package_contents": delivery.Package.Contents,
		"address":          delivery.Address,
		"status":           delivery.Status,
		"created_at":       delivery.CreatedAt,
	})
}
