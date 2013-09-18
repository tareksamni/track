package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/simonz05/track/server"
	"github.com/simonz05/track/util"
)

var (
	verbose    = flag.Bool("v", false, "verbose mode")
	help       = flag.Bool("h", false, "show help text")
	laddr      = flag.String("http", ":8080", "set bind address for the HTTP server")
	dsn        = flag.String("dsn", "testing:testing@tcp(localhost:3306)/testing?utf8", "MySQL Data Source Name")
	logLevel   = flag.Int("log", 0, "set log level")
	version    = flag.Bool("version", false, "show version number and exit")
	cpuprofile = flag.String("debug.cpuprofile", "", "write cpu profile to file")
)

var Version   = "0.1.1"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Fprintln(os.Stdout, Version)
		return
	}

	if *help {
		flag.Usage()
		os.Exit(1)
	}

	if *laddr == "" {
		fmt.Fprintln(os.Stderr, "listen address required")
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	if *verbose {
		util.LogLevel = 2
	} else {
		util.LogLevel = *logLevel
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	err := server.ListenAndServe(*laddr, *dsn)

	if err != nil {
		log.Println(err)
	}
}
