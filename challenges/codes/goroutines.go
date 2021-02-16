package main

import (
	"fmt"
)

// PAYLOAD
type Payload struct {
	ChannelID string
	Value     int
}

// A GO
func aGo(a chan Payload) {
	inc := 0
	for {
		inc++
		a <- Payload{"a", inc}
	}
}

// B GO
func bGo(a chan Payload) {
	inc := 0
	for {
		inc++
		a <- Payload{"b", inc}
	}
}

// MAIN
func main() {

	pChan := make(chan Payload)

	go aGo(pChan)
  go bGo(pChan)

  aCount, bCount := 0, 0

	for p := range pChan {
		fmt.Printf("Channel ID: %v, Value: %5v\n", p.ChannelID, p.Value)
    if p.ChannelID == "a" { aCount++ }
    if p.ChannelID == "b" { bCount++ }
		if p.Value == 100 { break }
	}

  fmt.Printf("Total: %v\n", aCount + bCount)
  fmt.Printf("A Receives: %v\n", aCount)
  fmt.Printf("B Receives: %v\n", bCount)

	close(pChan)
}
