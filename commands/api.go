package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	API = cli.Command{
		Name:  "api",
		Usage: "Execute CloudStack API",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}

			if len(c.Args()) == 0 {
				cli.ShowCommandHelp(c, c.Command.Name)
				return
			}

			command := c.Args()[0]
			var params map[string]string
			result, err := client.Request(command, params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}
			lockgate.PrettyPrint(result)
		},
	}
)
