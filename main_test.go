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
	html = `

	<i href="https://bbs.feng.com/read-htm-tid-11790208.html">
		<i ><img src="https://bbsimages.feng.com/data/attachment/common/41/common_371_icon.png"></i>
		<span>【汇总】苹果教育优惠来威锋下单，三重大礼等你拿！</span>
	</i>
`
	node, _ := Load(html)
	//node.FindAll("#docs-collapse-btn").ForEach(func(index int, node *Node) {
		println(&node)
	//})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
