package lockgate

import (
	"path"

	"github.com/atsaki/lockgate/util"
)

var (
	GlobalConfDir       = util.ExpandPath(path.Join("~", ".lg"))
	GlobalConfFile      = path.Join(GlobalConfDir, "config")
	DefaultGlobalConfig = GlobalConfig{
		Option: GlobalOption{
			Profile: "default",
		},
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
		},
	}
)
