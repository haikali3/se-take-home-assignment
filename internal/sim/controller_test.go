package sim

import (
	"se-take-home-assignment/common"
	"testing"
)

func TestOrderIDIncreasing(t *testing.T) {
	common.InitLogging()

	c := NewController(1001)

	order1 := c.CreateNormalOrder()
	order2 := c.CreateVIPOrder()
	order3 := c.CreateNormalOrder()

	if order1.ID != 1001 || order2.ID != 1002 || order3.ID != 1003 {
		t.Fatalf("expected sequential IDs 1001, 1002, 1003 but got %d, %d, %d", order1.ID, order2.ID, order3.ID)
	}
}

func TestVIPOrderPriority(t *testing.T) {
	c := NewController(1001)

	c.CreateNormalOrder()     // ID 1001
	c.CreateNormalOrder()     // ID 1002
	vip := c.CreateVIPOrder() // ID 1003

	pending := c.PendingOrders()

	// VIP should be 1st in queue
	if pending[0].ID != vip.ID {
		t.Fatalf("expected VIP order ID %d to be first in queue but got ID %d", vip.ID, pending[0].ID)
	}
}
