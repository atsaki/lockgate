package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	ZoneList = cli.Command{
		Name:      "zone-list",
		ShortName: "list",
		Usage:     "List zones",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListZonesParameter{}
			result, err := client.ListZones(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}

	Zone = cli.Command{
		Name:  "zone",
		Usage: "Manage zone",
		Subcommands: []cli.Command{
			ZoneList,
		},
	}
)
