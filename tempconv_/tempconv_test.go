package tempconv_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestKToC(t *testing.T) {
	colorlog.Info("%g°K = %g°C", Kelvin(0), KToC(Kelvin(0)))
	colorlog.Info("%g°K = %g°C", Kelvin(273.15), KToC(Kelvin(273.15)))
	colorlog.Info("%g°K = %g°C", Kelvin(373.15), KToC(Kelvin(373.15)))
}
