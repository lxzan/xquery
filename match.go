package parser

import (
	"github.com/kataras/iris/core/errors"
	"regexp"
	"strings"
)

// 匹配标签对
func MatchChild(html string) (string, error) {
	var matchCount = 0
	if getTagName(html) == "script" {
		return "", nil
	}
	var childs = InnterHtml(html)
	if childs == "" {
		return "", nil
	}

	re, _ := regexp.Compile(`(?imU:<.*>)`)
	if !re.MatchString(childs) {
		return "", nil
	}

	var cp = childs
	var tagName = getTagName(childs)
	var startTag = "<" + tagName
	var startTagLength = len(startTag)
	empty, _ := regexp.Compile(`^[\s>]$`)
	var endTag = "</" + tagName + ">"
	if InArray(singleTags, strings.ToLower(tagName)) || strings.Index(cp, endTag) == -1 {
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
		matchCount++
		if matchCount > 10000 {
			return "", errors.New("html parse error")
		}
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
		return "", nil
	}
	var child = Substr(cp, startIndex, endIndex-startIndex) + endTag
	return child, nil
}
