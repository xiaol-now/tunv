package main

import "github.com/xiaol-now/tunv/tun"

func main() {
	ifce := tun.CreateTun()
	println(ifce.Name(), ifce.IsTAP(), ifce.IsTUN())
	select {}
}
