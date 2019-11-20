package umlsrest

// This exists only to parse the CAS response

import (
	"golang.org/x/net/html"
	"io"
)

type Form struct {
	Action	string
	Method  string
}

func DecodeTGTMessage(responseBody io.ReadCloser)(Form, error){
	var frm Form
	doc, _ := html.Parse(responseBody)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "form" {
			for _, a := range n.Attr {
				if a.Key == "action" {
					frm.Action = a.Val
					break
				} else if a.Key == "method" {
					frm.Method = a.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return frm, nil
}

