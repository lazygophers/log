//go:build !release

package log

// ReleaseLogPath is empty in debug/default mode, logs go to stdout
var ReleaseLogPath string = ""
