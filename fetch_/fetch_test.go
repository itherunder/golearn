package fetch_

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

var urls = []string{"http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com", "http://www.baidu.com"}

func fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		colorlog.Error("fetch %s failed: %v", url, err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		colorlog.Error("read %s failed: %v", url, err)
		os.Exit(1)
	}
	if ch != nil {
		ch <- string(b)
	} else {
		colorlog.Info("%s: %d", url, len(string(b)))
	}
}

func TestFetch(t *testing.T) {
	start := time.Now()
	for _, url := range urls {
		fetch(url, nil)
	}
	end := time.Now() // 2.85s
	colorlog.Info("fetch all url cost %v", end.Sub(start))
}

func TestFetchAll(t *testing.T) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for _, url := range urls {
		s := <-ch
		colorlog.Info("%s: %d", url, len(s))
	}
	end := time.Now() // 1.76s
	colorlog.Info("fetch all url cost %v", end.Sub(start))
}
