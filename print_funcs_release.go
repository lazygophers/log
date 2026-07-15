//go:build (release || prod || production) && !discard

package log

// init initializes log output for release mode.
func init() {
	SetOutput(GetOutputWriterHourly(ReleaseLogDir))
}
