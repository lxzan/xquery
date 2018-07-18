package parser

import (
	"regexp"
	"strings"
)

func Build(message string, bind Form) string {
	for k, v := range bind {
		var re = "{{" + k + "}}"
		message = strings.Replace(message, re, v, -1)
	}
	return message
}

// 去除两边引号
func FilteSlashes(s string) string {
	re, _ := regexp.Compile(`^["']{1}`)
	s = re.ReplaceAllString(s, "")
	re, _ = regexp.Compile(`["']{1}$`)
	return re.ReplaceAllString(s, "")
}

// 获取root标签名
func getTagName(html string) string {
	// 解析标签名
	re, _ := regexp.Compile(`(?imU:<[\S]+.*>)`)
	var s = re.FindString(html)
	s = strings.Replace(s, "<", "", 1)
	s = strings.Replace(s, ">", "", 1)
	re, _ = regexp.Compile(`(?i:[\S]+)`)
	var tag = re.FindString(s)
	return tag
}

// 获取root标签id, class以及其他属性
func getAttrs(tagName string, html string) (string, []string, Attrs) {
	var id = ""
	var classes = make([]string, 0)
	var attrs = Attrs{}

	var expr = Build(`(?im:^<{{tag}}.*?>)`, Form{"tag": tagName})
	re, _ := regexp.Compile(expr)
	var s = re.FindString(html)
	s = strings.Replace(s, "<"+tagName, "", 1)
	s = strings.Replace(s, "/>", "", 1)
	s = strings.TrimSpace(strings.Replace(s, ">", "", 1))
	re, _ = regexp.Compile(`[\s]{2,}`)
	s = strings.TrimSpace(re.ReplaceAllString(s, " "))

	var arr = make([]string, 0)
	re,_ = regexp.Compile(`(?imU:[a-z0-9\-]+={0,}".*")`)
	arr = re.FindAllString(s, -1)
	if len(arr) == 0{
		re,_ = regexp.Compile(`(?imU:[a-z0-9\-]+={0,}'.*')`)
		arr = re.FindAllString(s, -1)
	}
	if len(arr) == 1 && arr[0] == "" {
		return id, classes, attrs
	}

	for _, item := range arr {
		var kv = strings.Split(item, "=")
		var k = kv[0]
		var v *Any
		if len(kv) == 1 && boolAttrs[k] == true {
			v = NewAny(true)
		} else {
			v = NewAny(FilteSlashes(kv[1]))
		}

		if k == "id" {
			id = v.String()
		} else if k == "class" {
			re, _ := regexp.Compile(`[\s]{2,}`)
			str := re.ReplaceAllString(v.String(), " ")
			classes = strings.Split(str, " ")
		} else {
			attrs[k] = v
		}
	}
	return id, classes, attrs
}

func InnterHtml(html string) string{
	var tagName = getTagName(html)
	var expr = Build(`(?im:^<{{tag}}.*?>)`, Form{"tag": tagName})
	re, _ := regexp.Compile(expr)
	s := re.ReplaceAllString(html, "")

	expr = Build(`(?iUm:</{{tag}}>$)`, Form{"tag": tagName})
	re, _ = regexp.Compile(expr)
	s = re.ReplaceAllString(s, "")
	return strings.TrimSpace(s)
}

func DoWhile(fn func() bool)  {
	if fn() == true {
		DoWhile(fn)
	}
}

func Substr(s string, start, length int) string  {
	bytes := []byte(s)
	tmp := bytes[start:start+length]
	return string(tmp)
}

func InArray(arr []string, ele string) bool {
	for _, item := range arr {
		if item == ele {
			return true
		}
	}
	return false
}