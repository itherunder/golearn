package signals_

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignals(t *testing.T) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
