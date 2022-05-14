package epoch_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestEpoch(t *testing.T) {
	now := time.Now()
	colorlog.Info(now.String())

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1_000_000

	colorlog.Info("secs: %d\nnanos: %d\nmillis: %d", secs, nanos, millis)

	colorlog.Info("%+v, %+v", time.Unix(secs, 0), time.Unix(0, nanos))
}
