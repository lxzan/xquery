package parser

import (
	"errors"
	"regexp"
	"strings"
)

// 匹配标签对
func MatchChild(html string) (string, error) {
	var childs = InnterHtml(html)
	var cp = childs
	var tagName = getTagName(childs)
	var startTag = "<" + tagName
	var startTagLength = len(startTag)
	empty, _ := regexp.Compile(`^[\s>]$`)
	var endTag = "</" + tagName + ">"
	if InArray(singleTags, strings.ToLower(tagName)) {
		patt := Build("(?imU:<{{tagName}}.*>)", Form{"tagName": tagName})
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
	if startIndex != -1 && empty.MatchString(string(childs[startIndex+startTagLength])) {
		a++
		childs = strings.Replace(childs, startTag, sa, 1)
	}
	var endIndex = -1

	for {
		if a == b && a != 0 {
			break
		}

		var j = strings.Index(childs, endTag)
		if j != -1 {
			b++
			childs = strings.Replace(childs, endTag, sb, 1)
		}

		var i = strings.Index(childs, startTag)
		if i != -1 && i < j && empty.MatchString(string(childs[i+startTagLength])) {
			a++
			childs = strings.Replace(childs, startTag, sa, 1)
		}

		if a == b && a != 0 {
			endIndex = j
		}

		if a == b {
			break
		}
	}

	if a == 0 {
		return "", errors.New("no match")
	}
	var child = Substr(cp, startIndex, endIndex-startIndex) + endTag
	return child, nil
}
