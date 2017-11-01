package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var args struct {
}

func parseArgs() {
	//flag.StringVar(&args.configFile, "configfile", "", "set config file path")

	flag.Parse()
}

func main() {
	parseArgs()
	args := flag.Args()
	showMsg("arguments: %s", strings.Join(args, " "))

	detectGoWorldPath()

	if len(args) == 0 {
		showMsg("no command to execute")
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\tgoworld <build|start|stop|kill|reload|status> [server-id]\n")
		os.Exit(1)
	}

	cmd := args[0]

	if cmd == "build" || cmd == "start" || cmd == "stop" || cmd == "reload" || cmd == "kill" {
		if len(args) != 2 {
			showMsgAndQuit("server id is not given")
		}
	}

	if cmd == "build" {
		build(ServerID(args[1]))
	} else if cmd == "start" {
		start(ServerID(args[1]))
	} else if cmd == "stop" {
		if IsWindows {
			showMsgAndQuit("stop does not work on Windows, use kill instead (will lose player data)")
		}

		stop(ServerID(args[1]))
	} else if cmd == "reload" {
		reload(ServerID(args[1]))
	} else if cmd == "kill" {
		kill(ServerID(args[1]))
	} else if cmd == "status" {
		status()
	} else {
		showMsgAndQuit("unknown command: %s", cmd)
	}
}
