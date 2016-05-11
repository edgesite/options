package main

import (
	"fmt"

	"github.com/rikuayanokozy/options"
)

func main() {
	v := struct {
		Foo    string `env:"Foo" flag:"foo"`
		Bar    bool   `options:"auto"`
		FooBar bool   `options:"auto"`
	}{}
	options.Parse(&v, true, true)
	fmt.Println("Foo:", v.Foo)
	fmt.Println("Bar:", v.Bar)
	fmt.Println("FooBar:", v.FooBar)
}
