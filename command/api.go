package command

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
)

var (
	API = cli.Command{
		Name: "api",
		Help: "Execute CloudStack API",
		Args: []cli.Argument{
			cli.Argument{
				Name:     "command",
				Help:     "CloudStack API command",
				Required: true,
				Type:     cli.String,
			},
			cli.Argument{
				Name: "params",
				Help: "API parameters. key1=val1 key2=val2 ...",
				Type: cli.StringMap,
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			command := c.Command.Arg("command").Value().(string)
			params := c.Command.Arg("params").Value().(map[string]string)

			client, err := lockgate.GetClient(c)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal(err)
			}

			result, err := client.Request(command, params)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal(err)
			}

			items := []interface{}{}
			var key string
			var value json.RawMessage

			temp := map[string]json.RawMessage{}
			err = json.Unmarshal(result, &temp)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				log.Fatal(err)
			}
			for key, value = range temp {
				if key != "count" {
					break
				}
			}

			if key == "success" {
				if string(value) == "true" {
					items = append(items, map[string]bool{"success": true})
				} else {
					items = append(items, map[string]bool{"success": false})
				}
			} else if strings.HasPrefix(command, "list") {
				err = json.Unmarshal(value, &items)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					log.Fatal(err)
				}
			} else {
				var item interface{}
				err := json.Unmarshal(value, &item)
				if err != nil {
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						log.Fatal(err)
					}
				}
				items = append(items, item)
			}

			w := lockgate.GetTabWriter(c)
			w.Print(items)
		},
	}
)
