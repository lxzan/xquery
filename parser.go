package parser

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/kataras/iris/core/errors"
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
var singleTags = []string{"img", "input", "hr", "br", "link", "meta", "source", "path"}

func Load(html string) (*Node, error) {
	html = strings.Replace(html, "\n\t", " ", -1)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.TrimSpace(html)
	for _, item := range singleTags {
		tag := strings.ToLower("</" + item + ">")
		html = strings.Replace(html, tag, "", -1)
		tag = strings.ToUpper("</" + item + ">")
		html = strings.Replace(html, tag, "", -1)
	}
	re, _ := regexp.Compile(`(?imU:<!doctype.*>)`)
	html = re.ReplaceAllString(html, "")
	re, _ = regexp.Compile(`(?imU:<!--.*-->)`)
	html = strings.TrimSpace(re.ReplaceAllString(html, ""))

	if valid(html) == false {
		return nil, errors.New("html not valid")
	}
	return build(html), nil
}

type Node struct {
	html     string
	tagName  string
	classes  []string
	id       string
	attrs    Attrs
	children []*Node
}

func valid(html string) bool {
	re, _ := regexp.Compile(`(?m:^<.*>$)`)
	return re.MatchString(html)
}

func build(html string) *Node {
	var obj = new(Node)
	obj.attrs = Attrs{}
	obj.classes = []string{}
	obj.children = make([]*Node, 0)

	obj.html = html
	obj.tagName = getTagName(obj.html)
	obj.id, obj.classes, obj.attrs = getAttrs(obj.tagName, obj.html)
	var cp = obj.html

	for {
		if cp == "" {
			break
		}
		child, err := MatchChild(cp)
		if err == nil && child != "" {
			cp = strings.TrimSpace(strings.Replace(cp, child, "", 1))
			obj.children = append(obj.children, build(child))
		} else {
			break
		}
	}
	return obj
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

func (u *Node) query(selector string, limit int) []*Node {
	var res = make([]*Node, 0)
	var arr = strings.Split(selector, " ")
	type Q struct {
		N *Node
		A []string
	}
	var q = arraylist.New()
	var v = Q{
		N: u,
		A: arr,
	}
	q.Add(v)

	for q.Size() > 0 {
		var flag = 0
		i, _ := q.Get(0)
		q.Remove(0)
		v = i.(Q)
		if len(v.A) == 0 {
			res = append(res, v.N)
			if len(res) == limit {
				break
			}
		} else {
			var patt = v.A[0]
			if patt[0] == '#' {
				re, _ := regexp.Compile(`(?i:^#[0-9a-z\-]+)`)
				var id = re.FindString(patt)
				id = strings.Replace(id, "#", "", 1)
				if v.N.id == id {
					flag++
					patt = strings.Replace(patt, "#"+id, "", 1)
					if patt == "" && len(v.A) == 1 {
						q.Add(Q{
							N: v.N,
							A: make([]string, 0),
						})
					} else if patt != "" && len(patt) < len(arr[0]) {
						var tmp = []string{patt}
						for i := 1; i < len(arr); i++ {
							tmp = append(tmp, arr[i])
						}
						q.Add(Q{
							N: v.N,
							A: tmp,
						})
					} else if patt == "" && len(v.A) > 1 {
						for _, son := range v.N.children {
							q.Add(Q{
								N: son,
								A: v.A[1:len(v.A)],
							})
						}
					}
				}
			} else if patt[0] == '.' {
				re, _ := regexp.Compile(`(?i:^.[0-9a-z\-]+)`)
				var myclass = re.FindString(patt)
				myclass = strings.Replace(myclass, ".", "", 1)

				if InArray(v.N.classes, myclass) == true {
					flag++
					patt = strings.Replace(patt, "."+myclass, "", 1)
					if patt == "" && len(v.A) == 1 {
						q.Add(Q{
							N: v.N,
							A: make([]string, 0),
						})
					} else if patt != "" && len(patt) < len(arr[0]) {
						var tmp = []string{patt}
						for i := 1; i < len(arr); i++ {
							tmp = append(tmp, arr[i])
						}
						q.Add(Q{
							N: v.N,
							A: tmp,
						})
					} else if patt == "" && len(v.A) > 1 {
						for _, son := range v.N.children {
							q.Add(Q{
								N: son,
								A: v.A[1:len(v.A)],
							})
						}
					}
				}
			} else if (patt[0] >= 'a' && patt[0] <= 'z') || (patt[0] >= 'A' && patt[0] <= 'Z') {
				re, _ := regexp.Compile(`(?i:^[0-9a-z\-]+)`)
				var tagName = re.FindString(patt)
				tagName = strings.Replace(tagName, ".", "", 1)

				if v.N.tagName == tagName {
					flag++
					patt = strings.Replace(patt, tagName, "", 1)
					if patt == "" && len(v.A) == 1 {
						q.Add(Q{
							N: v.N,
							A: make([]string, 0),
						})
					} else if patt != "" && len(patt) < len(arr[0]) {
						var tmp = []string{patt}
						for i := 1; i < len(arr); i++ {
							tmp = append(tmp, arr[i])
						}
						q.Add(Q{
							N: v.N,
							A: tmp,
						})
					} else if patt == "" && len(v.A) > 1 {
						for _, son := range v.N.children {
							q.Add(Q{
								N: son,
								A: v.A[1:len(v.A)],
							})
						}
					}
				}
			} else {
				re, _ := regexp.Compile(`(?iU:^\[.*\])`)
				var s = re.FindString(patt)
				var cp = s
				s = strings.Replace(s, "[", "", 1)
				s = strings.Replace(s, "]", "", 1)
				var kv = strings.Split(s, "=")
				var k = kv[0]
				re, _ = regexp.Compile(`['"]`)
				var val = re.ReplaceAllString(kv[1], "")
				attr, ok := v.N.attrs[k]
				if ok && attr.String() == val {
					flag++
					patt = strings.Replace(patt, cp, "", 1)
					if patt == "" && len(v.A) == 1 {
						q.Add(Q{
							N: v.N,
							A: make([]string, 0),
						})
					} else if patt != "" && len(patt) < len(arr[0]) {
						var tmp = []string{patt}
						for i := 1; i < len(arr); i++ {
							tmp = append(tmp, arr[i])
						}
						q.Add(Q{
							N: v.N,
							A: tmp,
						})
					} else if patt == "" && len(v.A) > 1 {
						for _, son := range v.N.children {
							q.Add(Q{
								N: son,
								A: v.A[1:len(v.A)],
							})
						}
					}
				}
			}

			if flag == 0 {
				for _, son := range v.N.children {
					q.Add(Q{
						N: son,
						A: v.A,
					})
				}
			}
		}
	}
	return res
}

func (u *Node) FindAll(selector string) *Nodes {
	var nodes = Nodes{
		data: u.query(selector, -1),
	}
	return &nodes
}

func (u *Node) Find(selector string) *Node {
	nodes := u.query(selector, 1)
	if len(nodes) > 0 {
		return nodes[0]
	}
	return &Node{}
}

func (u *Node) Attr(key string) string {
	return u.attrs[key].String()
}

func (u *Node) Text() string {
	re, _ := regexp.Compile("<.*?>")
	return strings.TrimSpace(re.ReplaceAllString(u.html, ""))
}

type Nodes struct {
	data []*Node
}

func (u *Nodes) ForEach(f func(index int, node *Node)) {
	for i, item := range u.data {
		f(i, item)
	}
}

func (u *Nodes) Result() []*Node {
	return u.data
}
