package main

import (
	"fmt"
	"github.com/rikuayanokozy/options"
)

func main() {
	v := struct {
                Foo string `env:"FOO" flag:"foo"`
                Bar bool   `env:"BAR" flag:"bar"`
        }{}
	options.Parse(&v, true, true)
	fmt.Println("Foo:", v.Foo)
	fmt.Println("Bar:", v.Bar)
}
