package sha256_

import (
	"crypto/sha256"
	"testing"

	"github.com/yezihack/colorlog"
)

var pc [256]int

func init() {
	for i := 0; i < 256; i++ {
		pc[i] = i%2 + pc[i/2]
	}
}

func TestSha256(t *testing.T) {
	sig1 := sha256.Sum256([]byte("x"))
	sig2 := sha256.Sum256([]byte("X"))
	colorlog.Info("sig1 == sig2 ? %v %T\n", sig1 == sig2, sig1)
	colorlog.Info("\nsig1 = %x\nsig2 = %x\n", sig1, sig2)
	colorlog.Info("sig1 1 cnt: %d", PopCountOfSha256(&sig1))
	colorlog.Info("sig2 1 cnt: %d", PopCountOfSha256(&sig2))
}

func PopCountOfSha256(sig *[32]uint8) int {
	cnt := 0
	for _, x := range sig {
		cnt += pc[x]
	}
	return cnt
}

func TestSha(t *testing.T) {
	Sha("SHA384")
}
