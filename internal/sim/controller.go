package sim

import "time"

type Order struct {
	ID    int
	IsVIP bool
}

type Bot struct {
	ID      int
	Current *Order
	BusyEnd time.Time
}

type Controller struct {
	now         time.Time
	processTime time.Duration
	nextOrderID int
	nextBotID   int

	pending  []Order
	complete []Order
	bots     []Bot
	logs     []string
}

func controller() {
}

func (c *Controller) CreateNormalOrder() Order {
	order := Order{
		ID:    c.nextOrderID,
		IsVIP: false,
	}
	c.nextOrderID++
	c.pending = append(c.pending, order)

	return order
}

func (c *Controller) CreateVIPOrder() Order {
	order := Order{
		ID:    c.nextOrderID,
		IsVIP: true,
	}
	// increment nextOrderId
	c.nextOrderID++

	i := 0
	for i < len(c.pending) && c.pending[i].IsVIP {
		i++
	}

	// append order to pending
	c.pending = append(c.pending, order)

	// shift right
	copy(c.pending[i+1:], c.pending[i:])
	c.pending[i] = order //place vip

	return order
}

// add small getter (for testing)
func (c *Controller) PendingOrders() []Order {
	out := make([]Order, len(c.pending))
	copy(out, c.pending)
	return out
}

func (c *Controller) NextOrderID() int {
	return c.nextOrderID
}
