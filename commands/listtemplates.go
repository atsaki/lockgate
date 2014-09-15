package commands

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/codegangsta/cli"
)

var (
	ListTemplates = cli.Command{
		Name:      "templates",
		ShortName: "images",
		Usage:     "list templates",
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListTemplatesParameter{}
			params.SetTemplatefilter("executable")
			result, err := client.ListTemplates(params)
			if err != nil {
				fmt.Println(err)
				log.Fatal(err)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(result)
		},
	}
)
