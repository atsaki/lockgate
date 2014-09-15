package lockgate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate/util"
	"github.com/codegangsta/cli"
	"github.com/vaughan0/go-ini"
)

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

	configfile := util.ExpandPath(c.GlobalString("config-file"))
	log.Println("configfile:", configfile)
	cfg, err := ini.LoadFile(configfile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load config:", configfile)
		log.Fatal(err)
	}
	var ok bool

	profile, ok := cfg.Get("core", "profile")
	if !ok {
		log.Println("profile is not specified. use local profile.")
		profile = "local"
	}
	if c.GlobalString("profile") != "" {
		profile = c.GlobalString("profile")
	}
	log.Println("profile:", profile)

	endpointUrl, ok := cfg.Get(profile, "url")
	if !ok {
		msg := fmt.Sprintln("url is missing in config:", configfile, ",",
			"profile:", profile)
		fmt.Fprint(os.Stderr, msg)
		log.Fatal(msg)
	}
	log.Println("url:", endpointUrl)

	endpoint, err := url.Parse(endpointUrl)
	if err != nil {
		msg := fmt.Sprintln("Failed to parse endpoint URL")
		fmt.Fprintln(os.Stderr, msg)
		log.Println(msg)
		log.Fatal(err)
	}
	log.Println("endpoint:", endpoint)

	re := regexp.MustCompile(".")
	apikey, ok := cfg.Get(profile, "apikey")
	if !ok {
		apikey = ""
	}
	log.Println("apikey:", apikey)

	secretkey, ok := cfg.Get(profile, "secretkey")
	if !ok {
		secretkey = ""
		log.Println("secretkey:", "")
	}
	log.Println("secretkey:", re.ReplaceAllString(secretkey, "*"))

	username, ok := cfg.Get(profile, "username")
	if !ok {
		username = ""
	}
	log.Println("username:", username)

	password, ok := cfg.Get(profile, "password")
	if !ok {
		password = ""
	}
	log.Println("password:", re.ReplaceAllString(password, "*"))

	if (apikey == "" || secretkey == "") && username == "" {
		msg := fmt.Sprintln("apikey/secretkey or usename/password must be specified.")
		fmt.Fprint(os.Stderr, msg)
		log.Fatal(msg)
	}

	return cloudstack.NewClient(*endpoint, apikey, secretkey,
		username, password)
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

	keys := []string{}
	for _, k := range strings.Split(c.GlobalString("keys"), ",") {
		k = strings.TrimSpace(k)
		if k != "" {
			keys = append(keys, k)
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
