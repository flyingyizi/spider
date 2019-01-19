package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/flyingyizi/spider"
	//或使用下面的库
	//spider ""github.com/gocolly/colly"
)

type ChannelItem struct {
	Title    string `json:"title"`
	OpenUrl  string `json:"open_url"`
	GroupID  string `json:"group_id"`
	ImageUrl string `json:"image_url"`
}
type ToutiaoChannel struct {
	RealTimeNews []ChannelItem `json:"real_time_news"`
}

// Save saves model in the file
func (f *ToutiaoChannel) Save(path string) error {

	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(f)
	}
	file.Close()
	return err
}

func main() {
	//startURL := "https://www.toutiao.com/ch/news_entertainment/"
	startURL := "https://www.toutiao.com/ch/news_tech/"

	urls := spider.UrlFlags(startURL)

	start := time.Now()
	for i, url := range urls {
		out := getToutiaoChannel(url)
		fmt.Printf("------output-----%d url----------\n", i)
		fmt.Println(out)

	}
	elapsed := time.Since(start)
	fmt.Printf("Time required to complete: %s\n", elapsed)

}

func getToutiaoChannel(url string) (out string) {
	d := ToutiaoChannel{}

	var (
		bodyData []byte
	)

	// Instantiate default collector
	c := spider.NewCollector()

	c.OnResponse(func(r *spider.Response) {
		bodyData = r.Body

		//var _data = {"real_time_news"
		//通过观察头条频道的网页，频道文章网址存放在js代码中，下面从js代码中将感兴趣的内容解析出来
		htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyData))
		if err != nil {
			return
		}
		scripts := htmlDoc.Find("script")
		destScript := ""
		key1, key2 := "_data", "real_time_news"
		for _, n := range scripts.Nodes {
			if n.FirstChild != nil {
				//找了存放内容js代码的特征
				if strings.Contains(n.FirstChild.Data, key1) && strings.Contains(n.FirstChild.Data, key2) {
					destScript = n.FirstChild.Data
					break
				}
			}
		}
		//解析出js代码中 "_data"变量的值
		startIndex := strings.Index(destScript, key1)
		jsonData := getContentInBrace(destScript, startIndex)

		json.Unmarshal(jsonData, &d)
		for i := 0; i < len(d.RealTimeNews); i++ {
			d.RealTimeNews[i].OpenUrl = `https://www.toutiao.com` + d.RealTimeNews[i].OpenUrl
		}
		// fmt.Println(d.RealTimeNews)

	})

	//xpath: `//head/title`
	// c.OnHTML(`head>title`, func(e *spider.HTMLElement) {
	// 	d.Title = strings.TrimSpace(e.Text)
	// 	//fmt.Println(d.Title)
	// })

	c.OnScraped(func(_ *spider.Response) {
		bData, _ := json.MarshalIndent(d, "", "\t")
		out = string(bData)
	})

	err := c.Visit(url)
	if err != nil {
		return
	}
	return
	// d.OrigHTML = base64.StdEncoding.EncodeToString(bodyData)
	//d.Save("abc.json")

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

//将中文转换为unicode码，使用golang中的strconv包中的QuoteToASCII直接进行转换，
//将unicode码转换为中文就比较麻烦一点，先对unicode编码按\u进行分割，然后使用strconv.ParseInt，
//将16进制数字转换Int64，在使用fmt.Sprintf将数字转换为字符，最后将其连接在一起，
//这样就变成了中文字符串了。 参考代码如下：
func unicode2ZH(sText string) {
	sText = "\u8fd9\u4e2a\u4e16\u754c\u4e0a\u6bcf\u4e2a\u4eba\u90fd\u5f88\u82e6\uff0c\u51ed\u4ec0\u4e48\u53ea\u6709\u4f60\u625b\u4e0d\u4f4f\uff1f"
	textQuoted := strconv.QuoteToASCII(sText)
	textUnquoted := textQuoted[1 : len(textQuoted)-1]
	fmt.Println(textUnquoted)

	sUnicodev := strings.Split(textUnquoted, "\\u")
	var context string
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			panic(err)
		}
		context += fmt.Sprintf("%c", temp)
	}
	fmt.Println(context)
}
