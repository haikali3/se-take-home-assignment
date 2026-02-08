package sim

import "testing"

// go test ./internal/sim/ -v
// go test ./internal/sim/ -run TestAddBotNoPendingOrder -v
func TestAddBotNoPendingOrder(t *testing.T) {
	c := NewController(1001)
	bot := c.AddBot()

	if bot.Current != nil { //if bot has order assigned
		t.Error("bot should have no order")
	}
}

func TestAddBotPendingOrder(t *testing.T) {
	c := NewController(1001)
	order := c.CreateNormalOrder()
	bot := c.AddBot()

	if bot.Current == nil { //if current bot order is 0, test fail
		t.Error("bot should have an order")
	}

	if bot.Current.ID != order.ID {
		t.Errorf("expected order %d, got %d", order.ID, bot.Current.ID)
	}

	if len(c.PendingOrders()) != 0 {
		t.Error("pending orders should be empty")
	}
}

func TestAddBotPicksUpVIPOrderFirst(t *testing.T) {
	c := NewController(1001)
	c.CreateNormalOrder()          // ID 1001
	c.CreateNormalOrder()          // ID 1002
	vipOrder := c.CreateVIPOrder() // ID 1003
	bot := c.AddBot()

	if bot.Current == nil { //if current bot order is 0, test fail
		t.Error("bot should have an order")
	}

	if bot.Current.ID != vipOrder.ID {
		t.Errorf("expected bot pickup VIP order %d, but bot picked up order %d", vipOrder.ID, bot.Current.ID)
	}

	if len(c.PendingOrders()) != 2 {
		t.Errorf("expected 2 pending orders, got %d", len(c.PendingOrders()))
	}
}

func TestMultipleBotsTest(t *testing.T) {
	c := NewController(1001)
	c.CreateNormalOrder() // ID 1001
	c.CreateNormalOrder() // ID 1002
	c.CreateNormalOrder() // ID 1003

	bot1 := c.AddBot()
	bot2 := c.AddBot()

	if bot1.ID == bot2.ID {
		t.Errorf("both bot can have same order bot1: %d bot2: %d", bot1.ID, bot2.ID)
	}

	if bot1.Current.ID != 1001 {
		t.Errorf("bot1 expected to pick up order 1001")
	}

	if bot2.Current.ID != 1002 {
		t.Errorf("bot2 expected to pick up order 1002")
	}

	if len(c.PendingOrders()) != 1 {
		t.Errorf("should have pending orders")
	}
}

func TestTickCompletesOrder(t *testing.T) {
	c := NewController(1001)
	bot := c.AddBot()

}
func TestTickIdleBotPicksUpOrder(t *testing.T) {
}

func TestRemoveBotReturnsOrderToPending(t *testing.T) {
}
