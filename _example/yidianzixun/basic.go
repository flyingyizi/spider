package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/query"
)

func main() {
	startURL := "https://hw.yidianzixun.com/article/0KwMhPzx?s=hwbrowser&appid=hwbrowser&ctype=news"
	getYidianZixunNews(startURL)
}

//todo
func getYidianZixunNews(url string) {
	d := spider.DocCard{}

	var (
		bodyData  []byte
		mapResult map[string]string
	)

	// Instantiate default collector
	c := spider.NewCollector()

	c.OnResponse(func(r *spider.Response) {
		bodyData = r.Body

		//
		//通过观察一点资讯的网页，图片网址存放在js代码中，下面从js代码中将感兴趣的内容解析出来
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyData))
		if err != nil {
			return
		}
		scripts := htmlDoc.Find("script")
		destScript := ""
		key1 := "globalInfo.show_imgs"
		for _, n := range scripts.Nodes {
			if n.FirstChild != nil {
				//找了存放内容js代码的特征
				if strings.Contains(n.FirstChild.Data, key1) {
					destScript = n.FirstChild.Data
					break
				}
			}
		}
		//解析出js代码中 "globalInfo.show_imgs"变量的值
		startIndex := strings.Index(destScript, key1)
		imageJSONData := getContentInBrace(destScript, startIndex)

		json.Unmarshal(imageJSONData, &mapResult)
		//fmt.Println(mapResult)

	})

	//xpath: `//head/title`
	c.OnHTML(`head>title`, func(e *query.HTMLElement) {
		d.Title = strings.TrimSpace(e.Text)
		//fmt.Println(d.Title)
	})

	//正文+图片
	c.OnHTML(`#imedia-article`, func(e *query.HTMLElement) {
		//正文
		content := ""
		e.ForEach("p", func(_ int, ss *spider.HTMLElement) {
			temp := strings.TrimSpace(ss.Text)
			if temp != "" {
				content = content + temp + "\n"
			}
		})

		//fmt.Println(content)
		d.Content = content

		//图片
		d.Images = make([]spider.Image, 0)
		e.ForEach("div", func(i int, s *spider.HTMLElement) {
			image := spider.Image{}
			//image.Url = strings.TrimSpace(ss.Attr("data-src"))
			image.Index = i

			style := s.Attr("style")
			wr, _ := regexp.Compile("width:([0-9]+)px;height:([0-9]+)px")
			unr := wr.FindStringSubmatch(style)
			image.Width, _ = strconv.Atoi(strings.TrimSpace(unr[1]))
			image.Height, _ = strconv.Atoi(strings.TrimSpace(unr[2]))

			id := s.Attr("id")
			id = strings.Replace(id[8:], "-", "_", -1)
			if v, ok := mapResult[id]; ok {
				image.Url = v
			}
			//fmt.Println(image)
			d.Images = append(d.Images, image)

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

//抽取大括号中的内容,含定界的`{}`
func getContentInBrace(str string, startIndex int) []byte {
	buf := bytes.NewBufferString("")
	var count int = 0
	for i := startIndex; i < len(str); i++ {
		if str[i] == '{' {
			count++
		}
		if count != 0 {
			buf.WriteByte(str[i])
		}
		if str[i] == '}' {
			count--
			if count == 0 {
				break
			}
		}
	}
	if count != 0 {
		return nil
	}

	return buf.Bytes() //
}
