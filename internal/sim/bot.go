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

func Tick() {

}
