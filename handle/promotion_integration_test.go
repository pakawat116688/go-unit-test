//go:build integration

package handle_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/go-unit-test/handle"
	"github.com/pakawatkung/go-unit-test/repository"
	"github.com/pakawatkung/go-unit-test/service"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateIntegration(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		amount := 100
		want := 80

		promoRepo := repository.NewpromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repository.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoSrv := service.NewPromotionService(promoRepo)
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

}
