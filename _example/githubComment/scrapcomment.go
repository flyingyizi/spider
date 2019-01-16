package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/query"
)

type IssueComments struct {
}

//IssueComment 存放内容的卡片
type IssueComment struct {
	Time     string    `json:"time"`
	Title    string    `json:"title"` //,omitempty在序列化的时候忽略0值或者空值
	Contents []Content `json:"contents"`
}
type Content struct {
	Text string `json:"text"`
}

// Save saves model in the file
func (f *IssueComment) Save(path string) error {

	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(f)
	}
	file.Close()
	return err
}

// https://github.com/fxsjy/jieba/issues/87

func main() {
	//startURL := "https://github.com/fxsjy/jieba/issues/87"
	startURL1 := "https://github.com/fxsjy/jieba/issues"
	urls := spider.UrlFlags(startURL1)

	start := time.Now()
	for i, url := range urls {
		out := getIssueList(url)
		fmt.Printf("------output-----%d url----------\n", i)
		fmt.Println(out)

	}
	elapsed := time.Since(start)
	fmt.Printf("Time required to complete: %s\n", elapsed)
}

//爬取issue list url，以便后续爬取评论内容
func getIssueList(url string) []string {
	//urls := make([]string, 0)

	c := spider.NewCollector()

	// c.OnResponse(func(r *spider.Response) {
	// 	bodyData = r.Body
	// })

	//主题列表
	c.OnHTML(`div.repository-content > div > div.border-right.border-bottom.border-left > div > div > div > div >a `, func(e *query.HTMLElement) {

		// content := ""
		// e.ForEach("div > a", func(_ int, ss *spider.HTMLElement) {
		// 	temp := ss.Attr(`href`)
		// 	if !strings.EqualFold(temp, "") {
		// 		content = content + ss.Attr(`href`) + ":" + strings.TrimSpace(ss.Text) + "\n"
		// 	}
		// })
		// fmt.Println(content)
		//d.Content = content

		fmt.Println(e.Attr(`href`), e.Text)
		fmt.Println("*********")
	})
	//更多主题列表
	c.OnHTML(`div.repository-content > div > div.paginate-container > div > a`, func(e *query.HTMLElement) {

		fmt.Println(e.Attr(`href`), e.Text)
		fmt.Println("==========")
	})

	// c.OnScraped(func(_ *spider.Response) {
	// 	bData, _ := json.MarshalIndent(d, "", "\t")
	// 	out = string(bData)
	// })

	err := c.Visit(url)
	if err != nil {
		return nil
	}

	return nil
}

//爬取某个特定issue
func getIssueComment(url string) *IssueComment {
	d := IssueComment{}
	scraped := false

	// var bodyData []byte
	// Instantiate default collector
	c := spider.NewCollector()

	// c.OnResponse(func(r *spider.Response) {
	// 	bodyData = r.Body
	// })

	//主题
	c.OnHTML(`#partial-discussion-header > div.gh-header-show > h1 > span.js-issue-title`, func(e *query.HTMLElement) {
		d.Title = strings.TrimSpace(strings.TrimSpace(e.Text))
		//fmt.Println(d.Title)
		scraped = true
	})
	//时间
	c.OnHTML(`#partial-discussion-header > div.TableObject.gh-header-meta > div.TableObject-item.TableObject-item--primary > relative-time`, func(e *query.HTMLElement) {
		d.Time = strings.TrimSpace(strings.TrimSpace(e.Text))
		//fmt.Println(e.Text)
	})

	//评论内容
	c.OnHTML(`div.edit-comment-hide > task-lists > table > tbody > tr > td`, func(e *query.HTMLElement) {
		con := Content{Text: strings.TrimSpace(strings.TrimSpace(e.Text))}
		d.Contents = append(d.Contents, con)
		// fmt.Println(strings.TrimSpace(e.Text))
		// fmt.Println("------------------------------------------")
	})

	// c.OnScraped(func(_ *spider.Response) {
	// 	bData, _ := json.MarshalIndent(d, "", "\t")
	// 	out = string(bData)
	// })

	err := c.Visit(url)
	if err != nil {
		return nil
	}

	if scraped {
		return &d
	}

	return nil
}
