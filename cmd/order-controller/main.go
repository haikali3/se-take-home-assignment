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

	order1 := c.CreateNormalOrder()
	log.Info().Msgf("Created Normal Order #%d - Status: PENDING", order1.ID)
	order2 := c.CreateVIPOrder()
	log.Info().Msgf("Created VIP Order #%d - Status: PENDING", order2.ID)
	order3 := c.CreateVIPOrder()
	c.CreateNormalOrder()
	log.Info().Msgf("Created Normal Order #%d - Status: PENDING", order3.ID)

}
