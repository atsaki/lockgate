package main

import (
	"os"

	"github.com/atsaki/lockgate/commands"
	"github.com/codegangsta/cli"
)

func main() {

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
			Usage: "Profile to connect CloudStack service",
		},
		cli.BoolFlag{
			Name:  "no-header, H",
			Usage: "Show no header line",
		},
		cli.StringFlag{
			Name:  "keys, k",
			Value: "",
			Usage: "Keys for output",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Show log messages for debug",
		},
	}

	app.Commands = []cli.Command{
		commands.Init,
		commands.ListVirtualMachines,
		commands.ListZones,
		commands.ListServiceOfferings,
		commands.ListTemplates,
		commands.ListNetworks,
		commands.ListPublicIpAddresses,
		commands.DeployVirtualMachine,
		commands.DestroyVirtualMachine,
		commands.StartVirtualMachine,
		commands.StopVirtualMachine,
	}

	app.Run(os.Args)
}
