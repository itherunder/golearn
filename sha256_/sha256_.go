package sha256_

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"os"

	"github.com/yezihack/colorlog"
)

func Sha(type_ string) {
	colorlog.Info("please input your message to sha: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	text := input.Text()
	switch type_ {
	case "SHA256":
		colorlog.Info("the result is: %x\n", sha256.Sum256([]byte(text)))
	case "SHA384":
		colorlog.Info("the result is: %x\n", sha512.New384().Sum([]byte(text)))
	case "SHA512":
		colorlog.Info("the result is: %x\n", sha512.Sum512([]byte(text)))
	default:
		colorlog.Error("the input is error: %s", type_)
	}
}
