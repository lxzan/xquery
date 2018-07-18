package parser

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
	"regexp"
	"strings"
)

// 允许省略值得属性
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

// 非闭合标签
var singleTags = []string{"img", "input", "hr", "br", "link", "meta", "source"}

func Load(html string) *Node {
	html = strings.Replace(html, "\n\t", " ", -1)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.TrimSpace(html)
	for _, item := range singleTags {
		tag := "</" + item + ">"
		html = strings.Replace(html, tag, "", -1)
	}
	var obj = new(Node)
	obj.attrs = Attrs{}
	obj.classes = []string{}
	re, _ := regexp.Compile(`(?i:^<!DOCTYPE html.*?>)`)
	html = strings.TrimSpace(re.ReplaceAllString(html, ""))
	return obj.build(html)
}

type Node struct {
	html     string
	tagName  string
	classes  []string
	id       string
	attrs    Attrs
	children []*Node
}

func (u *Node) build(html string) *Node {
	u.html = html
	u.tagName = getTagName(u.html)
	u.id, u.classes, u.attrs = getAttrs(u.tagName, u.html)
	var cp = u.html

	for  {
		if cp == "" {
			break
		}
		child, err := MatchChild(cp)
		if err == nil {
			cp = strings.TrimSpace(strings.Replace(cp, child, "", 1))
			u.children = append(u.children, u.build(child))
		} else {
			break
		}
	}
	return u
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

func (u *Node) Select(selector string) []*Node {
	var res = make([]*Node, 0)
	var arr = strings.Split(selector, " ")
	type Q struct {
		N *Node
		A []string
	}
	var q = linkedliststack.New()
	var v = Q{
		N: u,
		A: arr,
	}
	q.Push(v)

	DoWhile(func() bool {
		var flag = 0
		i, _ := q.Pop()
		v = i.(Q)
		if len(v.A) == 0 {
			res = append(res, v.N)
		} else {
			var patt = v.A[0]
			if patt[0] == '#' {
				re, _ := regexp.Compile("(?i:^#[0-9a-z]+)")
				var id = re.FindString(patt)
				id = strings.Replace(id, "#", "", 1)
				if v.N.id == id {
					flag++
					patt = strings.Replace(patt, "#"+id, "", 1)
					if patt == "" && len(v.A) == 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						q.Push(Q{
							N: v.N,
							A: a,
						})
					} else if patt == "" && len(v.A) > 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						for _, son := range v.N.children {
							q.Push(Q{
								N: son,
								A: a,
							})
						}
					} else {
						a := make([]string, len(v.A))
						copy(a, v.A)
						a[0] = patt
						q.Push(Q{
							N: v.N,
							A: a,
						})
					}
				}
			} else if patt[0] == '.' {
				re, _ := regexp.Compile("(?i:^.[0-9a-z]+)")
				var myclass = re.FindString(patt)
				myclass = strings.Replace(myclass, ".", "", 1)

				if InArray(v.N.classes, myclass) == true {
					flag++
					patt = strings.Replace(patt, "."+myclass, "", 1)
					if patt == "" && len(v.A) == 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						q.Push(Q{
							N: v.N,
							A: a,
						})
					} else if patt == "" && len(v.A) > 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						for _, son := range v.N.children {
							q.Push(Q{
								N: son,
								A: a,
							})
						}
					} else {
						a := make([]string, len(v.A))
						copy(a, v.A)
						a[0] = patt
						q.Push(Q{
							N: v.N,
							A: a,
						})
					}
				}
			} else if (patt[0] >= 'a' && patt[0] <= 'z') || (patt[0] >= 'A' && patt[0] <= 'Z') {
				re, _ := regexp.Compile("(?i:^[0-9a-z]+)")
				var tagName = re.FindString(patt)
				tagName = strings.Replace(tagName, ".", "", 1)

				if v.N.tagName == tagName {
					flag++
					patt = strings.Replace(patt, tagName, "", 1)
					if patt == "" && len(v.A) == 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						q.Push(Q{
							N: v.N,
							A: a,
						})
					} else if patt == "" && len(v.A) > 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						for _, son := range v.N.children {
							q.Push(Q{
								N: son,
								A: a,
							})
						}
					} else {
						a := make([]string, len(v.A))
						copy(a, v.A)
						a[0] = patt
						q.Push(Q{
							N: v.N,
							A: a,
						})
					}
				}
			} else {
				re, _ := regexp.Compile(`(?i:^\[.*?\])`)
				var s = re.FindString(patt)
				var cp = s
				s = strings.Replace(s, "[", "", 1)
				s = strings.Replace(s, "]", "", 1)
				var kv = strings.Split(s, "=")
				var k = kv[0]
				re, _ = regexp.Compile(`['"]`)
				var val = re.ReplaceAllString(kv[1], "")
				if v.N.attrs[k].String() == val {
					flag++
					patt = strings.Replace(patt, cp, "", 1)
					if patt == "" && len(v.A) == 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						q.Push(Q{
							N: v.N,
							A: a,
						})
					} else if patt == "" && len(v.A) > 1 {
						a := make([]string, len(v.A)-1)
						copy(a, v.A[1:len(v.A)])
						for _, son := range v.N.children {
							q.Push(Q{
								N: son,
								A: a,
							})
						}
					} else {
						a := make([]string, len(v.A))
						copy(a, v.A)
						a[0] = patt
						q.Push(Q{
							N: v.N,
							A: a,
						})
					}
				}
			}

			if flag == 0 {
				for _, son := range v.N.children {
					q.Push(Q{
						N: son,
						A: arr,
					})
				}
			}
		}
		return q.Size() > 0
	})
	return res
}

func (u *Node) Attr(key string) string {
	return u.attrs[key].String()
}

func (u *Node) Text() string {
	re, _ := regexp.Compile("<.*?>")
	return strings.TrimSpace(re.ReplaceAllString(u.html, ""))
}
