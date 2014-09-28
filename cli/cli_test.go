package cli

import (
	"fmt"
	"testing"

	"github.com/atsaki/lockgate/cli"
)

func TestApp(t *testing.T) {
	app := cli.Application{
		Name:    "test",
		Help:    "TestApp",
		Version: "0.1",
		Args: []cli.Argument{
			cli.Argument{
				Name: "arg1",
				Help: "arg1 help",
				Type: cli.String,
			},
		},
		Flags: []cli.Flag{
			cli.Flag{
				Name:  "flag1",
				Short: 'F',
				Help:  "flag1 help",
				Type:  cli.String,
			},
		},
	}
	app.Run([]string{})
}

func TestCommand(t *testing.T) {
	app := cli.Application{
		Name:    "test",
		Help:    "TestApp",
		Version: "0.1",
		Commands: []cli.Command{
			cli.Command{
				Name: "hello",
				Help: "cmd help",
				Flags: []cli.Flag{
					cli.Flag{
						Name:    "name",
						Short:   'N',
						Default: "Alice",
						Help:    "name",
						Type:    cli.String,
					},
				},
				Action: func(c *cli.Context) {
					name := c.Command.Flag("name").String()
					fmt.Println("Hello", name)
				},
			},
		},
	}
	app.Run([]string{"hello", "--name", "Bob"})
}
