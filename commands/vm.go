package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	VMList = cli.Command{
		Name:      "vm-list",
		ShortName: "list",
		Usage:     "List virtualmachines",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListVirtualMachinesParameter{}
			result, err := client.ListVirtualMachines(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}

	VMStart = cli.Command{
		Name:      "vm-start",
		ShortName: "start",
		Usage:     "Start virtualmachine",
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.StartVirtualMachineParameter{}
			if c.String("id") != "" {
				params.SetId(c.String("id"))
			}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Args()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.StartVirtualMachine(params)
				if err != nil {
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(vms)
		},
	}

	VMStop = cli.Command{
		Name:      "vm-stop",
		ShortName: "stop",
		Usage:     "Stop virtualmachine",
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.StopVirtualMachineParameter{}
			if c.String("id") != "" {
				params.SetId(c.String("id"))
			}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Args()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.StopVirtualMachine(params)
				if err != nil {
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(vms)
		},
	}

	VMDeploy = cli.Command{
		Name:      "vm-deploy",
		ShortName: "deploy",
		Usage:     "Deploy virtualmachine",
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

	VMDestroy = cli.Command{
		Name:      "vm-destroy",
		ShortName: "destroy",
		Usage:     "Destroy virtualmachine",
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.DestroyVirtualMachineParameter{}
			if c.String("id") != "" {
				params.SetId(c.String("id"))
			}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Args()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.DestroyVirtualMachine(params)
				if err != nil {
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(vms)
		},
	}

	VM = cli.Command{
		Name:  "vm",
		Usage: "Manage virtualmachine",
		Subcommands: []cli.Command{
			VMList,
			VMStart,
			VMStop,
			VMDeploy,
			VMDestroy,
		},
	}
)
