package html_link_parser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href, Text string
}

func GetLinks(htmlSource io.ReadCloser) []Link {
	fmt.Println("HTML link parser!")
	// file_ptr, err := os.Open("E:\\Go\\go projects\\html_link_parser\\ex2.html")
	// if err != nil {
	// 	panic(err)
	// }

	var links []Link
	html_tokenizer := html.NewTokenizer(htmlSource)
	for {
		tt := html_tokenizer.Next()
		token := html_tokenizer.Token()
		if tt == html.ErrorToken {
			break
		}
		if tt == html.StartTagToken && token.DataAtom.String() == "a" {
			attrs := token.Attr
			var url string
			for _, attr := range attrs {
				if attr.Key == "href" {
					url = attr.Val
					break
				}
			}
			texts := getLinkTexts(html_tokenizer)
			link := Link{Href: url, Text: texts}
			links = append(links, link)
		}
	}
	return links
	// val, err := json.Marshal(links)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(val))
}

func getLinkTexts(html_tokenizer *html.Tokenizer) string {
	var texts []string
	for {
		tag := html_tokenizer.Next()
		token := html_tokenizer.Token()
		if tag == html.EndTagToken && token.DataAtom.String() == "a" {
			break
		}
		if tag == html.TextToken {
			texts = append(texts, strings.Fields(token.Data)...)
		}
	}
	return strings.Join(texts, " ")
}
