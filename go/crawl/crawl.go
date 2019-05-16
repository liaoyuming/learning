package main

import (
	"golang.org/x/net/html"
	"net/http"
	"log"
	"net/url"
	"flag"
	"sync"
)

func extractLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = extractLinks(links, c)
	}
	return links
}

func isURL(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		log.Println("url parse error: ", err)
		return false
	}
	return u.Scheme != "" && u.Host != ""
}

var tokens = make(chan struct{}, 20)

func crawl(link string) []string {
	if !isURL(link) {
		return []string{}
	}
	log.Println("crawl:", link)
	tokens <- struct{}{}
	response, err := http.Get(link)
	<- tokens
	if err != nil {
		log.Println("http 请求失败", link)
		return []string{}
	}
	defer response.Body.Close()
	links := []string{}
	doc, err := html.Parse(response.Body)
	if err != nil {
		log.Println("html 解析失败", response.Body)
		return []string{}
	}
	links = extractLinks(links, doc)
	return links
}

type LinkInfo struct {
	seen bool
	deep int
}

func main() {
	link := ""
	maxDeep := 0
	flag.StringVar(&link, "url", "", "the url that you want to crawl")
	flag.IntVar(&maxDeep, "deep", 3, "the depth of crawl")
	flag.Parse()
	if !isURL(link) {
		log.Fatal("请输入正确的 url 参数")
		return
	}

	worklist := make(chan []string, 3)
	worklist <- []string{link}

	count := 1
	linkInfo := new(LinkInfo)

	linkInfoSyncMap := sync.Map{}

	for ; count >0; count--{
		list := <- worklist
		for _, linkUrl := range list {
			v, ok := linkInfoSyncMap.Load(linkUrl)
			if !ok {
				linkInfo = &LinkInfo{seen:false, deep:0}
				linkInfoSyncMap.Store(linkUrl, linkInfo)
			} else {
				linkInfo = v.(*LinkInfo)
			}
			if !linkInfo.seen && linkInfo.deep < maxDeep {
				linkInfo.seen = true
				count++
				go func(url string, linkInfo *LinkInfo) {
					urls := crawl(url)
					for _,l := range urls {
						linkInfoSyncMap.Store(l, &LinkInfo{
							seen:false,
							deep:linkInfo.deep+1,
						})
					}
					worklist <- urls
				}(linkUrl, linkInfo)
			}
		}
	}
}