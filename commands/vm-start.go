package commands

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
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

			var ids []string
			if !termutil.Isatty(os.Stdin.Fd()) {
				scanner := bufio.NewScanner(os.Stdin)
				for scanner.Scan() {
					ids = append(ids, strings.Fields(scanner.Text())...)
				}
			}
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
)
