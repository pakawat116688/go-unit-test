package service

import (

	"github.com/pakawatkung/go-unit-test/repository"
)

type promotionService struct {
	promoRepo  repository.PromotionRepository
}

type PromotionServicce interface {
	CalculateDiscount(amount int) (int, error)
}

func NewPromotionService(promoRepo  repository.PromotionRepository) PromotionServicce {
	return promotionService{promoRepo}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {
	
	if amount <= 0 {
		return 0, ErrZeroAmount 
	}

	promotion, err := s.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}