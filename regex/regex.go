package main

import (
	"fmt"
	"regexp"
)

const content = `<script>!function(n){var t=n(".ct-nba-header"),s=[],e=null,a="nba-header-nav-submenu-enable",o=t.find("#ct-nav");o.on("mouseenter",function(){e=setTimeout(function(){for(o.addClass(a);0<s.length;)s.pop().call()},300)}).on("mouseleave",function(){e&&clearTimeout(e),o.removeClass(a)}),t.on("mouseenter","[data-bosshover]",function(){"function"==typeof window.ExposureBoss&&window.ExposureBoss(null,n(this).attr("data-bosshover"))}),s.push(function(){n.ajax({url:"//matchweb.sports.qq.com/rank/team?competitionId=100000&from=NBA_PC",dataType:"jsonp",scriptCharset:"utf-8",success:function(t){function s(t){var e="";return e+="<ul>",n.each(t,function(t,s){3<s.name.length?e+='<li><a href="https://sports.qq.com/kbsweb/teams.htm?tid='+s.teamId+'&cid=100000"  target="_blank"><img class="logo" src="'+s.logoNew+'" alt="'+s.name+'" /><strong class="line2">'+s.name.substr(0,2)+"<br/>"+s.name.substr(2)+"</strong></a></li>":e+='<li><a href="https://sports.qq.com/kbsweb/teams.htm?tid='+s.teamId+'&cid=100000"  target="_blank"><img class="logo" src="'+s.logoNew+'" alt="'+s.name+'" /><strong>'+s.name+"</strong></a></li>"}),e+="</ul>"}n("#ls-atlantic").html(s(t[1].atlantic)),n("#ls-westnorth").html(s(t[1].westnorth)),n("#ls-central").html(s(t[1].central)),n("#ls-pacific").html(s(t[1].pacific)),n("#ls-eastsouth").html(s(t[1].eastsouth)),n("#ls-westsouth").html(s(t[1].westsouth))}})});var l=t.find(".nba-header-top"),r=!1;/\bMSIE 6/.test(navigator.userAgent)&&!window.opera||n(window).on("scroll",function(){n(window).scrollTop()>l.height()?r||(t.addClass("nba-header-magnet"),t.find(".submenu").addClass("submenu-top"),r=!0):r&&(t.removeClass("nba-header-magnet"),t.find(".submenu").removeClass("submenu-top"),r=!1)})}(jQuery);</script>
<script src="//mat1.gtimg.com/pingjs/ext2020/sports/libs/dist/login.js"></script>
<script>
`

func main() {
	compile := regexp.MustCompile("(n\\.ajax\\({url:\"//)([a-zA-Z0-9&=_/.?]+)(\")")
	s1 := compile.FindStringSubmatch(content)
	fmt.Println(s1[2])

	compile2 := regexp.MustCompile("(https://[a-zA-Z0-9&=_/.?]+)(['+s.teamId]*)([&cid=0-9]*)")
	s2 := compile2.FindStringSubmatch(content)
	fmt.Println(s2[1])
	fmt.Println(s2[3])
}
