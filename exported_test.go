package colorlog

import "testing"

func Test_ColorLog(t *testing.T) {
	Println("println")
	Trace("trace")
	Info("info")
	Warn("warning")
	Error("error")
}
