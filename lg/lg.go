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

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/codegangsta/cli"
	"github.com/vaughan0/go-ini"
)

const (
	sep = '\t'
)

var (
	client *cloudstack.Client
)

func expandPath(path string) string {
	usr, _ := user.Current()
	home := usr.HomeDir

	if strings.HasPrefix(path, "~/") {
		path = strings.Replace(path, "~/", home+"/", 1)
	}
	return path
}

func setup(c *cli.Context) {

	if !c.GlobalBool("debug") {
		log.SetOutput(ioutil.Discard)
	}

	configfile := expandPath(c.GlobalString("config-file"))
	log.Println("configfile:", configfile)
	cfg, err := ini.LoadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	var ok bool

	profile, ok := cfg.Get("core", "profile")
	if !ok {
		profile = "local"
	}
	if c.GlobalString("profile") != "" {
		profile = c.GlobalString("profile")
	}
	log.Println("profile:", profile)

	endpointUrl, ok := cfg.Get(profile, "url")
	if !ok {
		log.Fatalf("URL is not specified")
	}
	log.Println("url:", endpointUrl)

	apikey, ok := cfg.Get(profile, "apikey")
	if !ok {
		apikey = ""
	}
	log.Println("apikey:", apikey)

	secretkey, ok := cfg.Get(profile, "secretkey")
	if !ok {
		secretkey = ""
	}
	log.Println("secretkey:", secretkey)

	username, ok := cfg.Get(profile, "username")
	if !ok {
		username = ""
	}
	log.Println("username:", username)

	password, ok := cfg.Get(profile, "password")
	if !ok {
		password = ""
	}
	log.Println("password:", password)

	endpoint, err := url.Parse(endpointUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("endpoint:", endpoint)

	client, err = cloudstack.NewClient(*endpoint, apikey, secretkey,
		username, password)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	tabw := new(tabwriter.Writer)
	tabw.Init(os.Stdout, 0, 8, 0, byte(sep), 0)

	app := cli.NewApp()
	app.Name = "lg"
	app.Usage = "lg comand"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config-file, c",
			Value: "~/.cloudmonkey/config",
			Usage: "Config file path",
		},
		cli.StringFlag{
			Name:  "profile, P",
			Value: "",
			Usage: "Server profile",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Show debug messages",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "virtualmachines",
			ShortName: "vms",
			Usage:     "List virtualmachines",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.ListVirtualMachinesParameter{}
				vms, err := client.ListVirtualMachines(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, vm := range vms {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								vm.Id.String,
								vm.Name.String,
								vm.Displayname.String,
								vm.State.String,
								vm.Zonename.String,
								vm.Templatename.String,
								vm.Serviceofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:  "deploy",
			Usage: "Deploy virtualmachine",
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
				setup(c)
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
					log.Fatal(err)
				}
				fmt.Fprintln(
					tabw,
					strings.Join(
						[]string{
							vm.Id.String,
							vm.Name.String,
							vm.Displayname.String,
							vm.State.String,
							vm.Zonename.String,
							vm.Templatename.String,
							vm.Serviceofferingname.String,
						}, string(sep)))
				tabw.Flush()
			},
		},
		{
			Name:  "start",
			Usage: "Start virtualmachine",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.StartVirtualMachineParameter{}
				var ids []string
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						ids = append(ids, strings.Split(scanner.Text(), string(sep))[0])
					}
				} else if len(c.Args()) == 1 {
					ids = append(ids, c.Args()[0])
				}
				log.Println("ids:", ids)
				for _, id := range ids {
					params.SetId(id)
					vm, err := client.StartVirtualMachine(params)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								vm.Id.String,
								vm.Name.String,
								vm.Displayname.String,
								vm.State.String,
								vm.Zonename.String,
								vm.Templatename.String,
								vm.Serviceofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:  "stop",
			Usage: "Stop virtualmachine",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.StopVirtualMachineParameter{}
				var ids []string
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						ids = append(ids, strings.Split(scanner.Text(), string(sep))[0])
					}
				} else if len(c.Args()) == 1 {
					ids = append(ids, c.Args()[0])
				}
				log.Println("ids:", ids)
				for _, id := range ids {
					params.SetId(id)
					vm, err := client.StopVirtualMachine(params)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								vm.Id.String,
								vm.Name.String,
								vm.Displayname.String,
								vm.State.String,
								vm.Zonename.String,
								vm.Templatename.String,
								vm.Serviceofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:  "destroy",
			Usage: "Destroy virtualmachine",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.DestroyVirtualMachineParameter{}
				var ids []string
				if len(c.Args()) == 0 {
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						ids = append(ids, strings.Split(scanner.Text(), string(sep))[0])
					}
				} else if len(c.Args()) == 1 {
					ids = append(ids, c.Args()[0])
				}
				log.Println("ids:", ids)
				for _, id := range ids {
					params.SetId(id)
					vm, err := client.DestroyVirtualMachine(params)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								vm.Id.String,
								vm.Name.String,
								vm.Displayname.String,
								vm.State.String,
								vm.Zonename.String,
								vm.Templatename.String,
								vm.Serviceofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:  "zones",
			Usage: "List zones",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.ListZonesParameter{}
				zones, err := client.ListZones(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, zone := range zones {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								zone.Id.String,
								zone.Name.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:      "serviceofferings",
			ShortName: "sizes",
			Usage:     "List serviceofferings",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.ListServiceOfferingsParameter{}
				serviceofferings, err := client.ListServiceOfferings(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, serviceoffering := range serviceofferings {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								serviceoffering.Id.String,
								serviceoffering.Name.String,
								fmt.Sprint(serviceoffering.Cpunumber.Int64),
								fmt.Sprint(serviceoffering.Cpuspeed.Int64),
								fmt.Sprint(serviceoffering.Memory.Int64),
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
				setup(c)
				params := cloudstack.ListTemplatesParameter{}
				params.SetTemplatefilter("featured")
				templates, err := client.ListTemplates(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, template := range templates {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								template.Id.String,
								template.Name.String,
								template.Displaytext.String,
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
				setup(c)
				params := cloudstack.ListNetworksParameter{}
				networks, err := client.ListNetworks(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, network := range networks {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								network.Id.String,
								network.Name.String,
								network.Networkofferingname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
		{
			Name:      "publicipaddresses",
			ShortName: "ips",
			Usage:     "list ipaddresses",
			Action: func(c *cli.Context) {
				setup(c)
				params := cloudstack.ListPublicIpAddressesParameter{}
				ips, err := client.ListPublicIpAddresses(params)
				if err != nil {
					log.Fatal(err)
				}
				for _, ip := range ips {
					fmt.Fprintln(
						tabw,
						strings.Join(
							[]string{
								ip.Id.String,
								ip.Zonename.String,
								ip.Associatednetworkname.String,
								fmt.Sprint(ip.Issourcenat.Bool),
								ip.Ipaddress.String,
								ip.Virtualmachinedisplayname.String,
							}, string(sep)))
				}
				tabw.Flush()
			},
		},
	}

	app.Run(os.Args)
}
