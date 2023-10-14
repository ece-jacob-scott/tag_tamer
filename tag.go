package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

type Flags struct {
	filePath string
	query    string
	text     bool
}

func (f *Flags) Parse() {
	flag.StringVar(&f.filePath, "file", "", "path to an html file to query")
	flag.StringVar(&f.query, "query", "", "css selector query (required)")
	flag.BoolVar(&f.text, "text", false, "print text instead of html element")
	flag.Parse()
}

func main() {
	flags := Flags{}
	flags.Parse()

	if flags.query == "" {
		flag.Usage()
		panic("--query is required")
	}

	// read from stdin if no file path is provided
	var input *os.File
	if flags.filePath == "" {
		input = os.Stdin
	} else {
		var err error
		input, err = os.Open(flags.filePath)
		if err != nil {
			panic(err)
		}
	}

	doc, err := html.Parse(input)
	if err != nil {
		panic(err)
	}

	query := cascadia.MustCompile(flags.query)

	nodes := cascadia.QueryAll(doc, query)

	for _, node := range nodes {
		if flags.text {
			fmt.Printf("%s\n", node.FirstChild.Data)
		} else {
			// turn render into a string
			var b bytes.Buffer
			_ = html.Render(&b, node)
			fmt.Printf("%s\n", b.String())
		}

	}
}
