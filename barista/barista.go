package barista

import (
	"coffe-shop/order"
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(99))

type Barista interface {
	MakeOrder(order.Order) (order.Order, string)
	Id() int
	String() string
}

type barista struct {
	speed int

	id int
}

func (b *barista) MakeOrder(order order.Order) (order.Order, string) {

	time.Sleep(time.Duration(b.speed) * time.Second)
	return order, fmt.Sprintf("Barista %s has prepared an order ", b.String())
}

func (b *barista) Id() int {
	return b.id
}

func (b *barista) String() string {
	return fmt.Sprintf("Barista %d (%d)", b.id, b.speed)
}

var New = func(id int) Barista {
	return &barista{id: id, speed: r.Intn(10)}
}

var RandomGroupOf = func(n int) []Barista {
	baristas := make([]Barista, n)
	for i := 0; i < n; i++ {
		baristas[i] = New(i + 1)
	}
	return baristas
}
