package main

import (
	"os"

	"github.com/atsaki/lockgate/commands"
	"github.com/codegangsta/cli"
)

func main() {

	cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} {{if .Flags}}[global options] {{end}}<command>{{if .Flags}} [command options]{{end}} [arguments...]

COMMANDS:
   {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
GLOBAL OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`

	cli.SubcommandHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.Name}} <command>{{if .Flags}} [command options]{{end}} [arguments...]

COMMANDS:
   {{range .Commands}}{{.ShortName}}{{ "\t" }}{{.Usage}}
   {{end}}{{if .Flags}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
	app := cli.NewApp()
	app.Name = "lg"
	app.Usage = "CLI for CloudStack"
	app.Version = "0.0.1"
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
		commands.API,

		commands.IP,
		commands.Network,
		commands.ServiceOffering,
		commands.Template,
		commands.VM,
		commands.Zone,
	}

	app.Run(os.Args)
}
