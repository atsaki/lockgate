package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	DeployVirtualMachine = cli.Command{
		Name:  "deploy",
		Usage: "Deploy virtualmachine",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "zone, z",
				Value: "",
				Usage: "The zoneid or zonename of the virtualmachine",
			},
			cli.StringFlag{
				Name:  "serviceoffering, s",
				Value: "",
				Usage: "The serviceofferingid or serviceofferingname of the virtualmachine",
			},
			cli.StringFlag{
				Name:  "template, t",
				Value: "",
				Usage: "The templateid or templatename of the virtualmachine",
			},
			cli.StringFlag{
				Name:  "displayname",
				Value: "",
				Usage: "The displayname of the virtualmachine",
			},
		},
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.DeployVirtualMachineParameter{}
			if c.String("zone") != "" {
				params.SetZoneid(c.String("zone"))
			}
			if c.String("serviceoffering") != "" {
				params.SetServiceofferingid(c.String("serviceoffering"))
			}
			if c.String("template") != "" {
				params.SetTemplateid(c.String("template"))
			}
			if c.String("displayname") != "" {
				params.SetDisplayname(c.String("displayname"))
			}
			vm, err := client.DeployVirtualMachine(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print([]cloudstack.Virtualmachine{vm})
		},
	}
)
