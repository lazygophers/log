//go:build !discard && (debug || canary || (!debug && !release && !canary))

package log

import "os"

func init() {
	SetOutput(os.Stdout)
}
