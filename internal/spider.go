package internal

import (
	"sort"
	"strings"
	"time"
)

func Spider(webPath string, urls map[string]string) {
	var targetUrls []string
	for targetUrl, _ := range urls {
		targetUrls = append(targetUrls, targetUrl)
	}
	sort.Strings(targetUrls)

	for order, targetUrl := range targetUrls {
		SLogger.Info(" 次序：", order+1, "/", len(targetUrls), " ----- ", urls[targetUrl], " ----- ", targetUrl)
		job(webPath, targetUrl)
	}
}

func job(webPath, targetUrl string) {
	tn := time.Now()
	switch {
	case strings.Contains(targetUrl, webPath+"/t"): // find all
		urls := FindDocument(webPath, targetUrl, "t")
		for i, val := range urls {
			SLogger.Info("[", i+1, "/", len(urls), "]")
			download(val)
		}
	case strings.Contains(targetUrl, webPath+"/x"): // find all list
		urls := FindDocument(webPath, targetUrl, "x")
		for i, val := range urls {
			SLogger.Info("[", i+1, "/", len(urls), "]")
			download(val)
		}
	case strings.Contains(targetUrl, webPath+"/a"): // find one
		download(targetUrl)
	}
	SLogger.Info("done...", time.Now().Sub(tn).String())
}

func download(url string) {
	document := NewDocument(url)
	if document != nil {
		name := document.Name()
		SLogger.Info("出境模特：", name)
		document.FindAll()
		document.SaveContents(url)
	}
}
