package main

import (
	"net/url"
	"fmt"
	"log"
)

type Link struct {
	parent *Link
	seen bool
	deep int
	url string
}

func NewLink (url string, parentLink *Link) *Link {
	deep := 0
	if parentLink != nil {
		deep = parentLink.deep + 1
	}
	return &Link{
		parent:parentLink,
		url:url,
		deep:deep,
		seen:false,
	}
}

func (this *Link)UpdateURLByParent() (err error) {
	u, err := url.Parse(this.url)
	if err != nil {
		return fmt.Errorf("url parse error: %q", this.url)
	}
	u.Fragment = ""
	if u.Scheme == "" || u.Host == "" {
		if this.parent == nil {
			return fmt.Errorf("url has not scheme or host: %q", this.url)
		}
		p, err := url.Parse(this.parent.url)
		if err != nil {
			return fmt.Errorf("url parse error: %q", this.parent.url)
		}
		if u.Scheme == "" {
			u.Scheme = p.Scheme
		}
		if u.Host == "" {
			u.Host = p.Host
		}
		if u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("url has not scheme or host: %q", this.url)
		}
	}
	this.url = u.String()
	return nil
}


func (this *Link) HasValidUrl() bool {
	u, err := url.Parse(this.url)
	if err != nil {
		log.Println("url parse error: ", err)
		return false
	}
	return u.Scheme != "" && u.Host != ""
}