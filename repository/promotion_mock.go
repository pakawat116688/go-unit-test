package repository

import "github.com/stretchr/testify/mock"

type promotionRepositoryMock struct {
	mock.Mock
}

func NewpromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (m *promotionRepositoryMock) GetPromotion() (Promotion, error)  {
	arg := m.Called()
	return arg.Get(0).(Promotion), arg.Error(1)
}