package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/flyingyizi/spider"
	"github.com/flyingyizi/spider/query"
)

type WeiboRealTimeHot struct {
	Top     int    `json:"top"`
	KeyWord string `json:"keyWord"`
	Mount   string `json:"mount"`
	Tag     string `json:"tag"`
}

func main() {
	//weibo热搜榜
	realtimehotURL := "https://s.weibo.com/top/summary?cate=realtimehot"
	GetweiborealTimeHotUC(realtimehotURL)
}

func GetweiborealTimeHotUC(url string) {

	// Instantiate default collector
	c := spider.NewCollector()

	//热搜榜表格正文,表格片段如下
	// 	<table>
	//     <thead>
	//         <tr class="thead_tr">
	//             <th class="th-01">序号</th>
	//             <th class="th-02">关键词</th>
	//             <th class="th-03"></th>
	//         </tr>
	//     </thead>
	//     <tbody>
	//         <tr class="">
	//             <td class="td-01"><i class="icon-top"></i></td>
	//             <td class="td-02">
	//                 <a href="/weibo?q=%E9%99%86%E5%86%9B%E5%AE%A3%E4%BC%A0%E7%89%87%20%E5%9C%B0%E8%A1%A8%E6%9C%80%E5%BC%BA&amp;Refer=new_time"
	//                     target="_blank">陆军宣传片 地表最强</a>
	//                 <span>2347804</span>
	//             </td>
	//             <td class="td-03"><i class="icon-txt icon-txt-hot">热</i></td>
	//         </tr>
	//     </tbody>
	// </table>

	dlist := make([]*WeiboRealTimeHot, 0)
	c.OnHTML(`#pl_top_realtimehot > table > tbody`, func(e *query.HTMLElement) {
		e.ForEach("tr", func(_ int, trSel *spider.HTMLElement) {
			d := WeiboRealTimeHot{}
			trSel.ForEach("td", func(j int, tbSel *spider.HTMLElement) {
				switch j {
				case 0:
					d.Top, _ = strconv.Atoi(strings.TrimSpace(tbSel.Text))
				case 1:
					d.KeyWord = strings.TrimSpace(tbSel.ChildText("a"))
					d.Mount = strings.TrimSpace(tbSel.ChildText("span"))
				case 2:
					d.Tag = strings.TrimSpace(tbSel.Text)
				}
			})
			dlist = append(dlist, &d)
		})
	})

	c.OnScraped(func(_ *spider.Response) {
		bData, _ := json.MarshalIndent(dlist, "", "\t")
		fmt.Println(string(bData))
	})

	err := c.Visit(url)
	if err != nil {
		return
	}
}
