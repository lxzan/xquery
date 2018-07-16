package parser

import (
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
	html := `
<dl id="cloud" name="caster">
	<dt>
		<a href="/forums/CloudComputing">云计算</a>
		<img src="https://baidu.com/">
	</dt>

	<dd>
	<a href="/forums/AWS">IaaS</a>
	<a href="/forums/CloudFoundry">Pass/SaaS</a>
	<a href="/forums/hadoop">Cluster Calculate/Hadoop</a>
	</dd>
</dl>`

	var tree = Load(html)
	nodes := tree.Select("dl dt img")
	println(nodes[0].Attr("src"))
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
