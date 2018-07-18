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
	nodes := Load(string(bytes))
	//nodes := Load(`<meta name="author" content="slene, Unknown" />`)
		//Select("#wrapper .sr-only")
	println(&nodes)
}

func TestWhile(t *testing.T) {
	var i = 0
	DoWhile(func() bool {
		i++
		println(i)
		return i < 10
	})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
