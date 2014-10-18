package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	IPList = cli.Command{
		Name: "list",
		Help: "List ipaddresses",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListPublicIpAddressesParameter{}
			resp, err := client.ListPublicIpAddresses(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(resp)
		},
	}

	IP = cli.Command{
		Name: "ip",
		Help: "Manage ipaddresses",
		Commands: []cli.Command{
			IPList,
		},
	}
)
