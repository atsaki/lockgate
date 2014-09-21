package command

import (
	"fmt"
	"log"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	TemplateList = cli.Command{
		Name: "list",
		Help: "List templates",
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

	Template = cli.Command{
		Name: "template",
		Help: "Manage template",
		Commands: []cli.Command{
			TemplateList,
		},
	}
)
