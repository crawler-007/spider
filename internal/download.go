package internal

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func HttpBody(uri string) []byte {
	body, cancel := HttpReader(uri)
	if body == nil {
		return nil
	}
	defer cancel()

	data, _ := ioutil.ReadAll(body)
	return data
}

func HttpReader(uri string) (io.ReadCloser, func()) {
	var (
		resp *http.Response
		err  error
	)

	for retry := 3; retry > 0; retry-- {
		resp, err = http.Get(uri)
		if err == nil {
			break
		}
	}
	if err != nil {
		SLogger.Info("status code error:", err)
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		SLogger.Info("status code error:", resp.StatusCode, resp.Status)
		return nil, nil
	}

	return resp.Body, func() {
		resp.Body.Close()
	}
}

var saved = map[string]bool{}
var savedLogPath = path.Join(LogPatg, "saved.txt")

func isSaved(s string) bool {
	if len(saved) == 0 {
		savedinfo, _ := ioutil.ReadFile(savedLogPath)
		for _, line := range strings.Split(string(savedinfo), "\n") {
			if item := strings.Split(line, "##"); len(item) > 0 && len(line) > 0 {
				saved[item[0]] = true
			}
		}
	}
	if saved[s] {
		SLogger.Info("saved:")
		return true
	}
	return false
}

func updateSaved(s, uri string) {
	if saved[s] {
		return
	}
	write(savedLogPath, []byte(s+"##"+uri))
}

func checkPath(filepath string) {
	_, err := os.Stat(filepath)
	if err == nil {
		return
	}
	os.Mkdir(filepath, os.ModePerm)
}

func write(filename string, body []byte) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		SLogger.Info(err)
		return
	}
	defer f.Close()
	f.Write(body)
}

func HttpDownload(uri string, filepath, filename string) {
	fname := path.Join(filepath, filename+".jpg")
	if isSaved(fname) {
		return
	}

	body := HttpBody(uri)
	if len(body) == 0 {
		return
	}

	checkPath(filepath)

	write(fname, body)

	updateSaved(fname, uri)
}
