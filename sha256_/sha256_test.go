package sha256_

import (
	"crypto/sha256"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestSha256(t *testing.T) {
	sig1 := sha256.Sum256([]byte("x"))
	sig2 := sha256.Sum256([]byte("X"))
	colorlog.Info("sig1 == sig2 ? %v %T\n", sig1 == sig2, sig1)
	colorlog.Info("\nsig1 = %x\nsig2 = %x\n", sig1, sig2)
}
