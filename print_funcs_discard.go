//go:build discard

package log

import "io"

// init sets the standard logger output to io.Discard during package initialization.
// io.Discard consumes and discards all written data, enabling zero-cost logging.
func init() {
	SetOutput(io.Discard)
}
