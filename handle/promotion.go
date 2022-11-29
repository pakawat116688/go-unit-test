package handle

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/go-unit-test/service"
)

type promotionHandle struct {
	promoSrv service.PromotionServicce
}

type PromotionHandle interface {
	CalculateDiscount(c *fiber.Ctx) error 
}

func NewPromotionHandle(promoSrv service.PromotionServicce) PromotionHandle {
	return promotionHandle{promoSrv}
}

func (h promotionHandle) CalculateDiscount(c *fiber.Ctx) error  {
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		// return fiber.NewError(fiber.StatusBadRequest, err.Error())
		return c.SendStatus(fiber.StatusBadRequest)
	}
	discount, err := h.promoSrv.CalculateDiscount(amount)
	if err != nil {
		// return fiber.NewError(fiber.StatusNotFound, err.Error())
		c.SendStatus(fiber.StatusNotFound)
	}

	// return c.Status(fiber.StatusOK).SendString(strconv.Itoa(discount))
	return c.SendString(strconv.Itoa(discount))
}