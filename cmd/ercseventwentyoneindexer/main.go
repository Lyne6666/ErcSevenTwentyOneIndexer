// cmd/ercseventwentyoneindexer/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"ercseventwentyoneindexer/internal/ercseventwentyoneindexer"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := ercseventwentyoneindexer.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
