package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
)

func main() {
	if err := run(os.Args[1]); err != nil {
		log.Fatal(err)
	}
}

func run(file string) error {
	converter := md.NewConverter("", true, &md.Options{
		BulletListMarker: "*",
	})
	converter.AddRules(
		md.Rule{
			Filter: []string{"br"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				content = strings.TrimSpace(content)
				return md.String(content + "\n")
			},
		},
		md.Rule{
			Filter: []string{"strong", "b", "h1", "h2", "h3", "h4", "h5", "h6"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				content = strings.TrimSpace(content)
				return md.String(aurora.Bold(content).String() + "\n")
			},
		},
		md.Rule{
			Filter: []string{"i"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				content = strings.TrimSpace(content)
				return md.String(aurora.Italic(content).String() + "\n")
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
