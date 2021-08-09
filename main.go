package main

import (
 "os"
 "log"
 "fmt"
 cli "github.com/urfave/cli/v2"
)


func main() {
	app := &cli.App{
		Name: "vCard: matt-forster",
		Usage: "A virtual business card, by a developer for developers",
		Action: func (context *cli.Context) error {
			fmt.Println("This is strange.")
			return nil
		},
	}

	err := app.Run(os.Args)
	if (err != nil) {
		log.Fatal(err)
	}
}

