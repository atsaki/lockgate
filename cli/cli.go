package cli

import (
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/alecthomas/units"
	"gopkg.in/alecthomas/kingpin.v1"
)

const (
	Bool = iota
	Bytes
	Duration
	Enum
	ExistingDir
	ExistingFile
	File
	Float
	IP
	Int
	Int64
	OpenFile
	String
	StringMap
	Strings
	TCP
	TCPList
	URL
	URLList
	Uint64
)

type Argument struct {
	Name         string
	Help         string
	Type         int
	Required     bool
	valuePointer unsafe.Pointer
}

func (arg *Argument) init(karg *kingpin.ArgClause) {

	if arg.Required {
		karg.Required()
	}

	switch arg.Type {
	case Bool:
		arg.valuePointer = unsafe.Pointer(karg.Bool())
	case Bytes:
		arg.valuePointer = unsafe.Pointer(karg.Bytes())
	case Duration:
		arg.valuePointer = unsafe.Pointer(karg.Duration())
	case Enum:
		arg.valuePointer = unsafe.Pointer(karg.Enum())
	case ExistingDir:
		arg.valuePointer = unsafe.Pointer(karg.ExistingDir())
	case ExistingFile:
		arg.valuePointer = unsafe.Pointer(karg.ExistingFile())
	case File:
		arg.valuePointer = unsafe.Pointer(karg.File())
	case Float:
		arg.valuePointer = unsafe.Pointer(karg.Float())
	case IP:
		arg.valuePointer = unsafe.Pointer(karg.IP())
	case Int:
		arg.valuePointer = unsafe.Pointer(karg.Int())
	case Int64:
		arg.valuePointer = unsafe.Pointer(karg.Int64())
	case OpenFile:
		/* arg.valuePointer = unsafe.Pointer(karg.OpenFile()) */
	case String:
		arg.valuePointer = unsafe.Pointer(karg.String())
	case StringMap:
		arg.valuePointer = unsafe.Pointer(karg.StringMap())
	case Strings:
		arg.valuePointer = unsafe.Pointer(karg.Strings())
	case TCP:
		arg.valuePointer = unsafe.Pointer(karg.TCP())
	case TCPList:
		arg.valuePointer = unsafe.Pointer(karg.TCPList())
	case URL:
		arg.valuePointer = unsafe.Pointer(karg.URL())
	case URLList:
		arg.valuePointer = unsafe.Pointer(karg.URLList())
	case Uint64:
		arg.valuePointer = unsafe.Pointer(karg.Uint64())
	}
}

func (arg *Argument) initApplicationArg(kapp *kingpin.Application) {
	arg.init(kapp.Arg(arg.Name, arg.Help))
}

func (arg *Argument) initCommandArg(kcmd *kingpin.CmdClause) {
	arg.init(kcmd.Arg(arg.Name, arg.Help))
}

func (arg *Argument) Bool() bool {
	return *(*bool)(arg.valuePointer)
}

func (arg *Argument) Bytes() units.Base2Bytes {
	return *(*units.Base2Bytes)(arg.valuePointer)
}

func (arg *Argument) Duration() time.Duration {
	return *(*time.Duration)(arg.valuePointer)
}

func (arg *Argument) Enum() string {
	return *(*string)(arg.valuePointer)
}

func (arg *Argument) ExistingDir() string {
	return *(*string)(arg.valuePointer)
}

func (arg *Argument) ExistingFile() string {
	return *(*string)(arg.valuePointer)
}

func (arg *Argument) File() *os.File {
	return *(**os.File)(arg.valuePointer)
}

func (arg *Argument) Float() float64 {
	return *(*float64)(arg.valuePointer)
}

func (arg *Argument) IP() net.IP {
	return *(*net.IP)(arg.valuePointer)
}

func (arg *Argument) Int() int {
	return *(*int)(arg.valuePointer)
}

func (arg *Argument) Int64() int64 {
	return *(*int64)(arg.valuePointer)
}

func (arg *Argument) OpenFile() *os.File {
	return *(**os.File)(arg.valuePointer)
}

func (arg *Argument) String() string {
	return *(*string)(arg.valuePointer)
}

func (arg *Argument) StringMap() map[string]string {
	return *(*map[string]string)(arg.valuePointer)
}

func (arg *Argument) Strings() []string {
	return *(*[]string)(arg.valuePointer)
}

func (arg *Argument) TCP() *net.TCPAddr {
	return *(**net.TCPAddr)(arg.valuePointer)
}

func (arg *Argument) TCPList() []*net.TCPAddr {
	return *(*[]*net.TCPAddr)(arg.valuePointer)
}

func (arg *Argument) URL() *url.URL {
	return *(**url.URL)(arg.valuePointer)
}

func (arg *Argument) URLList() []*url.URL {
	return *(*[]*url.URL)(arg.valuePointer)
}

func (arg *Argument) Uint64() uint64 {
	return *(*uint64)(arg.valuePointer)
}

type Flag struct {
	Name         string
	Short        byte
	Help         string
	Type         int
	Default      string
	Required     bool
	valuePointer unsafe.Pointer
}

func (flag *Flag) init(kflag *kingpin.FlagClause) {

	kflag.Short(flag.Short)

	if flag.Default != "" {
		kflag.Default(flag.Default)
	}

	if flag.Required {
		kflag.Required()
	}

	switch flag.Type {
	case Bool:
		flag.valuePointer = unsafe.Pointer(kflag.Bool())
	case Bytes:
		flag.valuePointer = unsafe.Pointer(kflag.Bytes())
	case Duration:
		flag.valuePointer = unsafe.Pointer(kflag.Duration())
	case Enum:
		flag.valuePointer = unsafe.Pointer(kflag.Enum())
	case ExistingDir:
		flag.valuePointer = unsafe.Pointer(kflag.ExistingDir())
	case ExistingFile:
		flag.valuePointer = unsafe.Pointer(kflag.ExistingFile())
	case File:
		flag.valuePointer = unsafe.Pointer(kflag.File())
	case Float:
		flag.valuePointer = unsafe.Pointer(kflag.Float())
	case IP:
		flag.valuePointer = unsafe.Pointer(kflag.IP())
	case Int:
		flag.valuePointer = unsafe.Pointer(kflag.Int())
	case Int64:
		flag.valuePointer = unsafe.Pointer(kflag.Int64())
	case OpenFile:
		/* flag.valuePointer = unsafe.Pointer(kflag.OpenFile()) */
	case String:
		flag.valuePointer = unsafe.Pointer(kflag.String())
	case StringMap:
		flag.valuePointer = unsafe.Pointer(kflag.StringMap())
	case Strings:
		flag.valuePointer = unsafe.Pointer(kflag.Strings())
	case TCP:
		flag.valuePointer = unsafe.Pointer(kflag.TCP())
	case TCPList:
		flag.valuePointer = unsafe.Pointer(kflag.TCPList())
	case URL:
		flag.valuePointer = unsafe.Pointer(kflag.URL())
	case URLList:
		flag.valuePointer = unsafe.Pointer(kflag.URLList())
	case Uint64:
		flag.valuePointer = unsafe.Pointer(kflag.Uint64())
	}
}

func (flag *Flag) initApplicationFlag(kapp *kingpin.Application) {
	flag.init(kapp.Flag(flag.Name, flag.Help))
}

func (flag *Flag) initCommandFlag(kcmd *kingpin.CmdClause) {
	flag.init(kcmd.Flag(flag.Name, flag.Help))
}

func (flag *Flag) Bool() bool {
	return *(*bool)(flag.valuePointer)
}

func (flag *Flag) Bytes() units.Base2Bytes {
	return *(*units.Base2Bytes)(flag.valuePointer)
}

func (flag *Flag) Duration() time.Duration {
	return *(*time.Duration)(flag.valuePointer)
}

func (flag *Flag) Enum() string {
	return *(*string)(flag.valuePointer)
}

func (flag *Flag) ExistingDir() string {
	return *(*string)(flag.valuePointer)
}

func (flag *Flag) ExistingFile() string {
	return *(*string)(flag.valuePointer)
}

func (flag *Flag) File() *os.File {
	return *(**os.File)(flag.valuePointer)
}

func (flag *Flag) Float() float64 {
	return *(*float64)(flag.valuePointer)
}

func (flag *Flag) IP() net.IP {
	return *(*net.IP)(flag.valuePointer)
}

func (flag *Flag) Int() int {
	return *(*int)(flag.valuePointer)
}

func (flag *Flag) Int64() int64 {
	return *(*int64)(flag.valuePointer)
}

func (flag *Flag) OpenFile() *os.File {
	return *(**os.File)(flag.valuePointer)
}

func (flag *Flag) String() string {
	return *(*string)(flag.valuePointer)
}

func (flag *Flag) StringMap() map[string]string {
	return *(*map[string]string)(flag.valuePointer)
}

func (flag *Flag) Strings() []string {
	return *(*[]string)(flag.valuePointer)
}

func (flag *Flag) TCP() *net.TCPAddr {
	return *(**net.TCPAddr)(flag.valuePointer)
}

func (flag *Flag) TCPList() []*net.TCPAddr {
	return *(*[]*net.TCPAddr)(flag.valuePointer)
}

func (flag *Flag) URL() *url.URL {
	return *(**url.URL)(flag.valuePointer)
}

func (flag *Flag) URLList() []*url.URL {
	return *(*[]*url.URL)(flag.valuePointer)
}

func (flag *Flag) Uint64() uint64 {
	return *(*uint64)(flag.valuePointer)
}

type Command struct {
	Name     string
	Help     string
	Commands []Command
	Args     []Argument
	Flags    []Flag
	Action   func(*Context)
	kcmd     *kingpin.CmdClause
}

func (cmd *Command) init() {

	for i := range cmd.Args {
		flag := &cmd.Args[i]
		flag.initCommandArg(cmd.kcmd)
	}

	for i := range cmd.Flags {
		flag := &cmd.Flags[i]
		flag.initCommandFlag(cmd.kcmd)
	}

	for i := range cmd.Commands {
		c := &cmd.Commands[i]
		c.kcmd = cmd.kcmd.Command(c.Name, c.Help)
		c.init()
	}
}

func (cmd *Command) Arg(arg string) *Argument {
	for _, a := range cmd.Args {
		if arg == a.Name {
			return &a
		}
	}
	return nil
}

func (cmd *Command) Flag(flag string) *Flag {
	for _, f := range cmd.Flags {
		if flag == f.Name {
			return &f
		}
	}
	return nil
}

func (cmd *Command) Command(names []string) *Command {
	if len(names) == 0 {
		return nil
	}
	for _, c := range cmd.Commands {
		if names[0] == c.Name {
			if len(names) == 1 {
				return &c
			} else {
				return c.Command(names[1:])
			}
		}
	}
	return nil
}

type Application struct {
	Name     string
	Help     string
	Version  string
	Args     []Argument
	Flags    []Flag
	Commands []Command
	kapp     *kingpin.Application
}

func (app *Application) init() {

	app.kapp = kingpin.New(app.Name, app.Help)
	kingpin.Version(app.Version)

	for i := range app.Args {
		flag := &app.Args[i]
		flag.initApplicationArg(app.kapp)
	}

	for i := range app.Flags {
		flag := &app.Flags[i]
		flag.initApplicationFlag(app.kapp)
	}

	for i := range app.Commands {
		cmd := &app.Commands[i]
		cmd.kcmd = app.kapp.Command(cmd.Name, cmd.Help)
		cmd.init()
	}
}

func (app *Application) Arg(arg string) *Argument {
	for _, a := range app.Args {
		if arg == a.Name {
			return &a
		}
	}
	return nil
}

func (app *Application) Flag(flag string) *Flag {
	for _, f := range app.Flags {
		if flag == f.Name {
			return &f
		}
	}
	return nil
}

func (app *Application) Command(names []string) *Command {
	if len(names) == 0 {
		return nil
	}
	for _, c := range app.Commands {
		if names[0] == c.Name {
			if len(names) == 1 {
				return &c
			} else {
				return c.Command(names[1:])
			}
		}
	}
	return nil
}

func (app *Application) Usage() {
	if app.kapp != nil {
		app.kapp.Usage(os.Stderr)
	}
}

func (app *Application) Parse(args []string) (*Context, error) {

	context := &Context{
		App: app,
	}

	app.init()

	cmdname, err := app.kapp.Parse(args)
	if err != nil {
		return context, err
	}

	context.CommandName = cmdname
	context.Command = app.Command(strings.Fields(cmdname))
	return context, nil
}

func (app *Application) Run(args []string) {

	context, err := app.Parse(args)
	if err != nil {
		if len(args) == 0 {
			context.App.Usage()
			return
		}
		log.Fatal(err)
	}
	if context.Command != nil && context.Command.Action != nil {
		context.Action()
	}
}

type Context struct {
	App         *Application
	CommandName string
	Command     *Command
}

func (context *Context) Action() {
	context.Command.Action(context)
}
