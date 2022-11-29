package handle_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/go-unit-test/handle"
	"github.com/pakawatkung/go-unit-test/service"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		amount := 100
		want := 80
		promoSrv := service.NewPromotionServiceMock()
		promoSrv.On("CalculateDiscount", amount).Return(want, nil)

		promoHandle := handle.NewPromotionHandle(promoSrv)

		app := fiber.New()
		app.Get("/calulate", promoHandle.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calulate?amount=%v", amount), nil)

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(want), string(body))
		}
	})

	t.Run("Bad Request", func(t *testing.T) {
		// Arrange
		amount := 100
		want := 80
		promoSrv := service.NewPromotionServiceMock()
		promoSrv.On("CalculateDiscount", amount).Return(want, nil)

		promoHandle := handle.NewPromotionHandle(promoSrv)

		app := fiber.New()
		app.Get("/calulate", promoHandle.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calulate?amount=%v", strconv.Itoa(amount)), nil) //amount not Integer

		// Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		// Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(want), string(body))
		}
	})
}
