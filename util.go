package lockgate

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func expandPath(path string) string {
	usr, _ := user.Current()
	home := usr.HomeDir

	if strings.HasPrefix(path, "~/") {
		path = strings.Replace(path, "~/", home+"/", 1)
	}
	return path
}

func PrettyPrint(b []byte) {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		fmt.Println(string(b))
		return
	}

	out, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to marshal.")
		fmt.Println(string(b))
		return
	}

	fmt.Println(string(out))
}
