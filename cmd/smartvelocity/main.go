// cmd/smartvelocity/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartvelocity/internal/smartvelocity"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartvelocity.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
