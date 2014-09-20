package lockgate

import (
	"path"

	"github.com/atsaki/lockgate/util"
)

var (
	ConfigDir          = util.ExpandPath(path.Join("~", ".lg"))
	ConfigFile         = "config.yaml"
	VirtualMachineKeys = []string{
		"id",
		"name",
		"displayname",
		"state",
		"zonename",
		"templatename",
		"serviceofferingname",
	}
	DefaultConfig = Config{
		Account: Account{
			URL:       "http://localhost:8080/client/api",
			Username:  "admin",
			Password:  "password",
			APIKey:    "",
			SecretKey: "",
		},
		Commands: map[string]Command{
			"virtualmachines": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"start": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"stop": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"deploy": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"destroy": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"networks": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"networkofferingname",
				},
			},
			"serviceofferings": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"cpunumber",
					"cpuspeed",
					"memory",
				},
			},
			"publicipaddresses": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"zonename",
					"issourcenat",
					"ipaddress",
				},
			},
			"zones": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
				},
			},
		},
	}
)
