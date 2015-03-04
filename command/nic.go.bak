package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	NicList = cli.Command{
		Name: "list",
		Help: "List nics",
		Args: []cli.Argument{
			cli.Argument{
				Name: "virtualmachineid",
				Help: "VM id",
				Type: cli.Strings,
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListNicsParameter{}

			vmids := lockgate.GetArgumentsFromStdin()
			vmids = append(vmids, c.Command.Arg("virtualmachineid").Strings()...)

			nics := []cloudstack.Nic{}
			for _, vmid := range vmids {
				params.SetVirtualmachineid(vmid)
				resp, err := client.ListNics(params)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				nics = append(nics, resp...)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(nics)
		},
	}

	Nic = cli.Command{
		Name: "nic",
		Help: "Manage nics",
		Commands: []cli.Command{
			NicList,
		},
	}
)
