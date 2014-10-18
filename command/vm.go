package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	VMList = cli.Command{
		Name: "list",
		Help: "List virtualmachines",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListVirtualMachinesParameter{}
			resp, err := client.ListVirtualMachines(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			items := make([]interface{}, len(resp))
			for i, r := range resp {
				items[i] = r
			}

			w := lockgate.GetTabWriter(c)
			w.Print(items)
		},
	}

	VMStart = cli.Command{
		Name: "start",
		Help: "Start virtualmachine",
		Args: []cli.Argument{
			cli.Argument{
				Name: "ids",
				Help: "VM ids",
				Type: cli.Strings,
			},
		},
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.StartVirtualMachineParameter{}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Command.Arg("ids").Strings()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.StartVirtualMachine(params)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			items := make([]interface{}, len(vms))
			for i, vm := range vms {
				items[i] = vm
			}

			w := lockgate.GetTabWriter(c)
			w.Print(items)
		},
	}

	VMStop = cli.Command{
		Name: "stop",
		Help: "Stop virtualmachine",
		Args: []cli.Argument{
			cli.Argument{
				Name: "ids",
				Help: "VM ids",
				Type: cli.Strings,
			},
		},
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.StopVirtualMachineParameter{}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Command.Arg("ids").Strings()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.StopVirtualMachine(params)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			items := make([]interface{}, len(vms))
			for i, vm := range vms {
				items[i] = vm
			}

			w := lockgate.GetTabWriter(c)
			w.Print(items)
		},
	}

	VMDeploy = cli.Command{
		Name: "deploy",
		Help: "Deploy virtualmachine",
		Flags: []cli.Flag{
			cli.Flag{
				Name:     "zone",
				Short:    'z',
				Help:     "The zoneid or zonename of the virtualmachine",
				Type:     cli.String,
				Required: true,
			},
			cli.Flag{
				Name:     "serviceoffering",
				Short:    's',
				Help:     "The serviceofferingid or serviceofferingname of the virtualmachine",
				Type:     cli.String,
				Required: true,
			},
			cli.Flag{
				Name:     "template",
				Short:    't',
				Help:     "The templateid or templatename of the virtualmachine",
				Type:     cli.String,
				Required: true,
			},
			cli.Flag{
				Name: "name",
				Help: "The name of the virtualmachine",
				Type: cli.String,
			},
			cli.Flag{
				Name: "displayname",
				Help: "The displayname of the virtualmachine",
				Type: cli.String,
			},
			cli.Flag{
				Name: "keypair",
				Help: "The keypair of the virtualmachine",
				Type: cli.String,
			},
		},
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.DeployVirtualMachineParameter{}

			zone := c.Command.Flag("zone").String()
			if zone != "" {
				params.SetZoneid(zone)
			}
			serviceoffering := c.Command.Flag("serviceoffering").String()
			if serviceoffering != "" {
				params.SetServiceofferingid(serviceoffering)
			}
			template := c.Command.Flag("template").String()
			if template != "" {
				params.SetTemplateid(template)
			}
			name := c.Command.Flag("name").String()
			if name != "" {
				params.SetName(name)
			}
			displayname := c.Command.Flag("displayname").String()
			if displayname != "" {
				params.SetDisplayname(displayname)
			}
			keypair := c.Command.Flag("keypair").String()
			if keypair != "" {
				params.SetKeypair(keypair)
			}

			vm, err := client.DeployVirtualMachine(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print([]interface{}{vm})
		},
	}

	VMDestroy = cli.Command{
		Name: "destroy",
		Help: "Destroy virtualmachine",
		Args: []cli.Argument{
			cli.Argument{
				Name: "ids",
				Help: "VM ids",
				Type: cli.Strings,
			},
		},
		Action: func(c *cli.Context) {
			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.DestroyVirtualMachineParameter{}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Command.Arg("ids").Strings()...)

			log.Println("ids:", ids)
			vms := []cloudstack.Virtualmachine{}
			for _, id := range ids {
				params.SetId(id)
				vm, err := client.DestroyVirtualMachine(params)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				vms = append(vms, vm)
			}

			items := make([]interface{}, len(vms))
			for i, vm := range vms {
				items[i] = vm
			}

			w := lockgate.GetTabWriter(c)
			w.Print(items)
		},
	}

	VM = cli.Command{
		Name: "vm",
		Help: "Manage virtualmachine",
		Commands: []cli.Command{
			VMList,
			VMStart,
			VMStop,
			VMDeploy,
			VMDestroy,
		},
	}
)
