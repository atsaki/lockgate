package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	PortforwardingruleList = cli.Command{
		Name: "list",
		Help: "List Portforwardingrule",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListPortForwardingRulesParameter{}
			resp, err := client.ListPortForwardingRules(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(resp)
		},
	}

	Portforwardingrule = cli.Command{
		Name: "portforwardingrule",
		Help: "Manage portforwardingrule",
		Commands: []cli.Command{
			PortforwardingruleList,
		},
	}
)
