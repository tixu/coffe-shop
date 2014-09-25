package order

import (
	"fmt"
	"time"
)

var Coffees = [2]string{"Latte", "Expresso"}

type Order interface {
	Type() string
	Time() time.Duration
}

type order struct {
	taype string
	time  time.Duration
}

func (o *order) Type() string {
	return o.taype
}

func (o *order) Time() time.Duration {
	return o.time
}

func New(taype string, time time.Duration) Order {
	return &order{taype: taype, time: time}
}
func (o *order) String() string {
	return fmt.Sprintf("order %s", string(o.Type()))
}
