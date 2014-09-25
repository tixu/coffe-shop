package customer

import (
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	c := New(1)
	order, _ := c.PlaceOrder()
	if order == nil {
		t.Error("go an error")
	}

}
