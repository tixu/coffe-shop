package customer

import (
	"coffe-shop/barista"
	"coffe-shop/order"
	"fmt"
	"math/rand"
)

var r = rand.New(rand.NewSource(99))

type Customer interface {
	PlaceOrder() (order.Order, string)
	EnjoyBeverage(order.Order, barista.Barista) string
	Id() int
	String() string
}

type customer struct {
	id int
}

func (c *customer) PlaceOrder() (customerOrder order.Order, thanks string) {
	i := r.Intn(2)
	customerOrder = order.New(order.Coffees[i], 2)
	return customerOrder, fmt.Sprintf("Customer %d orders %s ", c.id, customerOrder)

}

func (c *customer) Id() int {
	return c.id
}

func (c *customer) EnjoyBeverage(order order.Order, provider barista.Barista) string {
	return fmt.Sprintf("Customer %d says Yum and thanks to %v", c.id, provider)
}

func (c *customer) String() string {
	return fmt.Sprintf("Customer %d", c.id)
}

func New(id int) Customer {
	return &customer{id: id}
}

func RandomGroupOf(n int) []Customer {
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = New(i + 1)
	}
	return customers
}
