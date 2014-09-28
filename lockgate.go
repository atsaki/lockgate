package lockgate

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/andrew-d/go-termutil"
	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate/cli"
	"gopkg.in/yaml.v1"
)

type AccountConfig struct {
	URL       string `yaml:"url"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	APIKey    string `yaml:"apikey"`
	SecretKey string `yaml:"secretkey"`
}

type CommandConfig struct {
	Options map[string]interface{} `yaml:"options"`
	Keys    []string               `yaml:"keys"`
}

type Config struct {
	Account  AccountConfig            `yaml:"account"`
	Commands map[string]CommandConfig `yaml:"command"`
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

func SetLogLevel(c *cli.Context) {
	if !c.App.Flag("debug").Bool() {
		log.SetOutput(ioutil.Discard)
	}
}

func GetProfile(c *cli.Context) string {
	profile := c.App.Flag("profile").String()
	if profile == "" {
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			if pair[0] == "LGPROF" {
				profile = pair[1]
			}
		}
		if profile == "" {
			profile = "default"
		}
	}
	return profile
}

func GetClient(c *cli.Context) (*cloudstack.Client, error) {

	profile := GetProfile(c)

	config, err := LoadConfig(profile)
	if err != nil {
		msg := "Failed to load config. profile: " + profile
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

func GetArgumentsFromStdin() []string {
	var args []string
	if !termutil.Isatty(os.Stdin.Fd()) {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fields := strings.Fields(scanner.Text())
			if len(fields) > 0 {
				args = append(args, fields[0])
			}
		}
	}
	return args
}
