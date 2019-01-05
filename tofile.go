package spider

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// "content":"<p>自。</p><p>1、</p><p>舰。</p><!--{img: 0}--><p>2、</p><!--{img: 1 }--><div id=\"hidtag\"></div><p>3、</p><!--{img: 2}--><p>4、</p><p>！</p><!--{img: 3}--><p>5、</p>",
// "thumbnails":[
//     {"url":"http://image.uczzd.cn/2563535540031726725.jpg?id=0&from=export","width":640,"height":360,"type":"jpg","preload":0,"daoliu_url":"","daoliu_title":""},
//     {"url":"http://image.uczzd.cn/11790983737676000867.jpg?id=0&from=export","width":640,"height":360,"type":"jpg","preload":0,"daoliu_url":"","daoliu_title":""},
//     {"url":"http://image.uczzd.cn/18224486047257093302.jpg?id=0&from=export","width":615,"height":346,"type":"jpg","preload":0,"daoliu_url":"","daoliu_title":""}],
// "images":[
//     {"title":"","url":"http://image.uczzd.cn/14043054703237486508.jpg?id=0&from=export","description":"","index":0,"width":640,"height":360,"preload":1,"type":"jpg","focus":"50_50","original_url":"http://image.uc.cn/s/wemedia/s/upload/2018/b39193abd84b820fb3af1e732581cd5e.jpg","daoliu_url":"","daoliu_title":"","gallery_id":"","gallery_type":0,"is_hd":0,"image_extra":{}},
//     {"title":"","url":"http://image.uczzd.cn/17372319831213558599.jpg?id=0&from=export","description":"","index":1,"width":640,"height":426,"preload":0,"type":"jpg","focus":"50_50","original_url":"","daoliu_url":"","daoliu_title":"","gallery_id":"","gallery_type":0,"is_hd":0,"image_extra":{}},
//     {"title":"","url":"http://image.uczzd.cn/1479968059362464066.jpg?id=0&from=export","description":"","index":2,"width":640,"height":394,"preload":0,"type":"jpg","focus":"50_50","original_url":"","daoliu_url":"","daoliu_title":"","gallery_id":"","gallery_type":0,"is_hd":0,"image_extra":{}},
//     {"title":"","url":"http://image.uczzd.cn/6819343685372118035.jpg?id=0&from=export","description":"","index":3,"width":640,"height":346,"preload":0,"type":"jpg","focus":"50_50","original_url":"","daoliu_url":"","daoliu_title":"","gallery_id":"","gallery_type":0,"is_hd":0,"image_extra":{}}],
type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
}
type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Type   string `json:"type"`
	Index  int    `json:"index"`
}

//DocCard 存放内容的卡片
type DocCard struct {
	Content    string      `json:"content"`
	Title      string      `json:"title"` //,omitempty在序列化的时候忽略0值或者空值
	Thumbnails []Thumbnail `json:"thumbnails"`
	Images     []Image     `json:"images"`
	OrigHTML   string      `json:"origHtml"`
}

// Save saves model in the file
func (f *DocCard) Save(path string) error {

	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(f)
	}
	file.Close()
	return err
}

// Load loads from the file
func (f *DocCard) Load(path string) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&f)
	}
	file.Close()

	return err
}

//UrlFlags 命令行输入网址，如果命令行不输入网址，则以initUrl作为命令行输入的网址
func UrlFlags(initUrl string) []string {
	usage := func() {
		fmt.Fprintf(os.Stderr, `default input paramenters are urls seprated by space
	`)
		flag.PrintDefaults()
	}

	flag.Usage = usage

	flag.Parse()
	others := flag.Args()
	if len(others) == 0 {
		others = append(others, initUrl)
	}
	return others
}
