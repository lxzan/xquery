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
	bytes,_ := ioutil.ReadFile("./test/demo.html")
	println(&bytes)
	nodes := Load(string(bytes)).Select("#wrapper .sr-only")
	//nodes := Load(`<gcse:search></gcse:search>`)
		//Select("#wrapper .sr-only")
	println(&nodes)
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
