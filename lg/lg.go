package main

import (
	"os"

	"github.com/atsaki/lockgate/cli"
	"github.com/atsaki/lockgate/command"
)

var (
	app = cli.Application{
		Name:    "lg",
		Help:    "CLI for CloudStack",
		Version: "0.0.1",
		Flags: []cli.Flag{
			cli.Flag{
				Name:    "profile",
				Short:   'P',
				Default: "default",
				Help:    "Profile to connect CloudStack",
				Type:    cli.String,
			},
			cli.Flag{
				Name:  "no-header",
				Short: 'H',
				Help:  "Show no header line",
				Type:  cli.Bool,
			},
			cli.Flag{
				Name:  "keys",
				Short: 'k',
				Help:  "Keys for output",
				Type:  cli.String,
			},
			cli.Flag{
				Name: "debug",
				Help: "Show log messages for debug",
				Type: cli.Bool,
			},
		},
		Commands: []cli.Command{
			command.Init,
			command.API,
			command.IP,
			command.Network,
			command.ServiceOffering,
			command.Template,
			command.VM,
			command.Zone,
		},
	}
)

func main() {
	app.Run(os.Args[1:])
}
