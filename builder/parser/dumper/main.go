package main

import (
	"fmt"
	"os"

	"github.com/Go-Docker-Hackathon/team-ZJU-FT/builder/parser"
)

func main() {
	var f *os.File
	var err error

	if len(os.Args) < 2 {
		fmt.Println("please supply filename(s)")
		os.Exit(1)
	}

	for _, fn := range os.Args[1:] {
		f, err = os.Open(fn)
		if err != nil {
			panic(err)
		}

		ast, err := parser.Parse(f)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(ast.Dump())
		}
	}
}
