package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	FirewallruleList = cli.Command{
		Name: "list",
		Help: "List firewallrule",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListFirewallRulesParameter{}
			result, err := client.ListFirewallRules(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}

	Firewallrule = cli.Command{
		Name: "firewallrule",
		Help: "Manage firewallrule",
		Commands: []cli.Command{
			FirewallruleList,
		},
	}
)
