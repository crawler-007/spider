package internal

import (
	"testing"
)

func TestSpider(t *testing.T) {
	//Spider(FilePath, WebPath, AllDoc)
	Spider(WebPath, map[string]string{
		"https://www.tujigu.com/t/5774/": "抖娘-利世",
	})
}
