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
<dl>
	<dt>
		<a href="/forums/CloudComputing">Cloud Calculate</a>
	</dt>
	<dd>
		<a href="/forums/AWS">IaaS</a>
		<a href="/forums/CloudFoundry">Pass/SaaS</a>
		<a href="/forums/hadoop">Cluster Calculate/Hadoop</a>
	</dd>
</dl>`
	var tree = Load(html)
	println(&tree)
	//<dt><a href="/forums/CloudComputing">云计算</a></dt> <dd> 	<a href="/forums/AWS">IaaS</a> 	<a hre

	//child,err := MatchChild(`<dd>
	//	<a href="/forums/AWS">IaaS</a>
	//	<a href="/forums/CloudFoundry">Pass/SaaS</a>
	//	<a href="/forums/hadoop">分布式计算/Hadoop</a>
	//</dd>`)
	//println(&child, err.Error())
}

func TestWhile(t *testing.T) {
	var i = 0
	DoWhile(func() bool {
		i++
		println(i)
		return i < 10
	})
}
