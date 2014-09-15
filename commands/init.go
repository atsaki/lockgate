package commands

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/util"
	"github.com/codegangsta/cli"
)

var (
	Init = cli.Command{
		Name:  "init",
		Usage: "Create configuration files",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "import-from-cloudmonkey, m",
				Value: "~/.cloudmonkey/config",
				Usage: "import profiles from cloudmonkey config",
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			profile := "local"
			if len(c.Args()) > 0 {
				profile = c.Args()[0]
			}

			lgroot := util.ExpandPath(path.Join("~", ".lg"))
			lgconf := path.Join(lgroot, "config")
			profroot := path.Join(lgroot, profile)
			profconf := path.Join(profroot, "config")

			var globalConfig lockgate.GlobalConfig

			if _, err := os.Stat(lgroot); os.IsNotExist(err) {
				msg := fmt.Sprintf("Createing %s ...", lgroot)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)
				err := os.Mkdir(lgroot, 0755)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					log.Println(err)
				}
			}

			if _, err := os.Stat(lgconf); os.IsNotExist(err) {
				msg := fmt.Sprintf("Writing %s ...", lgconf)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)

				f, err := os.Create(lgconf)
				if err != nil {
					msg := "Failed to create config: " + lgconf
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Println(err)
				}

				globalConfig := lockgate.DefaultGlobalConfig
				encoder := toml.NewEncoder(f)
				err = encoder.Encode(globalConfig)
				if err != nil {
					msg := "Failed to write config: " + lgconf
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Println(err)
				}
			}

			_, err := toml.DecodeFile(lgconf, &globalConfig)
			if err != nil {
				msg := fmt.Sprintf("Failed to decode %s ...", lgconf)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)

			}

			if _, err := os.Stat(profroot); os.IsNotExist(err) {
				msg := fmt.Sprintf("Createing %s ...", profroot)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)
				err := os.Mkdir(profroot, 0755)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					log.Println(err)
				}
			}

			if _, err := os.Stat(profconf); os.IsNotExist(err) {
				msg := fmt.Sprintf("Writing %s ...", profconf)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)

				f, err := os.Create(profconf)
				if err != nil {
					msg := "Failed to create config: " + profconf
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Println(err)
				}

				globalConfig := lockgate.ProfileConfig{}
				encoder := toml.NewEncoder(f)
				err = encoder.Encode(globalConfig)
				if err != nil {
					msg := "Failed to write config: " + profconf
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Println(err)
				}
			}
		},
	}
)
