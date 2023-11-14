package main

import (
	"flag"
	"fmt"

	"github.com/opengo-org/easy-netdisk/cmd"
)

func main() {
	port := flag.Int("port", 8080, "Set which port server run at")
	server_port := fmt.Sprintf(":%d", *port)
	cmd.RunApplication(server_port)
}
