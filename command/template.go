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
		Args: []cli.Argument{
			cli.Argument{
				Name: "ids",
				Help: "VM ids",
				Type: cli.Strings,
			},
		},
		Flags: []cli.Flag{
			cli.Flag{
				Name:    "templatefilter",
				Help:    "templatefilter",
				Default: "executable",
				Type:    cli.String,
			},
			cli.Flag{
				Name: "name",
				Help: "The name of the template",
				Type: cli.String,
			},
			cli.Flag{
				Name: "zone",
				Help: "The zoneid or zonename of the template",
				Type: cli.String,
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			client, err := lockgate.GetClient(c)
			if err != nil {
				log.Fatal(err)
			}
			params := cloudstack.ListTemplatesParameter{}

			ids := lockgate.GetArgumentsFromStdin()
			ids = append(ids, c.Command.Arg("ids").Strings()...)

			templatefilter := c.Command.Flag("templatefilter").String()
			if templatefilter != "" {
				params.SetTemplatefilter(templatefilter)
			}
			name := c.Command.Flag("name").String()
			if name != "" {
				params.SetName(name)
			}
			zone := c.Command.Flag("zone").String()
			if zone != "" {
				params.SetZoneid(zone)
			}

			var templates []cloudstack.Template
			if len(ids) > 0 {
				for _, id := range ids {
					params.SetId(id)
					resp, err := client.ListTemplates(params)
					if err != nil {
						fmt.Println(err)
						log.Fatal(err)
					}
					templates = append(templates, resp...)
				}
			} else {
				templates, err = client.ListTemplates(params)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
			}

			w := lockgate.GetTabWriter(c)
			w.Print(templates)
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
