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

	html = `
  <div class="position-relative"> 
	<!-- </textarea></xmp> -->
	<form class="js-site-search-form" data-unscoped-search-url="/search" action="/search" accept-charset="UTF-8" method="get">
	<input name="utf8" type="hidden" value="✓" />
	<label class="form-control header-search-wrapper header-search-wrapper-jump-to position-relative d-flex flex-justify-between flex-items-center js-chromeless-input-container"> <input type="text" class="form-control header-search-input jump-to-field js-jump-to-field js-site-search-focus " data-hotkey="s,/" name="q" value="" placeholder="Search GitHub" data-unscoped-placeholder="Search GitHub" data-scoped-placeholder="Search" autocapitalize="off" aria-autocomplete="list" aria-controls="jump-to-results" data-jump-to-suggestions-path="/_graphql/GetSuggestedNavigationDestinations#csrf-token=JcTpW9NAfcJjli0OaoGpGf2d1BR/hRvhjQW/ovoj01ioxdsu8nvufq9aLn8dVIYXakzPDDVMoFUnbTFZz8WE8g==" spellcheck="false" autocomplete="off" /> <input type="hidden" class="js-site-search-type-field" name="type" /> <img src="https://assets-cdn.github.com/images/search-shortcut-hint.svg" alt="" class="mr-2 header-search-key-slash" />
	<div class="Box position-absolute overflow-hidden d-none jump-to-suggestions js-jump-to-suggestions-container">
	<ul class="d-none js-jump-to-suggestions-template-container">
	<li class="d-flex flex-justify-start flex-items-center p-0 f5 navigation-item js-navigation-item"> <a tabindex="-1" class="no-underline d-flex flex-auto flex-items-center p-2 jump-to-suggestions-path js-jump-to-suggestion-path js-navigation-open" href="">
	<div class="jump-to-octicon js-jump-to-octicon mr-2 text-center d-none"></div> <img class="avatar mr-2 flex-shrink-0 js-jump-to-suggestion-avatar" alt="" aria-label="Team" src="" width="28" height="28" />
	<div class="jump-to-suggestion-name js-jump-to-suggestion-name flex-auto overflow-hidden text-left no-wrap css-truncate css-truncate-target">
	</div>
	<div class="border rounded-1 flex-shrink-0 bg-gray px-1 text-gray-light ml-1 f6 d-none js-jump-to-badge-search">
	<span class="js-jump-to-badge-search-text-default d-none" aria-label="in all of GitHub"> Search </span>
	<span class="js-jump-to-badge-search-text-global d-none" aria-label="in all of GitHub"> All GitHub </span>
	<span aria-hidden="true" class="d-inline-block ml-1 v-align-middle">↵</span>
	</div>
	<div aria-hidden="true" class="border rounded-1 flex-shrink-0 bg-gray px-1 text-gray-light ml-1 f6 d-none d-on-nav-focus js-jump-to-badge-jump">
		Jump to
	<span class="d-inline-block ml-1 v-align-middle">↵</span>
	</div> </a> </li>
	<svg height="16" width="16" class="octicon octicon-repo flex-shrink-0 js-jump-to-repo-octicon-template" title="Repository" aria-label="Repository" viewbox="0 0 12 16" version="1.1" role="img">
	<path fill-rule="evenodd" d="M4 9H3V8h1v1zm0-3H3v1h1V6zm0-2H3v1h1V4zm0-2H3v1h1V2zm8-1v12c0 .55-.45 1-1 1H6v2l-1.5-1.5L3 16v-2H1c-.55 0-1-.45-1-1V1c0-.55.45-1 1-1h10c.55 0 1 .45 1 1zm-1 10H1v2h2v-1h3v1h5v-2zm0-10H2v9h9V1z" />
	</svg>
	<svg height="16" width="16" class="octicon octicon-project flex-shrink-0 js-jump-to-project-octicon-template" title="Project" aria-label="Project" viewbox="0 0 15 16" version="1.1" role="img">
	<path fill-rule="evenodd" d="M10 12h3V2h-3v10zm-4-2h3V2H6v8zm-4 4h3V2H2v12zm-1 1h13V1H1v14zM14 0H1a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1z" />
	</svg>
	<svg height="16" width="16" class="octicon octicon-search flex-shrink-0 js-jump-to-search-octicon-template" title="Search" aria-label="Search" viewbox="0 0 16 16" version="1.1" role="img">
	<path fill-rule="evenodd" d="M15.7 13.3l-3.81-3.83A5.93 5.93 0 0 0 13 6c0-3.31-2.69-6-6-6S1 2.69 1 6s2.69 6 6 6c1.3 0 2.48-.41 3.47-1.11l3.83 3.81c.19.2.45.3.7.3.25 0 .52-.09.7-.3a.996.996 0 0 0 0-1.41v.01zM7 10.7c-2.59 0-4.7-2.11-4.7-4.7 0-2.59 2.11-4.7 4.7-4.7 2.59 0 4.7 2.11 4.7 4.7 0 2.59-2.11 4.7-4.7 4.7z" />
	</svg>
	</ul>
	<ul class="d-none js-jump-to-no-results-template-container">
	<li class="d-flex flex-justify-center flex-items-center p-3 f5 d-none"> <span class="text-gray">No suggested jump to results</span> </li>
	</ul>
	<ul id="jump-to-results" class="js-navigation-container jump-to-suggestions-results-container js-jump-to-suggestions-results-container">
	<li class="d-flex flex-justify-center flex-items-center p-0 f5"> <img src="https://assets-cdn.github.com/images/spinners/octocat-spinner-128.gif" alt="Octocat Spinner Icon" class="m-2" width="28" /> </li>
	</ul>
	</div> </label>
	</form>
	</div>
`
	node, _ := Load(html)
	//node.FindAll(".gobackBtn span").ForEach(func(index int, node *Node) {
		println(&node)
	//})
}

func TestSubstr(t *testing.T) {
	var s = "lxz 520"
	son := Substr(s, 1, 2)
	println(son)
}
