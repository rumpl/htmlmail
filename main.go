package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func main() {
	if err := run(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func run(file string) error {
	converter := md.NewConverter("", true, &md.Options{})
	converter.AddRules(
		md.Rule{
			Filter: []string{"br"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				content = strings.TrimSpace(content)
				return md.String(content + "\n")
			},
		},
	)

	fd, err := os.Open(file)
	if err != nil {
		return errors.Wrap(err, "open file")
	}

	b, err := converter.ConvertReader(fd)
	if err != nil {
		return errors.Wrap(err, "html to text")
	}

	fmt.Print(b.String())

	return nil
}
