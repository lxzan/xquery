package parser

import (
	"testing"
	"io/ioutil"
)

func TestParserLoad(t *testing.T) {
	var html = `<link    
	rel="stylesheet" 
	type="text/css"     
	href="//csdnimg.cn/pubfooter/css/pub_footer_1.0.3.css?v=201806111415">`
	Load(html)
}

func TestMatch(t *testing.T) {
	//resp, _ := hasaki.Get("https://www.feng.com/").GetBody()
	bytes, _ := ioutil.ReadFile("./test/demo.html")
	html := string(bytes)
	node, _ := Load(html)
	node.FindAll(".gobackBtn span").ForEach(func(index int, node *Node) {
		println(&node)
	})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
