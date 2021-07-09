package internal

import (
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type document struct {
	doc *goquery.Document

	url         string
	name        string // t5774__抖娘-利世
	imageset    string // _a_43313_
	filepath    string // /Users/coffee/Downloads/com.tujigu/download/
	downloaddir string // /Users/coffee/Downloads/com.tujigu/download/t5774__抖娘-利世

	content map[string]content //key=src
	pages   map[string]pages   //key=page
}

type content struct {
	src string
	alt string
}

type pages struct {
	page string
	href string
}

func NewDocument(urlStr string) *document {
	urlObj, err := url.Parse(urlStr)
	if urlStr == "" || err != nil {
		SLogger.Info(urlStr, err)
		return nil
	}

	doc := newDocument(urlObj.String())
	if doc == nil {
		return nil
	}
	return &document{
		filepath: FilePath,
		doc:      doc,
		url:      urlObj.String(),
		content:  map[string]content{},
		pages:    map[string]pages{},
	}
}

func FindDocument(webPath, url string, feature string) (a []string) {

	var urls = map[string]bool{url: true}
	searched := map[string]bool{}
	for !searched[url] {
		searched[url] = true

		if doc := newDocument(url); doc != nil {
			// Find all page
			doc.Find("center").Each(func(i int, s *goquery.Selection) {
				s.Find("a").Each(func(i int, s *goquery.Selection) {
					href, _ := s.Attr("href")
					if strings.Contains(href, feature+"/") {
						if !strings.Contains(href, "http") {
							href = webPath + href
						}
						urls[href] = true
					}
					class, _ := s.Attr("class")
					if strings.Contains(class, "next") {
						if !strings.Contains(href, "http") {
							href = webPath + href
						}
						url = href
						SLogger.Infoln(href)
					}
				})
			})
		}
	}

	var documents = map[string]bool{}
	for k, _ := range urls {
		doc := newDocument(k)
		if doc == nil {
			continue
		}
		doc.Find("div").Each(func(i int, s *goquery.Selection) {
			if val, exist := s.Attr("class"); !exist || val != "hezi" {
				return
			}
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				href, _ := s.Attr("href")
				if strings.Contains(href, "a") {
					documents[href] = true
				}
			})
		})
	}

	var ret []string
	for k, _ := range documents {
		ret = append(ret, k)
	}
	return ret
}

// 请求html页面
func newDocument(url string) (document *goquery.Document) {
	body, cancel := HttpReader(url)
	if body == nil {
		return
	}
	defer cancel()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		SLogger.Info(err)
		return
	}
	return doc
}

func (this *document) Name() string {
	name := ""
	if doc := this.doc; doc != nil {
		doc.Find("div").Each(func(i int, s *goquery.Selection) {
			if val, exist := s.Attr("class"); exist && val == "tuji" {
				s.Find("a").Each(func(i int, s *goquery.Selection) {
					href, _ := s.Attr("href")
					if strings.Contains(href, "/t/") {
						t := strings.Split(strings.Replace(strings.Replace(href[strings.Index(href, "/t/"):], " ", "", -1), "/", "", -1), "、")[0]
						if name == "" {
							name = t + "__" + s.Text()
						}
					}
				})
			}
		})
	}
	if len(name) == 0 {
		name = "default"
	}
	this.name = name
	return name
}

func (this *document) FindAll() {
	for this.url != "" {
		this.Find()

		if next := this.Next(); next == "" || next == this.url {
			this.url = ""
		} else {
			this.url = next
		}
	}
}

func (this *document) Find() {
	if doc := this.doc; doc != nil {
		this.FindContents(doc)
		this.FindPages(doc)
	}
}

func (this *document) Next() string {
	if len(this.pages) == 0 {
		return ""
	}
	page, ok := this.pages["下一页"]
	if !ok {
		return ""
	}
	return page.href
}

func (this *document) FindContents(document *goquery.Document) {
	if this.content == nil {
		this.content = make(map[string]content)
	}

	// Find the review items
	document.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if val, exist := s.Attr("class"); exist && val == "content" {
			s.Find("img").Each(func(i int, s *goquery.Selection) {
				src, _ := s.Attr("src")
				alt, _ := s.Attr("alt")
				class, _ := s.Attr("class")
				if class == "tupian_img" {
					this.content[src] = content{
						src: src,
						alt: strings.Replace(strings.Replace(alt, " ", "", -1), "/", "", -1),
					}
				}
			})
		}
	})
}

func (this *document) FindPages(document *goquery.Document) {
	if this.pages == nil {
		this.pages = make(map[string]pages)
	}

	// Find the review items
	document.Find("div").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if val, exist := s.Attr("id"); !exist || val != "pages" {
			return
		}
		s.Find("a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			page := s.Text()
			this.pages[page] = pages{
				page: page,
				href: href,
			}
		})
	})
}

func (this *document) SaveContents(urlStr string) {
	urlObj, _ := url.Parse(urlStr)
	this.imageset = strings.Replace(urlObj.EscapedPath(), "/", "_", -1)
	this.downloaddir = path.Join(this.filepath, this.name)
	SLogger.Info("document = ", urlStr, this.downloaddir, this.imageset)

	i := 0
	for _, content := range this.content {
		SLogger.Info("content: [", i, "/", len(this.content), "], download...")
		i++
		HttpDownload(content.src, this.downloaddir, this.name+this.imageset+content.alt)
	}
}
