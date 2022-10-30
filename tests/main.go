package main

import ping "github.com/budimanlai/go-services-ping"

func main() {
	ping := &ping.ServicePing{}
	ping.Init("config/main.conf", "ping-01")

	ping.Start()
}
