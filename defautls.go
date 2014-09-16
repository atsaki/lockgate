package lockgate

import (
	"path"

	"github.com/atsaki/lockgate/util"
)

var (
	ConfigDir     = util.ExpandPath(path.Join("~", ".lg"))
	ConfigFile    = "config.toml"
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
				Keys: []string{
					"id",
					"name",
					"displayname",
					"state",
					"zonename",
					"templatename",
					"serviceofferingname",
				},
				Args: map[string]interface{}{},
			},
			"zones": Command{
				Keys: []string{
					"id",
					"name",
				},
				Args: map[string]interface{}{},
			},
		},
	}
)
