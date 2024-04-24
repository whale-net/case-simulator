package main

import (
	"github.com/whale-net/case-simulator/simulator"
)

func main() {
	//simulator.Simulate(nil, nil)
	simulator.Simulate(10, 100_000_000)
}
