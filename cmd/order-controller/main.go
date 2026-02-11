package main

import (
	"fmt"
	"se-take-home-assignment/common"
	"se-take-home-assignment/internal/cli"
)

func main() {
	common.InitLogging()
	fmt.Println("McDonald's Order Management System - Simulation Results")
	cli.New().Run()
}
