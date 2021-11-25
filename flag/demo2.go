package main

import (
	"flag"
	"fmt"
	"os"
)

var name2 string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init()  {
	cmdLine.StringVar(&name2, "name", "flag", "this flag package")
}

func main()  {
	for i, arg := range os.Args {
		fmt.Printf("第%d个参数是%q\n", i, arg)
	}


	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name2)
}
