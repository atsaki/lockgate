package lockgate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/atsaki/lockgate/cli"
)

func getKeys(x interface{}) []string {

	v := reflect.ValueOf(x)
	n := v.Type().NumField()
	keys := make([]string, n, n)

	for i := 0; i < n; i++ {
		var key string
		if v.Type().Field(i).Tag.Get("json") != "" {
			key = v.Type().Field(i).Tag.Get("json")
		} else {
			key = v.Type().Field(i).Name
		}
		keys[i] = key
	}
	sort.Strings(keys)

	return keys
}

func getFieldByTag(v reflect.Value, tag string) reflect.Value {

	if v.Kind() != reflect.Struct {
		log.Fatal("v's Kind must be struct")
	}

	for i := 0; i < v.Type().NumField(); i++ {
		field := v.Type().Field(i)
		if field.Tag.Get("json") == tag {
			return v.FieldByName(field.Name)
		}
	}
	return reflect.Value{}
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

	vs := reflect.ValueOf(xs)

	if vs.Kind() != reflect.Slice {
		log.Println("xs's Kind must be slice")
		return
	}

	if vs.Len() == 0 {
		return
	}

	if vs.Index(0).Kind() != reflect.Struct {
		log.Println("Elements of xs must be struct")
		return
	}

	keys := tw.keys
	if len(keys) == 0 && vs.Len() > 0 {
		log.Println("keys is not specified. use keys of first item.")
		keys = getKeys(vs.Index(0).Interface())
		if len(keys) > 0 {
			// move id column to left side
			var i int
			var k string
			for i, k = range keys {
				if k == "id" {
					break
				}
			}
			keys = append(append(keys[i:i+1], keys[0:i]...), keys[i+1:]...)
		}
	}
	log.Println("keys", keys)

	if tw.header {
		fmt.Fprintln(tw.writer, strings.Join(keys, string(tw.separator)))
	}
	for i := 0; i < vs.Len(); i++ {
		s := ""
		for _, key := range keys {

			field := getFieldByTag(vs.Index(i), key)
			if field.IsValid() {
				// check if filed is Marshaler (Especially cloudstack.Null*).
				p := reflect.New(field.Type())
				p.Elem().Set(field)
				marshaler, ok := p.Interface().(json.Marshaler)
				if ok {
					var v interface{}
					var b []byte
					var err error

					b, err = marshaler.MarshalJSON()
					if err != nil {
						log.Println("Failed to marshal.", marshaler)
					}
					err = json.Unmarshal(b, &v)
					if err != nil {
						log.Println("Failed to unmarshal.", string(b))
					}
					s += fmt.Sprint(v)
				} else {
					s += fmt.Sprint(field.Interface())
				}
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
