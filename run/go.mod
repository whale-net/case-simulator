module github.com/whale-net/case-simulator/run

go 1.22.2

replace github.com/whale-net/case-simulator/simulator => ../simulator

require github.com/whale-net/case-simulator/simulator v0.0.0-00010101000000-000000000000

require (
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
)
