package rev_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestRev(t *testing.T) {
	s := []byte("itherunder is liaozhou")
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	colorlog.Info("\nbefore: %s\nafter: %s\n", []byte("itherunder is liaozhou"), s)

	a := [...]int{0, 1, 2, 3, 4, 5}
	colorlog.Info("a is %v", a)
	reverse(a[:2]) // left move
	reverse(a[2:])
	reverse(a[:])
	colorlog.Info("a is %v", a)

	colorlog.Info("reverse a utf8 string with its []byte")
	// reverseUtf8WithByte([]byte("hello, 世界"))
	reverseUtf8WithByte([]byte("你好, world"))
}

func reverse(s []int) {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func cntRuneByByte(s []byte, i int) (int, int) {
	if s[i]&0b10000000 == 0 { // 1 byte rune 0...
		return i, 1
	} else {
		if s[i]&0b01000000 == 0 { // 10...
			return cntRuneByByte(s, i-1)
		}
		if s[i]&0b00100000 == 0 { // 2 byte rune 110... 10...
			return i, 2
		} else if s[i]&0b00010000 == 0 { // 3 byte rune 1110... 10... 10...
			return i, 3
		} else { // 4 byte rune 11110... 10... 10... 10...
			return i, 4
		}
	}
}

func findRuneWithByteBack(s []byte, j int) (int, int) {
	for i := 3; i > 0; i-- {
		if start, cnt := cntRuneByByte(s, j-i); cnt == i+1 {
			return start, cnt
		}
	}
	return j, 1
}

func reverseUtf8WithByte(s []byte) {
	colorlog.Info("the reverse before is %s\n %08b\n", s, s)
	n := len(s)
	for i, j := 0, n-1; i < j; {
		_, cnt := cntRuneByByte(s, i)
		start_, cnt_ := findRuneWithByteBack(s, j)
		colorlog.Debug("start_: %d, cnt_: %d, cnt: %d", start_, cnt_, cnt)
		if cnt > cnt_ { // left move cnt-cnt_
			for k := 0; k < cnt_; k++ {
				s[i+cnt-1-k], s[j-k] = s[j-k], s[i+cnt-1-k]
			}
			for k := cnt_; k < cnt; k++ {
				tmp := s[i+cnt-1-k]
				for l := i + cnt - 1 - k; l < start_-1; l++ {
					s[l] = s[l+1]
				}
				s[start_-1-(k-cnt_)] = tmp
			}
		} else if cnt_ > cnt { // right move cnt_-cnt
			for k := 0; k < cnt; k++ {
				s[i+cnt-1-k], s[j-k] = s[j-k], s[i+cnt-1-k]
			}
			for k := cnt; k < cnt_; k++ {
				tmp := s[j-cnt]
				for l := j - cnt; l > i; l-- {
					s[l] = s[l-1]
				}
				s[i] = tmp
			}
		} else {
			for k := 0; k < cnt; k++ {
				s[i+cnt-1-k], s[j-k] = s[j-k], s[i+cnt-1-k]
			}
		}
		// colorlog.Info("process: %s\n %08b\n", s, s)
		i += cnt_
		j -= cnt
	}
	colorlog.Info("the reverse result is %s\n %08b\n", s, s)
}
