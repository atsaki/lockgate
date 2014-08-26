package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/user"
	"strings"
	"text/tabwriter"

	"code.google.com/p/gcfg"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/codegangsta/cli"
)

const (
	configfile = "~/.cloudmonkey/config"
)

type Config struct {
	User struct {
		Username  string
		Password  string
		Secretkey string
		Apikey    string
	}
	Server struct {
		Protocol string
		Host     string
		Port     string
		Path     string
	}
}

func expandPath(path string) string {
	usr, _ := user.Current()
	home := usr.HomeDir

	if strings.HasPrefix(path, "~/") {
		path = strings.Replace(path, "~/", home+"/", 1)
	}
	return path
}

func main() {

	log.SetOutput(ioutil.Discard)

	cfg := Config{}

	err := gcfg.ReadFileInto(&cfg, expandPath(configfile))
	if err != nil {
		log.Fatal(err)
	}

	endpoint := url.URL{
		Scheme: cfg.Server.Protocol,
		Host:   fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Path:   cfg.Server.Path,
	}

	client, _ := cloudstack.NewClient(endpoint,
		cfg.User.Apikey, cfg.User.Secretkey, cfg.User.Username, cfg.User.Password)

	sep := '\t'
	tabw := new(tabwriter.Writer)
	tabw.Init(os.Stdout, 0, 8, 0, byte(sep), 0)

	app := cli.NewApp()
	app.Name = "lg"
	app.Usage = "lg comand"

	app.Commands = []cli.Command{
		{
			Name:      "virtualmachines",
			ShortName: "vms",
			Usage:     "list virtualmachines",
			Action: func(c *cli.Context) {
				params := cloudstack.ListVirtualMachinesParameter{}
				resp, _ := client.ListVirtualMachines(params)
				for _, v := range resp.Virtualmachine {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Displayname.String,
								v.State.String,
								v.Zonename.String,
								v.Templatename.String,
								v.Serviceofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:  "deploy",
			Usage: "deploy virtualmachine",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "zone",
					Value: "",
				},
				cli.StringFlag{
					Name:  "serviceoffering",
					Value: "",
				},
				cli.StringFlag{
					Name:  "template",
					Value: "",
				},
				cli.StringFlag{
					Name:  "displayname",
					Value: "",
				},
			},
			Action: func(c *cli.Context) {
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
				resp, _ := client.DeployVirtualMachine(params)
				v := resp.Virtualmachine
				fmt.Fprintln(
					tabw,
					strings.Join(
						[]string{
							v.Id.String,
							v.Name.String,
							v.Displayname.String,
							v.State.String,
							v.Zonename.String,
							v.Templatename.String,
							v.Serviceofferingname.String,
						}, string(sep)))
				tabw.Flush()
			},
		},
		{
			Name:  "start",
			Usage: "start virtualmachine",
			Action: func(c *cli.Context) {
				params := cloudstack.StartVirtualMachineParameter{}
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						id := strings.Split(scanner.Text(), string(sep))[0]
						params.SetId(id)
						resp, _ := client.StartVirtualMachine(params)
						v := resp.Virtualmachine
						fmt.Fprintln(
							tabw,
							strings.Join(
								[]string{
									v.Id.String,
									v.Name.String,
									v.Displayname.String,
									v.State.String,
									v.Zonename.String,
									v.Templatename.String,
									v.Serviceofferingname.String,
								}, string(sep)))
					}
					tabw.Flush()
				} else {
					params.SetId(c.Args()[0])
					resp, _ := client.StartVirtualMachine(params)
					v := resp.Virtualmachine
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Displayname.String,
								v.State.String,
								v.Zonename.String,
								v.Templatename.String,
								v.Serviceofferingname.String,
							}, string(sep)))
					tabw.Flush()
				}
			},
		},
		{
			Name:  "stop",
			Usage: "stop virtualmachine",
			Action: func(c *cli.Context) {
				params := cloudstack.StopVirtualMachineParameter{}
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						id := strings.Split(scanner.Text(), string(sep))[0]
						params.SetId(id)
						resp, _ := client.StopVirtualMachine(params)
						v := resp.Virtualmachine
						fmt.Fprintln(
							tabw,
							strings.Join(
								[]string{
									v.Id.String,
									v.Name.String,
									v.Displayname.String,
									v.State.String,
									v.Zonename.String,
									v.Templatename.String,
									v.Serviceofferingname.String,
								}, string(sep)))
					}
					tabw.Flush()
				} else {
					params.SetId(c.Args()[0])
					resp, _ := client.StopVirtualMachine(params)
					v := resp.Virtualmachine
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Displayname.String,
								v.State.String,
								v.Zonename.String,
								v.Templatename.String,
								v.Serviceofferingname.String,
							}, string(sep)))
					tabw.Flush()
				}
			},
		},
		{
			Name:  "destroy",
			Usage: "destroy virtualmachine",
			Action: func(c *cli.Context) {
				params := cloudstack.DestroyVirtualMachineParameter{}
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						id := strings.Split(scanner.Text(), string(sep))[0]
						params.SetId(id)
						resp, _ := client.DestroyVirtualMachine(params)
						v := resp.Virtualmachine
						fmt.Fprintln(
							tabw,
							strings.Join(
								[]string{
									v.Id.String,
									v.Name.String,
									v.Displayname.String,
									v.State.String,
									v.Zonename.String,
									v.Templatename.String,
									v.Serviceofferingname.String,
								}, string(sep)))
					}
					tabw.Flush()
				} else {
					params.SetId(c.Args()[0])
					resp, _ := client.DestroyVirtualMachine(params)
					v := resp.Virtualmachine
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Displayname.String,
								v.State.String,
								v.Zonename.String,
								v.Templatename.String,
								v.Serviceofferingname.String,
							}, string(sep)))
					tabw.Flush()
				}
			},
		},
		{
			Name:  "zones",
			Usage: "list zones",
			Action: func(c *cli.Context) {
				params := cloudstack.ListZonesParameter{}
				resp, _ := client.ListZones(params)
				for _, v := range resp.Zone {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:      "serviceofferings",
			ShortName: "sizes",
			Usage:     "list serviceofferings",
			Action: func(c *cli.Context) {
				params := cloudstack.ListServiceOfferingsParameter{}
				resp, _ := client.ListServiceOfferings(params)
				for _, v := range resp.Serviceoffering {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:      "templates",
			ShortName: "images",
			Usage:     "list templates",
			Action: func(c *cli.Context) {
				params := cloudstack.ListTemplatesParameter{}
				params.SetTemplatefilter("featured")
				resp, _ := client.ListTemplates(params)
				for _, v := range resp.Template {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Displaytext.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:      "networks",
			ShortName: "nws",
			Usage:     "list network",
			Action: func(c *cli.Context) {
				params := cloudstack.ListNetworksParameter{}
				resp, _ := client.ListNetworks(params)
				for _, v := range resp.Network {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								v.Id.String,
								v.Name.String,
								v.Networkofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
	}

	app.Run(os.Args)
}
