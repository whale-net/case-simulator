package main

import (
	"github.com/whale-net/case-simulator/simulator"
)

func main() {
	//simulator.Simulate(nil, nil)
	//simulator.Simulate(10, 400)
	// 2 min 35 seconds
	//simulator.Simulate(10, 1_000_000_000)

	// 2min6s
	// 10 billion = 126 seconds
	// 100 billion = 1260
	// 1 trillion = 12600
	// 12600 / 3600 = ~4 hours
	//simulator.Simulate(16, 625_000_000)
	simulator.Simulate(16, 62_500_000_000)
}
