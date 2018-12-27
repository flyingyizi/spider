package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/query"
)

// https://cpu.baidu.com/api/1022/b9268e89/detail/25738154162040646/news?no_list=1&scene=2&forward=api&bside=1&api_version=2&appid=hwbrowser&ctype=news

func main() {
	startURL := "https://cpu.baidu.com/api/1022/b9268e89/detail/25738154162040646/news?no_list=1&scene=2&forward=api&bside=1&api_version=2&appid=hwbrowser&ctype=news"
	getbaiduNews(startURL)
}

func getbaiduNews(url string) {
	d := spider.DocCard{}

	var bodyData []byte
	// Instantiate default collector
	c := spider.NewCollector()

	c.OnResponse(func(r *spider.Response) {
		bodyData = r.Body
	})

	//xpath: `//head/title`
	c.OnHTML(`head>title`, func(e *query.HTMLElement) {
		d.Title = strings.TrimSpace(e.Text)
		fmt.Println(d.Title)
	})

	//正文+图片
	c.OnHTML(`#main-content>div>div>div>div.article.container >div`, func(e *query.HTMLElement) {
		//正文
		content := ""
		e.ForEach("p", func(_ int, ss *spider.HTMLElement) {
			content = content + strings.TrimSpace(ss.Text) + "\n"
		})
		//fmt.Println(content)
		d.Content = content

		//图片
		d.Images = make([]spider.Image, 0)
		e.ForEach("div", func(i int, s *spider.HTMLElement) {
			s.ForEach("div > img ", func(_ int, ss *spider.HTMLElement) {
				image := spider.Image{}
				image.Url = strings.TrimSpace(ss.Attr("data-src"))
				image.Width, _ = strconv.Atoi(strings.TrimSpace(ss.Attr("data-width")))
				image.Height, _ = strconv.Atoi(strings.TrimSpace(ss.Attr("data-height")))
				image.Index = i
				d.Images = append(d.Images, image)
				//fmt.Println(ss.Attr("data-src"))
			})
		})

	})

	c.OnScraped(func(_ *spider.Response) {
		bData, _ := json.MarshalIndent(d, "", "\t")
		fmt.Println(string(bData))
	})

	err := c.Visit(url)
	if err != nil {
		return
	}

	d.OrigHTML = base64.StdEncoding.EncodeToString(bodyData)
	d.Save("abc.json")

}
