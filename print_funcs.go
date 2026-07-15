//go:build !discard && !test && !canary && (debug || (!debug && !release && !canary && !prod && !production))

package log

import "os"

func init() {
	SetOutput(os.Stdout)
}
