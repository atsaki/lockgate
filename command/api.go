package command

import (
	"fmt"
	"log"
	"os"

	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	API = cli.Command{
		Name: "api",
		Help: "Execute CloudStack API",
		Args: []cli.Argument{
			cli.Argument{
				Name:     "command",
				Help:     "CloudStack API command",
				Required: true,
				Type:     cli.String,
			},
			cli.Argument{
				Name: "params",
				Help: "API parameters. key1=val1 key2=val2 ...",
				Type: cli.StringMap,
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			command := c.Command.Arg("command").Value().(string)
			params := c.Command.Arg("params").Value().(map[string]string)

			client, err := lockgate.GetClient(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal(err)
			}

			result, err := client.Request(command, params)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal(err)
			}
			lockgate.PrettyPrint(result)
		},
	}
)
