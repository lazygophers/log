package constant

// Hook defines the interface for log processing hooks
// Hooks can modify, filter, or enrich log entries before they are written
type Hook interface {
	// OnWrite processes the log entry before writing
	// Returns modified entry or nil to skip logging
	OnWrite(entry interface{}) interface{}
}

// HookFunc is a convenience type for implementing Hook with a function
type HookFunc func(entry interface{}) interface{}

// OnWrite implements Hook interface
func (h HookFunc) OnWrite(entry interface{}) interface{} {
	return h(entry)
}

// Format defines log formatting interface
type Format interface {
	// Format formats log entry to byte array
	Format(entry interface{}) []byte
}

// FormatFull extends Format with additional controls
type FormatFull interface {
	Format

	// ParsingAndEscaping controls message parsing and escaping
	ParsingAndEscaping(disable bool)

	// Caller enables or disables caller information in log output
	Caller(disable bool)

	// Clone creates a deep copy of the formatter
	Clone() Format
}

// Writer defines the output writer interface
type Writer interface {
	Write(p []byte) (n int, err error)
}
