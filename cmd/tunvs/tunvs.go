package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/xiaol-now/tunv/tun"
)

func main() {
	ifce := tun.CreateTun()
	println(ifce.Name(), ifce.IsTAP(), ifce.IsTUN())
	var frame ethernet.Frame
	for {
		frame.Resize(1500)
		n, err := ifce.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
	}
}
