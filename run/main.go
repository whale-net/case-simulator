package main

import (
	"github.com/whale-net/case-simulator/simulator"
)

func main() {
	//simulator.Simulate(nil, nil)
	simulator.SimulateMany(100_000)
}
