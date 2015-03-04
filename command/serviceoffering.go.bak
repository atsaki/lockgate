package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	ServiceOfferingList = cli.Command{
		Name: "list",
		Help: "List serviceofferings",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListServiceOfferingsParameter{}
			resp, err := client.ListServiceOfferings(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(resp)
		},
	}

	ServiceOffering = cli.Command{
		Name: "serviceoffering",
		Help: "Manage serviceoffering",
		Commands: []cli.Command{
			ServiceOfferingList,
		},
	}
)
