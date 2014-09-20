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
			"vm-list": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm-start": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm-stop": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm-deploy": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm-destroy": Command{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"network-list": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"networkofferingname",
				},
			},
			"serviceoffering-list": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"cpunumber",
					"cpuspeed",
					"memory",
				},
			},
			"template-list": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"displaytext",
				},
			},
			"ip-list": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"zonename",
					"issourcenat",
					"ipaddress",
				},
			},
			"zone-list": Command{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
				},
			},
		},
	}
)
