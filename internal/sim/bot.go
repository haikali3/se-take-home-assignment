package sim

func (c *Controller) AddBot() Bot {
	// create bot with next ID
	bot := Bot{
		ID: c.nextBotID,
	}
	c.nextBotID++

	if len(c.pending) > 0 {
		order := c.pending[0]
		c.pending = c.pending[1:] // [1:] the slice drop the first element
		bot.Current = &order
		bot.BusyEnd = c.now.Add(c.processTime)
	}

	c.bots = append(c.bots, bot)
	return bot
}

func (c *Controller) RemoveBot() {
	if len(c.bots) == 0 {
		return
	}

	// c.bots = [Bot0, Bot1, Bot2]
	// remove the newest bot on last position
	newestBot := c.bots[len(c.bots)-1]
	c.bots = c.bots[:len(c.bots)-1]

	// if the bot is processing an order, put it back to pending
	if newestBot.Current != nil {
		// prepend back to front of pending (VIP priority still applies)
		c.pending = append([]Order{*newestBot.Current}, c.pending...)
	}
}

// simulation clock by one time step
func (c *Controller) Tick() {
	// 1. complete finish orders - check every bot
	for i := range c.bots {
		bot := &c.bots[i]
		if bot.Current != nil && !c.now.Before(bot.BusyEnd) { // if current bot is busy and the time has reached or passed the BusyEnd
			// order is complete
			c.complete = append(c.complete, *bot.Current)
			bot.Current = nil
		}
	}

	// 2. assign pending orders to idle bots
	for idleBot := range c.bots {
		bot := &c.bots[idleBot]
		if bot.Current == nil && len(c.pending) > 0 { // if bot has no order and order pending more than 0
			order := c.pending[0]     // get first pending order
			c.pending = c.pending[1:] // and remove the first slice of array
			bot.Current = &order
			bot.BusyEnd = c.now.Add(c.processTime)
		}
	}

	// 3. advance simulation time
	c.now = c.now.Add(c.timeStep)
}
