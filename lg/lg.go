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
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
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

		commands.VMList,
		commands.VMStart,
		commands.VMStop,
		commands.VMDeploy,
		commands.VMDestroy,

		commands.NetworkList,
		commands.IPList,
		commands.ServiceOfferingList,
		commands.TemplateList,
		commands.ZoneList,
	}

	app.Run(os.Args)
}
