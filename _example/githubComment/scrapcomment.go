//导出github某个archive的issue
//example: go run scrapcomment.go -o "abc.md"  "https://github.com/gonum/plot/issues"

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flyingyizi/spider"
)

func main() {
	outfile := flag.String("o", "", "file stores the logs, default is outputing to terminal")

	startURL := "https://github.com/fxsjy/jieba/issues"
	urls := spider.UrlFlags(startURL)
	var wg sync.WaitGroup
	pageUrls := make([]string, 100)
	start := time.Now()
	//step 1:
	for _, issueURL := range urls {
		ps := producerContainer(issueURL)
		pageUrls = append(pageUrls, ps...)
	}

	//step 2
	topicURLs := make(chan string, 0)
	var wg2 sync.WaitGroup
	var out sync.Map
	for _, pageURL := range pageUrls {
		wg.Add(1)
		go getIssueTopicList(&wg, pageURL, topicURLs)

		wg2.Add(1)
		go getIssueComment(&wg2, topicURLs, &out)
	}
	wg.Wait()
	close(topicURLs)
	wg2.Wait()

	result := make(CommentSlice, 0)
	out.Range(func(key, value interface{}) bool {
		tmpval, valid := key.(*IssueComment)
		if !valid {
			return true //返回true，则range下一个key
		}
		result = append(result, tmpval)
		return true
	})
	result.Sort()

	md := result.ToMarkDown(*outfile)
	if *outfile == "" {
		fmt.Println(md)
	}
	elapsed := time.Since(start)
	fmt.Printf("total %d, Time required to complete: %s\n", len(result), elapsed)
}

//通过导航获取该issue的所有page list,
// 获取到的page url 放到pageUrls chan
func producerContainer(issueURL string) (pageUrls []string) {
	c := spider.NewCollector()

	exp := regexp.MustCompile(`page=\d+`)
	base, maxPage := "", 0
	//更多主题列表
	c.OnHTML(`div.repository-content > div > div.paginate-container > div > a`, func(e *spider.HTMLElement) {
		if val, err := strconv.Atoi(e.Text); err == nil {
			if val > maxPage {
				maxPage = val
				base = e.Attr(`href`)
			}
		}
		// fmt.Printf("==========%s%s\n", e.Attr(`href`), e.Text)
		// fmt.Println("==================")
	})
	//
	c.OnScraped(func(_ *spider.Response) {
		first, second := "", ""
		if loc := exp.FindIndex([]byte(base)); len(loc) == 2 {
			first, second = base[0:loc[0]], base[loc[1]:]
		}
		pageUrls = make([]string, 0)
		for i := 1; i <= maxPage; i++ {
			p := fmt.Sprintf("https://github.com%spage=%d%s", first, i, second)
			pageUrls = append(pageUrls, p)
		}
	})

	err := c.Visit(issueURL)
	if err != nil {
		return
	}
	return
}

//爬取issue topic url list，以便后续爬取评论内容
func getIssueTopicList(wg *sync.WaitGroup, pageURL string, topicURLs chan<- string) {
	defer wg.Done()
	c := spider.NewCollector()

	//register 获取主题列表
	c.OnHTML(`div.repository-content > div > div.border-right.border-bottom.border-left > div > div > div > div >a `, func(e *spider.HTMLElement) {
		topicURLs <- "https://github.com" + e.Attr(`href`)
		// fmt.Println(e.Attr(`href`), e.Text)
		// fmt.Println("*********")
	})

	// c.OnScraped(func(_ *spider.Response) {
	// })

	err := c.Visit(pageURL)
	if err != nil {
		return
	}
	return
}

//爬取某个特定issue
func getIssueComment(wg *sync.WaitGroup, topicURLs <-chan string, out *sync.Map) {
	defer wg.Done()

	d := IssueComment{}
	c := spider.NewCollector()

	//主题
	c.OnHTML(`#partial-discussion-header > div.gh-header-show > h1 > span.js-issue-title`, func(e *spider.HTMLElement) {
		d.Title = strings.TrimSpace(strings.TrimSpace(e.Text))
	})
	//时间
	c.OnHTML(`#partial-discussion-header > div.TableObject.gh-header-meta > div.TableObject-item.TableObject-item--primary > relative-time`, func(e *spider.HTMLElement) {
		d.Time = strings.TrimSpace(strings.TrimSpace(e.Text))
		//fmt.Println(e.Text)
	})

	//评论内容
	c.OnHTML(`div.edit-comment-hide > task-lists > table > tbody > tr > td`, func(e *spider.HTMLElement) {
		d.Contents = append(d.Contents, strings.TrimSpace(strings.TrimSpace(e.Text)))
		// fmt.Println(strings.TrimSpace(e.Text))
		// fmt.Println("------------------------------------------")
	})

	c.OnScraped(func(r *spider.Response) {
		copyd := new(IssueComment)
		*copyd = d
		out.Store(copyd, r.Request.URL)
	})
	for topicURL := range topicURLs {
		err := c.Visit(topicURL)
		if err != nil {
			//return
		}
	}

	return
}

// CommentSlice attaches the methods of Interface to []string, sorting in increasing order.
type CommentSlice []*IssueComment

func (p CommentSlice) Len() int      { return len(p) }
func (p CommentSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p CommentSlice) Less(i, j int) bool {
	layout1, layout2 := `Jan 2, 2006`, `Jan 02, 2006`
	var (
		tm1, tm2   time.Time
		err1, err2 error
	)

	first, second := p[i].Time, p[j].Time
	if tm1, err1 = time.Parse(layout1, first); err1 != nil {
		tm1, err1 = time.Parse(layout2, first)
	}
	if tm2, err2 = time.Parse(layout1, second); err1 != nil {
		tm2, err2 = time.Parse(layout2, second)
	}
	//如果出现这个panic，那说明layout不对
	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	return tm1.Unix() > tm2.Unix()
}

// Sort is a convenience method.
func (p CommentSlice) Sort() { sort.Sort(p) }

func (p CommentSlice) ToMarkDown(path string) (out string) {
	out = ""
	for i, j := range p {
		out = out + fmt.Sprintf("## %d.%s\n%s\n", i, j.Title, j.Time)

		for _, k := range j.Contents {
			out = out + fmt.Sprintf("```text\n %s", k)
			out = out + "\n```\n"
		}
	}

	if path != "" {
		// 可写方式打开文件
		file, err := os.OpenFile(
			path,
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
		)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		// 写字节到文件中
		file.Write([]byte(out))
	}
	return
}

//IssueComment 存放内容的卡片
type IssueComment struct {
	Time     string   `json:"time"`
	Title    string   `json:"title"` //,omitempty在序列化的时候忽略0值或者空值
	Contents []string `json:"contents"`
}

// // Save saves model in the file
// func (f *IssueComment) Save(path string) error {

// 	file, err := os.Create(path)
// 	if err == nil {
// 		encoder := json.NewEncoder(file)
// 		err = encoder.Encode(f)
// 	}
// 	file.Close()
// 	return err
// }
