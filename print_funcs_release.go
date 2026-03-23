//go:build release && !discard

package log

// init initializes log output for release mode.
func init() {
	SetOutput(GetOutputWriterHourly(ReleaseLogPath))
}
