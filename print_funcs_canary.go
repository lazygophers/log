//go:build (test || canary) && !discard

package log

// init enables debug-level logging to stdout for test/canary builds.
func init() {
	SetLevel(DebugLevel)
	SetOutput(GetOutputWriterHourly(ReleaseLogDir))
}
