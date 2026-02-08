package main

import (
	"se-take-home-assignment/common"
	"se-take-home-assignment/internal/sim"

	"github.com/rs/zerolog/log"
)

// 1. init controller (startID=1001)
// 2. create normal order → print "[HH:MM:SS] Created Normal Order #1001 - PENDING"
// 3. create VIP order   → print "[HH:MM:SS] Created VIP Order #1002 - PENDING"
// 5. create normal order → print ...
// 6. add bot            → print bot created + what it picked up
// 7. add bot            → print bot created + what it picked up
// 8. loop Tick() for N seconds:
//      - after each tick, check if any bot just completed (compare complete list)
//      - check if any bot just picked up a new order
//      - print events as they happen
// 9. remove bot         → print bot destroyed
// 10. print final summary

func main() {
	common.InitLogging()
	log.Info().Msgf("System initialized with 0 bots")
	c := sim.NewController(1001)

	c.CreateNormalOrder()
	c.CreateVIPOrder()
	c.CreateNormalOrder()
	c.CreateNormalOrder()

	c.AddBot() // Bot #1 picks up VIP # 1002
	c.AddBot() // Bot #2 picks up VIP # 1001

	// run sim
	for i := 0; i < 25; i++ {
		c.Tick()
	}

	c.RemoveBot()
}
