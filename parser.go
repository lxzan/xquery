package parser

import (
	"regexp"
	"strings"
)

var boolAttrs = Switch{
	"autofocus": true,
	"autoplay":  true,
	"async":     true,
	"checked":   true,
	"controls":  true,
	"defer":     true,
	"disabled":  true,
	"hidden":    true,
	"loop":      true,
	"multiple":  true,
	"open":      true,
	"readonly":  true,
	"required":  true,
	"scoped":    true,
	"selected":  true,
}

func Load(html string) *Node {
	html = strings.Replace(html, "\n\t", " ", -1)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.TrimSpace(html)
	var obj = new(Node)
	obj.attrs = &Attrs{}
	obj.classes = []string{}
	re, _ := regexp.Compile(`(?i:^<!DOCTYPE html.*?>)`)
	obj.html = strings.TrimSpace(re.ReplaceAllString(html, ""))
	obj.tagName = getTagName(obj.html)
	obj.id, obj.classes, obj.attrs = getAttrs(obj.tagName, obj.html)

	var cp = obj.html
	DoWhile(func() bool {
		if cp == "" {
			return false
		}
		child, err := MatchChild(cp)
		if err == nil{
			cp = strings.TrimSpace(strings.Replace(cp, child, "", 1))
			obj.children = append(obj.children, Load(child))
		}
		return err == nil
	})
	return obj
}

type Node struct {
	html     string
	tagName  string
	classes  []string
	id       string
	attrs    *Attrs
	children []*Node
}

func (u *Node) InnterHtml() string {
	var expr = Build(`(?im:^<{{tag}}.*?>)`, Form{"tag": u.tagName})
	re, _ := regexp.Compile(expr)
	s := re.ReplaceAllString(u.html, "")

	expr = Build(`(?iUm:</{{tag}}>$)`, Form{"tag": u.tagName})
	re, _ = regexp.Compile(expr)
	s = re.ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}
