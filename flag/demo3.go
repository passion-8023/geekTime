package main

import (
	_ "flag"
	"fmt"
	"net/url"
	_ "os"
)

type URLValue struct {
	URL *url.URL
}

func (U *URLValue) String() string {
	if U.URL != nil {
		return U.URL.String()
	}
	return ""
}

func (U *URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*U.URL = *u
	}
	return nil
}

var u = &url.URL{}

//var name string = "hh"
func main()  {
	var name string
	name = "wx"
	fmt.Println(name)
	if value, ok := interface{}(name).(string); !ok {
		fmt.Println("不是string")
	} else {
		fmt.Println("是string", value)
	}

	//comline := flag.NewFlagSet("测试 flag.Value", flag.ExitOnError)
	//comline.Var(&URLValue{u}, "url", "url to parse")
	//comline.Parse(os.Args[1:])
	//fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
}

