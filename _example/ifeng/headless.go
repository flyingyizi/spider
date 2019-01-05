package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/client"
)

func main() {
	start := time.Now()
	startURL := "https://m.ifeng.com/huaweillqtest?appid=hwbrowser&ch=ref_hwllq_dl1&type=doc&aid=ucms_7j9jR0XW9MD&ctype=news"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cdp, err := HeadlesswithStandlone(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	var html string
	err = cdp.Run(ctx, GetHtmlAction(startURL, &html))
	if err != nil {
		// 错误处理
		return
	}
	// 成功取得HTML内容进行后续处理
	//fmt.Println(html)
	analyzeIfeng([]byte(html))

	//关闭chrome实例
	err = cdp.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 等待chrome实例关闭
	err = cdp.Wait()
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Time required to complete: %s\n", elapsed)

}

func GetHtmlAction(url string, res *string) chromedp.Tasks {
	return chromedp.Tasks{
		// 访问页面
		chromedp.Navigate(url),
		// 等待table渲染成功，成功则说明已经获取到了正确的页面
		chromedp.WaitVisible(`div.mod_news_text`, chromedp.ByQuery),
		// 获取 标签的html字符
		chromedp.OuterHTML("div.mod_news_text", res),
	}
}

// HeadlesswithStandlone 通过独立chrome（standlone chrome），比如已经启动的chrome docker
func HeadlesswithStandlone(ctxt context.Context) (c *chromedp.CDP, err error) {
	logger := func(string, ...interface{}) {
		return
	}
	// create chrome instance
	c, err = chromedp.New(ctxt, chromedp.WithTargets(client.New().WatchPageTargets(ctxt)), chromedp.WithLog(logger /* log.Printf */))
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func analyzeIfeng(bodyData []byte) {
	htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyData))
	if err != nil {
		return
	}

	ps := htmlDoc.Find(".visible")

	//通过观察ifeng的网页，内容是js动态生成，从js代码无法将感兴趣的内容解析出来
	for _, n := range ps.Nodes {
		doc := goquery.NewDocumentFromNode(n)
		val, exists := doc.Find("img").Attr("data-src")

		if exists {
			fmt.Println(val)
		} else {
			fmt.Println(doc.Text())
		}

	}

	//fmt.Println(destScript)

	//解析出js代码中xissJsonData变量的值
	// startIndex := strings.Index(destScript, key1)
	// xissJsonData := getContentInBrace(destScript, startIndex)
	//fmt.Println(string(xissJsonData))

	type XissJSONData struct {
		Content string `json:"content"`
	}
	// if err := json.Unmarshal(xissJsonData, &d); err == nil {
	// 	//fmt.Println(d)
	// } else {
	// 	fmt.Println(err)
	// 	return
	// }
	//d.OrigHTML = base64.StdEncoding.EncodeToString(bodyData)

	//d.Save("efg.json")

}

// // Headless 通过自身创建headless chrome实例,
// // 从实践看，该方式有问题， chrome 实例无法关闭。
// func Headless(ctx context.Context) (c *chromedp.CDP, err error) {
// 	logger := func(string, ...interface{}) {
// 		return
// 	}

// 	c, err = chromedp.New(ctx, chromedp.WithRunnerOptions(
// 		//runner.Headless(path, 9222),
// 		runner.Flag("headless", true),
// 		runner.Flag("disable-web-security", true),
// 		runner.Flag("no-first-run", true),
// 		runner.Flag("no-default-browser-check", true),
// 		runner.Flag("disable-gpu", true),
// 	), chromedp.WithLog(logger)) //, cdp.WithLog(log.Printf))

// 	// // Windows用户需要设置runner.Flag("disable-gpu", true)，具体信息参见文档的FAQ
// 	// run, err = runner.New(runner.Flag("headless", true), runner.Flag("disable-gpu", true) /* ,        runner.URL(starturl) */)
// 	// if err != nil {
// 	// 	return nil, nil, err
// 	// }

// 	// // run.Start启动实例
// 	// err = run.Start(ctx)
// 	// if err != nil {
// 	// 	return nil, nil, err
// 	// }

// 	// c, err = chromedp.New(ctx, chromedp.WithRunner(run), chromedp.WithErrorf(logger /* log.Printf */))
// 	// if err != nil {
// 	// 	return
// 	// }

// 	return
// }
