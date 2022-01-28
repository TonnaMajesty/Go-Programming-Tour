package main

import (
	"flag"
	"log"
)

var name string

func flag_sub() {
	flag.Parse()

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	// go run main.go go -name=eddycjy
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	// go run main.go php -n=tonna
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}