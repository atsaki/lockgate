package lockgate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/atsaki/golang-cloudstack-library"
	"github.com/atsaki/lockgate/cli"
)

func convertToArrayOfMap(xs []interface{}) ([]map[string]interface{}, error) {
	ys := make([]map[string]interface{}, len(xs))

	for i, x := range xs {
		var (
			b   []byte
			err error
			m   map[string]interface{}
		)

		switch x.(type) {
		case cloudstack.Firewallrule:
			v := x.(cloudstack.Firewallrule)
			b, err = json.Marshal(&v)
		case cloudstack.Network:
			v := x.(cloudstack.Network)
			b, err = json.Marshal(&v)
		case cloudstack.Portforwardingrule:
			v := x.(cloudstack.Portforwardingrule)
			b, err = json.Marshal(&v)
		case cloudstack.Serviceoffering:
			v := x.(cloudstack.Serviceoffering)
			b, err = json.Marshal(&v)
		case cloudstack.Template:
			v := x.(cloudstack.Template)
			b, err = json.Marshal(&v)
		case cloudstack.Virtualmachine:
			v := x.(cloudstack.Virtualmachine)
			b, err = json.Marshal(&v)
		case cloudstack.Publicipaddress:
			v := x.(cloudstack.Publicipaddress)
			b, err = json.Marshal(&v)
		case cloudstack.Sshkeypair:
			v := x.(cloudstack.Sshkeypair)
			b, err = json.Marshal(&v)
		case cloudstack.Zone:
			v := x.(cloudstack.Zone)
			b, err = json.Marshal(&v)
		default:
			b, err = json.Marshal(&x)
		}

		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = json.Unmarshal(b, &m)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ys[i] = m
	}
	return ys, nil
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

func (tw *TabWriter) Print(xs []interface{}) {

	table, err := convertToArrayOfMap(xs)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to convert result to map")
		log.Fatal(err)
	}

	keys := tw.keys
	if len(keys) == 0 && len(table) > 0 {
		log.Println("keys is not specified. use keys of first item.")
		log.Println("first item:", table[0])
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
		if len(keys) > 0 {
			// move id column to left side
			keys = append(append(keys[i:i+1], keys[0:i]...), keys[i+1:]...)
		}
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

	profile := GetProfile(c)
	config, err := LoadConfig(profile)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	keys := []string{}
	if c.App.Flag("keys").String() != "" {
		keys = strings.Split(c.App.Flag("keys").String(), ",")
		for i := range keys {
			keys[i] = strings.TrimSpace(keys[i])
		}
	} else {
		command, ok := config.Commands[c.CommandName]
		if ok {
			keys = command.Keys
		} else {
			log.Println("Unable to get keys from config: " + c.Command.Name)
		}
	}

	tw := TabWriter{
		separator: '\t',
		writer:    new(tabwriter.Writer),
		minwidth:  0,
		tabwidth:  8,
		padding:   0,
		header:    !c.App.Flag("no-header").Bool(),
		keys:      keys,
	}

	tw.writer.Init(os.Stdout, tw.minwidth, tw.tabwidth, tw.padding,
		tw.separator, 0)
	return &tw
}
