package main

import (
	"flag"
	"fmt"
	"jvm/jvmgo/classpath"
	"os"
	"strings"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	xjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xjreOption, "xjre", "", "path to jre")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("usage: %s, [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)

	parse := classpath.Parse(cmd.xjreOption, cmd.cpOption)
	class, entry, err := parse.ReadClass(strings.Replace(cmd.class, ".", "/", -1))

	fmt.Print(class)
	fmt.Print(entry)
	fmt.Print(err)
}
