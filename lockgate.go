package lockgate

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/andrew-d/go-termutil"
	"github.com/atsaki/golang-cloudstack-library"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v1"
)

type Account struct {
	URL       string `yaml:"url"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	APIKey    string `yaml:"apikey"`
	SecretKey string `yaml:"secretkey"`
}

type Command struct {
	Options map[string]interface{} `yaml:"options"`
	Keys    []string               `yaml:"keys"`
}

type Config struct {
	Account  Account            `yaml:"account"`
	Commands map[string]Command `yaml:"command"`
}

func convertToArrayOfMap(v interface{}) ([]map[string]interface{}, error) {
	var m []map[string]interface{}
	b, err := json.Marshal(v)
	if err != nil {
		log.Println("Faild to Marshal input")
		return nil, err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Println("Faild to convert input to []map[string]interface{}")
		return nil, err
	}
	return m, nil
}

func SetLogLevel(c *cli.Context) {
	if !c.GlobalBool("debug") {
		log.SetOutput(ioutil.Discard)
	}
}

func GetClient(c *cli.Context) (*cloudstack.Client, error) {

	profile := c.GlobalString("profile")
	config, err := LoadConfig(profile)
	if err != nil {
		msg := "Failed to load config. profile: " + c.GlobalString("profile")
		fmt.Fprintln(os.Stderr, msg)
		log.Println(msg)
		log.Fatal(err)
	}

	url, err := url.Parse(config.Account.URL)
	if err != nil {
		msg := "Failed to parse URL: " + config.Account.URL
		fmt.Fprintln(os.Stderr, msg)
		log.Println(msg)
		log.Fatal(err)
	}

	return cloudstack.NewClient(*url,
		config.Account.APIKey, config.Account.SecretKey,
		config.Account.Username, config.Account.Password)
}

type Writer interface {
	Print(interface{})
}

type TabWriter struct {
	separator byte
	writer    *tabwriter.Writer
	minwidth  int
	tabwidth  int
	padding   int
	padchar   byte
	header    bool
	keys      []string
}

func (tw *TabWriter) Print(xs interface{}) {

	table, err := convertToArrayOfMap(xs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to convert result to map")
		log.Fatal(err)
	}

	keys := tw.keys
	if len(keys) == 0 && len(table) > 0 {
		log.Println("keys is not specified. use keys of first item.", keys)
		for k, _ := range table[0] {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var i int
		var k string
		for i, k = range keys {
			if k == "id" {
				break
			}
		}
		keys = append(append(keys[i:i+1], keys[0:i]...), keys[i+1:]...)
	}
	log.Println("keys:", keys)

	if tw.header {
		fmt.Fprintln(tw.writer, strings.Join(keys, string(tw.separator)))
	}

	for _, m := range table {
		s := ""
		for _, k := range keys {
			switch m[k].(type) {
			case string, float64, bool:
				s += fmt.Sprint(m[k])
			default:
				b, err := json.Marshal(m[k])
				if err != nil {
					log.Println("Faild to Marshal value of", k)
				}
				s += fmt.Sprint(string(b))
			}
			s += string(tw.separator)
		}
		fmt.Fprintln(tw.writer, s)
	}
	tw.writer.Flush()
}

func GetTabWriter(c *cli.Context) *TabWriter {

	profile := c.GlobalString("profile")
	config, err := LoadConfig(profile)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	keys := []string{}
	if c.GlobalString("keys") != "" {
		for _, k := range strings.Split(c.GlobalString("keys"), ",") {
			k = strings.TrimSpace(k)
			if k != "" {
				keys = append(keys, k)
			}
		}
	} else {
		command, ok := config.Commands[c.Command.Name]
		if ok {
			keys = command.Keys
		} else {
			log.Println("Failed to get keys from config: " + c.Command.Name)
		}
	}

	tw := TabWriter{
		separator: '\t',
		writer:    new(tabwriter.Writer),
		minwidth:  0,
		tabwidth:  8,
		padding:   0,
		header:    !c.GlobalBool("no-header"),
		keys:      keys,
	}

	tw.writer.Init(os.Stdout, tw.minwidth, tw.tabwidth, tw.padding,
		tw.separator, 0)
	return &tw
}

func GetConfigFilePath(profile string) string {
	return path.Join(ConfigDir, profile, ConfigFile)
}

func LoadConfig(profile string) (*Config, error) {
	config := new(Config)
	configFilePath := GetConfigFilePath(profile)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = GetConfigFilePath("default")
	}

	contents, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		msg := fmt.Sprintf("Failed to read %s ...", configFilePath)
		fmt.Fprintln(os.Stderr, msg)
		fmt.Fprintln(os.Stderr, err)
		log.Println(msg)
		log.Println(err)
		return nil, err
	}

	err = yaml.Unmarshal(contents, config)
	if err != nil {
		msg := fmt.Sprintf("Failed to unmarshal the contents of %s ...", configFilePath)
		fmt.Fprintln(os.Stderr, msg)
		fmt.Fprintln(os.Stderr, err)
		log.Println(msg)
		log.Println(err)
		return nil, err
	}
	return config, err
}

func GetArgumentsFromStdin() []string {
	var args []string
	if !termutil.Isatty(os.Stdin.Fd()) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			args = append(args, strings.Fields(scanner.Text())...)
		}
	}
	return args
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
