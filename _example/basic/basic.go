package main

import (
	"fmt"

	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/core"
	"github.com/flyingyizi/spider/query"
)

// https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=3091260259678083253&cid=100&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=859544137256752847&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D
// https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=16119487478051582585&cid=179223212&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=7511051609959712598&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D

func main() {
	startURL := "https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=3091260259678083253&cid=100&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=859544137256752847&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D"
	// Instantiate default collector
	c := spider.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	//spider.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	//*[@id="wrapper"]/div[13]
	// On every a element which has href attribute call callback
	c.OnHTML(`div[data-region="article_relate"]`, func(e *query.HTMLElement) {
		fmt.Println(e.Text)
		// link := e.Attr("href")
		// // Print link
		// fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// // Visit link found on page
		// // Only those links are visited which are in AllowedDomains
		// c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *core.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *spider.Response) {

		d := spider.DocCard{}
		d.Title = string(r.Body)

		d.Save("abc.json")

		//fmt.Println(string())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(startURL)
}
