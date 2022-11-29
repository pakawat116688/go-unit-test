package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main()  {

	c := CustomerRepository{}
	// ถ้ามีใครก็ตามส่ง ... มาแบบนี้จะได้แบบนี้กลับไป
	c.On("GetCustomer", 1).Return("Pkk", 22, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))
	// check ใส่ function เสร็จค่อยเรียกใช้งาน function นั้นๆ
	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(name, age)
}

type CustomerRepository struct {
	mock.Mock
}

func (m *CustomerRepository) GetCustomer(id int) (name string, age int, err error)  {
	arges := m.Called(id)
	return arges.String(0), arges.Int(1), arges.Error(2)
}