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
<<<<<<< HEAD
	filePath string
	query    string
	text     bool
=======
	filePath  string
	query     string
	text      bool
	attribute string
>>>>>>> 82a0501 (init commit)
}

func (f *Flags) Parse() {
	flag.StringVar(&f.filePath, "file", "", "path to an html file to query")
	flag.StringVar(&f.query, "query", "", "css selector query (required)")
	flag.BoolVar(&f.text, "text", false, "print text instead of html element")
<<<<<<< HEAD
=======
	flag.StringVar(&f.attribute, "attr", "", "print attribute instead of html element")
>>>>>>> 82a0501 (init commit)
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
<<<<<<< HEAD
		if flags.text {
=======
		if flags.attribute != "" {
			for _, attr := range node.Attr {
				if attr.Key == flags.attribute {
					fmt.Printf("%s\n", attr.Val)
				}
			}
		} else if flags.text {
>>>>>>> 82a0501 (init commit)
			fmt.Printf("%s\n", node.FirstChild.Data)
		} else {
			// turn render into a string
			var b bytes.Buffer
			_ = html.Render(&b, node)
			fmt.Printf("%s\n", b.String())
		}

	}
}
