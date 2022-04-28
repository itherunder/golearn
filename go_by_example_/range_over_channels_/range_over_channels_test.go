package range_over_channels_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestRangeOverChannels(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		colorlog.Info(elem)
	}
}
