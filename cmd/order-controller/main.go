package main

import (
	"fmt"
	"se-take-home-assignment/common"
	"se-take-home-assignment/internal/sim"

	"github.com/rs/zerolog/log"
)

func main() {
	common.InitLogging()
	fmt.Println("McDonald's Order Management System - Simulation Results")
	fmt.Println()
	log.Info().Msgf("System initialized with 0 bots")
	c := sim.NewController(1001)

	c.CreateNormalOrder() // Order #1001
	c.CreateVIPOrder()    // Order #1002
	c.CreateNormalOrder() // Order #1003

	c.AddBot() // Bot #1 picks up VIP # 1002
	c.AddBot() // Bot #2 picks up VIP # 1001

	// run sim
	for range 11 {
		c.Tick()
	}

	c.CreateVIPOrder() // Order #1004 - idle Bot #2 should pick it up on next tick

	for range 14 {
		c.Tick()
	}

	c.RemoveBot()

	completed := c.CompleteOrder()
	vipCount := 0
	for _, o := range completed {
		if o.IsVIP {
			vipCount++
		}
	}
	normalCount := len(completed) - vipCount

	fmt.Println()
	fmt.Println("Final Status:")
	fmt.Printf("- Total Orders Processed: %d (%d VIP, %d Normal)\n", len(completed), vipCount,
		normalCount)
	fmt.Printf("- Orders Completed: %d\n", len(completed))
	fmt.Printf("- Active Bots: %d\n", len(c.Bots()))
	fmt.Printf("- Pending Orders: %d\n", len(c.PendingOrders()))
}
