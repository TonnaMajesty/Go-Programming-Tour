package cmd

import (
	"flag"
	"log"
)

func flag_t() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}