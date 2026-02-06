package sim

import (
	"testing"
)

func TestOrderIDIncreasing(t *testing.T) {
	c := NewController(1001)

	order1 := c.CreateNormalOrder()
	order2 := c.CreateVIPOrder()
	order3 := c.CreateNormalOrder()

	if order1.ID != 1001 || order2.ID != 1002 || order3.ID != 1003 {

	}

}
