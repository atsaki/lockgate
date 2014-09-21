package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	ServiceOfferingList = cli.Command{
		Name:      "serviceoffering-list",
		ShortName: "list",
		Usage:     "List serviceofferings",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListServiceOfferingsParameter{}
			result, err := client.ListServiceOfferings(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}

	ServiceOffering = cli.Command{
		Name:  "serviceoffering",
		Usage: "Manage serviceoffering",
		Subcommands: []cli.Command{
			ServiceOfferingList,
		},
	}
)