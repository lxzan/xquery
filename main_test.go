package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserLoad(t *testing.T) {
	var html = `<link    
	rel="stylesheet" 
	type="text/css"     
	href="//csdnimg.cn/pubfooter/css/pub_footer_1.0.3.css?v=201806111415">`
	Load(html)
}

func TestMatch(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./test/demo.html")
	node, _ := Load(string(bytes))
	//println(node.Find("#docs-collapse-btn").InnterHtml())
	nodes := node.FindAll("#docs-collapse-btn")
	//.ForEach(func(index int, node *Node) {
		println(nodes)
	//})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
