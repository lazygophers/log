//go:build !discard && (debug || canary || (!debug && !release && !canary && !prod && !production))

package log

import "os"

func init() {
	SetOutput(os.Stdout)
}
