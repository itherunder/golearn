package sha256_hashes_

import (
	"crypto/sha256"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestSha256Hashes(t *testing.T) {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	colorlog.Info("%v", s)
	colorlog.Info("%x", bs)
}
