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
	//bytes, _ := hasaki.Get("https://github.com/").GetBody()
	bytes, _ := ioutil.ReadFile("./test/demo.html")
	html := string(bytes)
//	html := `
//<script>
//   with(document)with(body)with(insertBefore(createElement("script"),firstChild))setAttribute("exparams","category=&userid=&aplus&udpid=&&yunid=&&trid=9dff284515322261753058595e&asid=AQAAAAB/6lNbd85RSAAAAACD74CxVu+bXQ==",id="tb-beacon-aplus",src=(location>"https"?"//g":"//g")+".alicdn.com/alilog/mlog/aplus_v2.js")
//</script>
//`

	node, err := Load(html)
	//node.Find("auto-check").Attr("src")

	node.FindAll("auto-check").ForEach(func(index int, node *Node) {
		println(&node, &err)
	})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
