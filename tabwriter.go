package lockgate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/atsaki/lockgate/cli"
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

	profile := GetProfile(c)
	config, err := LoadConfig(profile)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	keys := []string{}
	if c.App.Flag("keys").Value().(string) != "" {
		keys = strings.Split(c.App.Flag("keys").Value().(string), ",")
		for i := range keys {
			keys[i] = strings.TrimSpace(keys[i])
		}
	} else {
		command, ok := config.Commands[c.CommandName]
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
		header:    !c.App.Flag("no-header").Value().(bool),
		keys:      keys,
	}

	tw.writer.Init(os.Stdout, tw.minwidth, tw.tabwidth, tw.padding,
		tw.separator, 0)
	return &tw
}
