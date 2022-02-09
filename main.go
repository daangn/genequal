package main

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

//go:embed genequal.tmpl
var tmpl string

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("no path provided")
	}

	filename := args[0]
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	source := string(buf)

	c := NewCollector()
	decls, err := c.Collect(source)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	t := template.New("equal method")
	t, err = t.Parse(tmpl)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	err = t.Execute(os.Stdout, decls)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
