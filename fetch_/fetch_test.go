package fetch_

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestFetch(t *testing.T) {
	urls := []string{"http://www.baidu.com", "https://github.com/itherunder"}
	for _, url := range urls {
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
		colorlog.Info("%s: %s", url, b)
	}
}
