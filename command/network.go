package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	NetworkList = cli.Command{
		Name: "list",
		Help: "List network",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListNetworksParameter{}
			result, err := client.ListNetworks(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}

	Network = cli.Command{
		Name: "network",
		Help: "Manage network",
		Commands: []cli.Command{
			NetworkList,
		},
	}
)
