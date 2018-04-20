package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the front page
	url := "https://gate.io/fee"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	matcher := func(n *html.Node) bool {
		return n.DataAtom == atom.Tr &&
			(scrape.Attr(n, "class") == "odd" || scrape.Attr(n, "class") == "even") &&
			n.Parent != nil &&
			n.Parent.DataAtom == atom.Tbody &&
			n.Parent.Parent != nil
	}
	// grab all articles and print them
	nodes := scrape.FindAll(root, matcher)

	for _, n := range nodes {
		tds := scrape.FindAll(n, func(n *html.Node) bool { return n.DataAtom == atom.Td })
		logrus.Info(scrape.Text(tds[0]), "  ", scrape.Text(tds[1]), "  ", scrape.Text(tds[2]), "  ", scrape.Text(tds[3]), "  ", scrape.Text(tds[4]), "  ", scrape.Text(tds[4]), "  ", scrape.Text(tds[6]), "  ", scrape.Text(tds[7]))
	}
}
