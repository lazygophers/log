package log

import (
	"fmt"
	"strconv"
)

// fastSprint optimizes fmt.Sprint for common cases
//
//go:inline
func fastSprint(args ...interface{}) string {
	switch len(args) {
	case 0:
		return ""
	case 1:
		return fastStringify(args[0])
	default:
		return fmt.Sprint(args...) // Fallback for multiple args
	}
}

// fastSprintf optimizes fmt.Sprintf for common cases
//
//go:inline
func fastSprintf(format string, args ...interface{}) string {
	if len(args) == 0 {
		return format // No formatting needed
	}
	if len(args) == 1 && format == "%s" {
		return fastStringify(args[0]) // Fast path for %s
	}
	return fmt.Sprintf(format, args...) // Fallback to standard
}

// fastStringify converts a single value to string without reflection for common types
//
//go:inline
func fastStringify(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val) // Simple conversion, still faster than fmt.Sprint
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	case error:
		if val != nil {
			return val.Error()
		}
		return "<nil>"
	case nil:
		return ""
	default:
		// Fallback to standard fmt for other types
		return fmt.Sprint(val)
	}
}
