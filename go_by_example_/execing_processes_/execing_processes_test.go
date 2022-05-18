package execing_processes_

import (
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestExecingProcesses(t *testing.T) {
	binary, lookErr := exec.LookPath("/bin/ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
