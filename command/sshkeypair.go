package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	SshkeypairList = cli.Command{
		Name: "list",
		Help: "List ssh keypairs",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListSSHKeyPairsParameter{}
			resp, err := client.ListSSHKeyPairs(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(resp)
		},
	}

	Sshkeypair = cli.Command{
		Name: "sshkeypair",
		Help: "Manage sshkeypair",
		Commands: []cli.Command{
			SshkeypairList,
		},
	}
)
