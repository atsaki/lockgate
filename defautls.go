package lockgate

import (
	"path"
)

var (
	ConfigDir          = expandPath(path.Join("~", ".lg"))
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
		Account: AccountConfig{
			URL:       "http://localhost:8080/client/api",
			Username:  "admin",
			Password:  "password",
			APIKey:    "",
			SecretKey: "",
		},
		Commands: map[string]CommandConfig{
			"vm list": CommandConfig{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm start": CommandConfig{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm stop": CommandConfig{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm deploy": CommandConfig{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"vm destroy": CommandConfig{
				Options: map[string]interface{}{},
				Keys:    VirtualMachineKeys,
			},
			"network list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"networkofferingname",
				},
			},
			"nic list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"isdefault",
					"ipaddress",
					"netmask",
				},
			},
			"firewallrule list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"cidrlist",
					"protocol",
					"startport",
					"endport",
					"icmpcode",
					"icmptype",
				},
			},
			"portforwardingrule list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"ipaddress",
					"virtualmachinename",
					"protocol",
					"publicport",
					"publicendport",
					"privateport",
					"privateendport",
				},
			},
			"serviceoffering list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"cpunumber",
					"cpuspeed",
					"memory",
				},
			},
			"template list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
					"displaytext",
				},
			},
			"ip list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"zonename",
					"issourcenat",
					"ipaddress",
				},
			},
			"zone list": CommandConfig{
				Options: map[string]interface{}{},
				Keys: []string{
					"id",
					"name",
				},
			},
		},
	}
)
