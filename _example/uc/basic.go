package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/flyingyizi/spider"
)

// https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=3091260259678083253&cid=100&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=859544137256752847&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D
// https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=16119487478051582585&cid=179223212&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=7511051609959712598&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D

func main() {
	//startURL := "https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=3091260259678083253&cid=100&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=859544137256752847&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D"
	startURL2 := "https://m.uczzd.cn/webview/news?app=huawei-iflow&aid=16119487478051582585&cid=179223212&zzd_from=huawei-iflow&uc_param_str=dndsfrvesvntnwpfgibicpkt&recoid=7511051609959712598&rd_type=reco&sp_gz=1&enuid=AANA1zeV8aROM%2FjV3SPP4HAxwsLLuvuIpl56ol5ZzlXQuw%3D%3D"
	getUC(startURL2)
}

func getUC(url string) {
	d := spider.DocCard{}

	var bodyData []byte
	// Instantiate default collector
	c := spider.NewCollector()

	c.OnResponse(func(r *spider.Response) {
		bodyData = r.Body
	})

	//xpath: `//head/title`
	c.OnHTML(`head>title`, func(e *spider.HTMLElement) {
		d.Title = strings.TrimSpace(e.Text)
		//fmt.Println(d.Title)
	})

	err := c.Visit(url)
	if err != nil {
		return
	}

	htmlDoc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyData))
	if err != nil {
		return
	}

	scripts := htmlDoc.Find("script")

	destScript := ""
	//通过观察uc的网页，内容存放在js代码中，下面从js代码中将感兴趣的内容解析出来
	key1, key2 := "xissJsonData", "content"
	for _, n := range scripts.Nodes {
		if n.FirstChild != nil {
			//找了两个存放内容js代码的特征
			if strings.Contains(n.FirstChild.Data, key1) && strings.Contains(n.FirstChild.Data, key2) {
				destScript = n.FirstChild.Data
				break
			}
		}
	}
	//fmt.Println(destScript)

	//解析出js代码中xissJsonData变量的值
	startIndex := strings.Index(destScript, key1)
	xissJsonData := getContentInBrace(destScript, startIndex)
	//fmt.Println(string(xissJsonData))

	type XissJSONData struct {
		Content string `json:"content"`
	}
	if err := json.Unmarshal(xissJsonData, &d); err == nil {
		//fmt.Println(d)
	} else {
		fmt.Println(err)
		return
	}
	d.OrigHTML = base64.StdEncoding.EncodeToString(bodyData)

	d.Save("efg.json")

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
