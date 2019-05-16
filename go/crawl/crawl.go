package main

import (
	"golang.org/x/net/html"
	"net/http"
	"log"
	"flag"
	"sync"
)

func extractLinks(links *[]string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*links = append(*links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLinks(links, c)
	}
}

var tokens = make(chan struct{}, 30)

func crawl(linkInfo *Link) []string {
	if !linkInfo.HasValidUrl() {
		return []string{}
	}
	if _, ok := linkFail.Load(linkInfo.url); ok {
		return []string{}
	}
	log.Println("crawl:", linkInfo.url, linkInfo.deep)

	tokens <- struct{}{}
	response, err := http.Get(linkInfo.url)
	<- tokens

	if err != nil {
		linkFail.Store(linkInfo.url, true)
		log.Println("http 请求失败", linkInfo.url, " parent Url:", linkInfo.parent.url)
		return []string{}
	}
	defer response.Body.Close()


	doc, err := html.Parse(response.Body)
	if err != nil {
		linkFail.Store(linkInfo.url, true)
		log.Println("html 解析失败", response.Body, " Url:", linkInfo.url)
		return []string{}
	}
	go handle(doc)
	links := []string{}
	extractLinks(&links, doc)
	return links
}

func handle(node *html.Node) {
	// do something
}

var (
	linkSeen = sync.Map{}  // 对已爬过的去重
	linkFail = sync.Map{}  // 对爬取失败的去重
	worklist = make(chan []string, 5)
	maxDeep = 0
)

func initCrawl() {
	linkURL := ""
	flag.StringVar(&linkURL, "url", "", "the url that you want to crawl")
	flag.IntVar(&maxDeep, "deep", 3, "the depth of crawl")
	flag.Parse()
	link := NewLink(linkURL, nil)
	if !link.HasValidUrl() {
		log.Fatal("请输入正确的 url 参数")
		return
	}
	worklist <- []string{link.url}
}

func main() {
	initCrawl()
	var linkInfo *Link

	for count:=1; count>0; count--{
		list := <- worklist
		for _, linkUrl := range list {
			v, ok := linkSeen.Load(linkUrl)
			if !ok {
				linkInfo = NewLink(linkUrl, nil)
				linkSeen.Store(linkUrl, linkInfo)
			} else {
				linkInfo = v.(*Link)
			}
			if !linkInfo.seen && linkInfo.deep < maxDeep {
				linkInfo.seen = true
				count++
				go func(linkInfo *Link) {
					urls := crawl(linkInfo)
					list := []string{}
					for _,l := range urls {
						t := NewLink(l, linkInfo)
						t.UpdateURLByParent()
						if _, ok := linkSeen.Load(t.url); !ok {
							linkSeen.Store(t.url, t)
							list = append(list, t.url)
						}
					}
					worklist <- list
				}(linkInfo)
			}
		}
	}
}