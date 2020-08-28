package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var ver = flag.Bool("version", false, "Show version")
	var help = flag.Bool("help", false, "Show help")

	flag.Usage = showUsage
	flag.Parse()

	if *help {
		showUsage()
		os.Exit(0)
	}

	if *ver {
		showVersion()
		os.Exit(0)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(UserAgents[rand.Intn(len(UserAgents))])
}
