package parser

import (
	"errors"
	"strings"
	"regexp"
)

// 匹配标签对
func MatchChild(html string) (string, error) {
	var childs = InnterHtml(html)
	var cp = childs
	var tagName = getTagName(childs)
	var startTag = "<" + tagName
	var endTag = "</" + tagName + ">"
	if InArray(singleTags, tagName) {
		patt := Build("<{{tagName}}.*?>", Form{"tagName": tagName})
		re, _ := regexp.Compile(patt)
		return re.FindString(html), nil
	}

	var a = 0
	var b = 0
	var sa = ""
	var sb = ""
	for i := 0; i < len(startTag); i++ {
		sa += "*"
	}
	for i := 0; i < len(endTag); i++ {
		sb += "*"
	}
	var startIndex = strings.Index(childs, startTag)
	if startIndex != -1 {
		a++
		childs = strings.Replace(childs, startTag, sa, 1)
	}
	var endIndex = -1

	DoWhile(func() bool {
		if a == b && a != 0 {
			return a != b
		}

		var j = strings.Index(childs, endTag)
		if j != -1 {
			b++
			childs = strings.Replace(childs, endTag, sb, 1)
		}

		var i = strings.Index(childs, startTag)
		if i != -1 && i < j {
			a++
			childs = strings.Replace(childs, startTag, sa, 1)
		}

		if a == b && a != 0 {
			endIndex = j
		}
		return a != b
	})

	if a == 0 {
		return "", errors.New("no match")
	}
	var child = Substr(cp, startIndex, endIndex-startIndex) + endTag
	return child, nil
}
