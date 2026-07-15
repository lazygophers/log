//go:build (test || canary) && !discard

package log

import "os"

// init enables debug-level logging to stdout for test/canary builds.
func init() {
	SetLevel(DebugLevel)
	SetOutput(os.Stdout)
}
