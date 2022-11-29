package service_test

import (
	"errors"
	"testing"

	"github.com/pakawatkung/go-unit-test/repository"
	"github.com/pakawatkung/go-unit-test/service"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCal(t *testing.T) {

	type testCase struct {
		name            string
		PurchaseMin     int
		DiscountPercent int
		amount          int
		want            int
	}

	cases := []testCase{
		{name: "a 100", PurchaseMin: 100, DiscountPercent: 20, amount: 100, want: 80},
		{name: "a 200", PurchaseMin: 100, DiscountPercent: 20, amount: 200, want: 160},
		{name: "a 300", PurchaseMin: 100, DiscountPercent: 20, amount: 300, want: 240},
		{name: "not applied", PurchaseMin: 100, DiscountPercent: 20, amount: 50, want: 50},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Arrage
			promoRepo := repository.NewpromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repository.Promotion{
				ID:              1,
				PurchaseMin:     c.PurchaseMin,
				DiscountPercent: c.DiscountPercent,
			}, nil)

			promoSrv := service.NewPromotionService(promoRepo)

			// Act
			discount, _ := promoSrv.CalculateDiscount(c.amount)
			want := c.want

			// Assert
			assert.Equal(t, want, discount)
		})
	}

	t.Run("perchase amount zero", func(t *testing.T) {
		// Arrage
		promoRepo := repository.NewpromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repository.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoSrv := service.NewPromotionService(promoRepo)

		// Act
		_, err := promoSrv.CalculateDiscount(0)

		// Assert
		assert.ErrorIs(t, err, service.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		// Arrage
		promoRepo := repository.NewpromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repository.Promotion{}, errors.New("repo error"))

		promoSrv := service.NewPromotionService(promoRepo)

		// Act
		_, err := promoSrv.CalculateDiscount(100)

		// Assert
		assert.ErrorIs(t, err, service.ErrRepository)
	})
}
