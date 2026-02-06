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
	// read current nextOrderId as new order ID
	order := Order{
		ID:    c.nextOrderID,
		IsVIP: false,
	}
	// build normal order (vip: valse)
	// increment nextOrderId
	c.nextOrderID++
	// append order to pending

	c.pending = append(c.pending, order)

	// return created order
	return order
}

func CreateVIPOrder() {

}

// add small getter (for testing)
func PendingOrders() {
}

func NextOrderID() {
}
