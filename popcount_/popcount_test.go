package popcount_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}

func PopCount_(x uint64) int {
	ans := 0
	for i := 0; i < 8; i++ {
		ans += int(pc[byte(x>>(i*8))])
	}
	return ans
}

func PopCount__(x uint64) int {
	ans := 0
	for x != 0 {
		ans += int(x & 1)
		x >>= 1
	}
	return ans
}

func PopCount___(x uint64) int {
	ans := 0
	for x != 0 {
		ans++
		x &= (x - 1)
	}
	return ans
}

func TestPopCount(t *testing.T) {
	colorlog.Info("%d = %d", 0, PopCount(0))
	colorlog.Info("%d = %d", 1, PopCount(1))
	colorlog.Info("%d = %d", 0x1234567890ABCDEF, PopCount(0x1234567890ABCDEF))
}

func TestPopCount_(t *testing.T) {
	colorlog.Info("%d = %d", 0, PopCount_(0))
	colorlog.Info("%d = %d", 1, PopCount_(1))
	colorlog.Info("%d = %d", 0x1234567890ABCDEF, PopCount_(0x1234567890ABCDEF))
}

func TestPopCount__(t *testing.T) {
	colorlog.Info("%d = %d", 0, PopCount__(0))
	colorlog.Info("%d = %d", 1, PopCount__(1))
	colorlog.Info("%d = %d", 0x1234567890ABCDEF, PopCount__(0x1234567890ABCDEF))
}

func TestPopCount___(t *testing.T) {
	colorlog.Info("%d = %d", 0, PopCount___(0))
	colorlog.Info("%d = %d", 1, PopCount___(1))
	colorlog.Info("%d = %d", 0x1234567890ABCDEF, PopCount___(0x1234567890ABCDEF))
}

func BenchmarkPopCount(b *testing.B) { // 0.553s
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount_(b *testing.B) { // 1.882s
	for i := 0; i < b.N; i++ {
		PopCount_(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount__(b *testing.B) { // 1.525s
	for i := 0; i < b.N; i++ {
		PopCount__(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount___(b *testing.B) { // 1.538s
	for i := 0; i < b.N; i++ {
		PopCount___(0x1234567890ABCDEF)
	}
}
