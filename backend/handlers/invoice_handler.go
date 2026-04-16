package handlers

import (
	"fleetify-backend/config"
	"fleetify-backend/models"
	"fleetify-backend/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateInvoice(c *fiber.Ctx) error {
	type InvoiceRequest struct {
		SenderName  string `json:"sender_name"`
		SenderAddress string `json:"sender_address"`
		ReceiverName string `json:"receiver_name"`
		ReceiverAddress string `json:"receiver_address"`
		Items []struct {
			ItemID uint `json:"item_id"`
			Quantity int `json:"quantity"`
		} `json:"items"`
	}

	var req InvoiceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid JSON format"})
	}

	if req.SenderName == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Sender name is required"})
	}

	if req.SenderAddress == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Sender address is required"})
	}

	if req.ReceiverName == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Receiver name is required"})
	}

	if req.ReceiverAddress == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Receiver address is required"})
	}

	if len(req.Items) == 0 {
		return c.Status(400).JSON(fiber.Map{"message": "Invoice must contain at least 1 item"})
	}

	userID := c.Locals("user_id").(float64)

	var finalInvoice models.Invoice

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var grandTotal int64 = 0
		var invoiceDetails []models.InvoiceDetail

		for _, itemReq := range req.Items {
			var item models.Item
			if err := tx.First(&item, itemReq.ItemID).Error; err != nil {
				return fmt.Errorf("item with ID %d not found", itemReq.ItemID)
			}

			subTotal := int64(itemReq.Quantity) * item.Price
			grandTotal += subTotal

			invoiceDetails = append(invoiceDetails, models.InvoiceDetail{
				ItemID: item.ID,
				Quantity: itemReq.Quantity,
				Price: item.Price,
				SubTotal: subTotal,
			})
		}

		finalInvoice = models.Invoice{
			InvoiceNumber: fmt.Sprintf("INV-%d", time.Now().Unix()),
			SenderName: req.SenderName,
			SenderAddress: req.SenderAddress,
			ReceiverName: req.ReceiverName,
			ReceiverAddress: req.ReceiverAddress,
			TotalAmount: grandTotal,
			CreatedBy: uint(userID),
			Details: invoiceDetails,
		} 

		if err := tx.Create(&finalInvoice).Error; err != nil {
			return err
		}

		return nil
	})
	
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	utils.SendInvoiceWebhook(finalInvoice)

	return c.Status(201).JSON(fiber.Map{"message": "Invoice created successfully"})
}