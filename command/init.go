package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/atsaki/lockgate"
	"github.com/atsaki/lockgate/cli"
	"gopkg.in/yaml.v1"
)

var (
	Init = cli.Command{
		Name: "init",
		Help: "Create profile configuration file",
		Args: []cli.Argument{
			cli.Argument{
				Name: "profile",
				Help: "Profile name",
				Type: cli.String,
			},
		},
		Action: func(c *cli.Context) {

			lockgate.SetLogLevel(c)

			newProfile := "default"
			if c.Command.Arg("profile") != nil {
				newProfile = c.Command.Arg("profile").Value().(string)
			}

			if _, err := os.Stat(lockgate.ConfigDir); os.IsNotExist(err) {
				msg := fmt.Sprintf("Createing %s ...", lockgate.ConfigDir)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)
				err := os.Mkdir(lockgate.ConfigDir, 0755)
				if err != nil {
					msg := "Failed to create " + lockgate.ConfigDir
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Fatal(err)
				}
			}

			configDirPath := path.Join(lockgate.ConfigDir, newProfile)
			configFilePath := path.Join(configDirPath, lockgate.ConfigFile)

			if _, err := os.Stat(configDirPath); os.IsNotExist(err) {
				msg := fmt.Sprintf("Createing %s ...", configDirPath)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)
				err := os.Mkdir(configDirPath, 0755)
				if err != nil {
					msg := "Failed to create " + configDirPath
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Fatal(err)
				}
			}

			if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
				msg := fmt.Sprintf("Writing %s ...", configFilePath)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)

				b, err := yaml.Marshal(lockgate.DefaultConfig)
				if err != nil {
					msg := "Failed to marshal default config"
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Fatal(err)
				}

				err = ioutil.WriteFile(configFilePath, b, 0644)
				if err != nil {
					msg := "Failed to write config: " + configFilePath
					fmt.Fprintln(os.Stderr, msg)
					fmt.Fprintln(os.Stderr, err)
					log.Println(msg)
					log.Fatal(err)
				}
			} else {
				msg := fmt.Sprintf("profile %s already exists.", newProfile)
				fmt.Fprintln(os.Stderr, msg)
				log.Println(msg)
			}
		},
	}
)
