package main

import (
	"flag"
	"fmt"
	"os"
)

var name1 string

func init()  {

	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name1, "name", "everyone", "hahahha")
}

func main()  {
	flag.Parse()
	fmt.Printf("Hello %s!\n", name1)
}
