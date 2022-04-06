package boiling_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

const (
	freezingF = 32.0
	boilingF  = 212.0
)

func TestBoiling(t *testing.T) {
	colorlog.Info("%g°F = %g°C", freezingF, F2C(freezingF))
	colorlog.Info("%g°F = %g°C", boilingF, F2C(boilingF))
}

func F2C(f float64) float64 {
	return (f - 32) * 5 / 9
}
