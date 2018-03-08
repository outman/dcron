package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/outman/dcron/app"
)

func main() {

	mode := flag.String("mode", "server", "-mode=server|cli")
	port := flag.String("port", "8080", "-port=8080")
	host := flag.String("host", "", "-host=")
	flag.Parse()

	if bytes.EqualFold([]byte(*mode), []byte("cli")) {
		fmt.Println("Starting...cli mode...")
		app.RunCrond()
	} else {
		fmt.Println("Starting...server mode...")
		app.RunServer(*host, *port)
	}
}
