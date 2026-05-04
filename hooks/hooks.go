package hooks

import (
	"regexp"
	"unicode/utf8"

	"github.com/lazygophers/log/constant"
)

// Hook is exported from constant package for convenience
type Hook = constant.Hook

// HookFunc is exported from constant package for convenience
type HookFunc = constant.HookFunc

// SensitiveDataMaskHook masks sensitive data in log messages
type SensitiveDataMaskHook struct {
	// Patterns to mask (regex patterns)
	patterns []*regexp.Regexp
	// Mask character
	mask string
	// Mask fields by key
	maskFields map[string]bool
}

// NewSensitiveDataMaskHook creates a new sensitive data masking hook
func NewSensitiveDataMaskHook() *SensitiveDataMaskHook {
	return &SensitiveDataMaskHook{
		patterns: []*regexp.Regexp{
			regexp.MustCompile(`\b\d{4}[- ]?\d{4}[- ]?\d{4}[- ]?\d{4}\b`), // Credit card
			regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`), // Email
			regexp.MustCompile(`\b\d{3}-\d{2}-\d{4}\b`), // SSN
			regexp.MustCompile(`\b(?:password|passwd|pwd|token|secret|key)\s*[:=]\s*\S+`), // Password fields
			regexp.MustCompile(`"(?:password|passwd|pwd|token|secret|key)"\s*:\s*"[^"]+"`), // JSON password
		},
		mask:       "***",
		maskFields: map[string]bool{"password": true, "token": true, "secret": true, "key": true},
	}
}

// AddPattern adds a custom regex pattern to mask
func (h *SensitiveDataMaskHook) AddPattern(pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	h.patterns = append(h.patterns, re)
	return nil
}

// SetMask sets the mask character/string
func (h *SensitiveDataMaskHook) SetMask(mask string) {
	h.mask = mask
}

// AddMaskField adds a field key to mask
func (h *SensitiveDataMaskHook) AddMaskField(key string) {
	h.maskFields[key] = true
}

// OnWrite implements Hook interface
func (h *SensitiveDataMaskHook) OnWrite(entry interface{}) interface{} {
	// Type assertion to access Entry fields
	// This works because we pass the actual *Entry from logger
	type entryLike struct {
		Message    string
		Fields     []interface{}
		File       string
		CallerName string
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	// Mask in message
	e.Message = h.maskString(e.Message)

	// Mask in fields
	for i, field := range e.Fields {
		if fieldMap, ok := field.(map[string]interface{}); ok {
			// Handle field as map
			for key, val := range fieldMap {
				if h.maskFields[key] {
					fieldMap[key] = h.mask
				} else if strVal, ok := val.(string); ok {
					fieldMap[key] = h.maskString(strVal)
				}
			}
			e.Fields[i] = fieldMap
		}
	}

	// Mask in caller info
	e.File = h.maskString(e.File)
	e.CallerName = h.maskString(e.CallerName)

	return entry
}

func (h *SensitiveDataMaskHook) maskString(s string) string {
	for _, pattern := range h.patterns {
		s = pattern.ReplaceAllString(s, h.mask)
	}
	return s
}

// ContextEnrichHook automatically adds contextual fields to log entries
type ContextEnrichHook struct {
	fields map[string]interface{}
}

// NewContextEnrichHook creates a new context enrichment hook
func NewContextEnrichHook(fields map[string]interface{}) *ContextEnrichHook {
	return &ContextEnrichHook{fields: fields}
}

// AddField adds a field to enrich
func (h *ContextEnrichHook) AddField(key string, value interface{}) {
	if h.fields == nil {
		h.fields = make(map[string]interface{})
	}
	h.fields[key] = value
}

// SetFields sets multiple fields
func (h *ContextEnrichHook) SetFields(fields map[string]interface{}) {
	h.fields = fields
}

// OnWrite implements Hook interface
func (h *ContextEnrichHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Fields []interface{}
	}

	if entry == nil || len(h.fields) == 0 {
		return entry
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	// Append enrichment fields
	for key, value := range h.fields {
		e.Fields = append(e.Fields, map[string]interface{}{key: value})
	}

	return entry
}

// LevelFilterHook filters logs by level
type LevelFilterHook struct {
	minLevel int
}

// NewLevelFilterHook creates a new level filter hook
func NewLevelFilterHook(minLevel int) *LevelFilterHook {
	return &LevelFilterHook{minLevel: minLevel}
}

// OnWrite implements Hook interface
func (h *LevelFilterHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Level int
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	if e.Level < h.minLevel {
		return nil // Filter out
	}
	return entry
}

// MessageFilterHook filters logs by message content
type MessageFilterHook struct {
	// Allow patterns (if any matches, allow)
	allow []*regexp.Regexp
	// Deny patterns (if any matches, deny)
	deny []*regexp.Regexp
}

// NewMessageFilterHook creates a new message filter hook
func NewMessageFilterHook() *MessageFilterHook {
	return &MessageFilterHook{}
}

// AddAllowPattern adds a regex pattern that allows logs
func (h *MessageFilterHook) AddAllowPattern(pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	h.allow = append(h.allow, re)
	return nil
}

// AddDenyPattern adds a regex pattern that denies logs
func (h *MessageFilterHook) AddDenyPattern(pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	h.deny = append(h.deny, re)
	return nil
}

// OnWrite implements Hook interface
func (h *MessageFilterHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Message string
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	msg := e.Message

	// Check deny patterns first
	for _, pattern := range h.deny {
		if pattern.MatchString(msg) {
			return nil // Filter out
		}
	}

	// If no allow patterns, allow all
	if len(h.allow) == 0 {
		return entry
	}

	// Check allow patterns
	for _, pattern := range h.allow {
		if pattern.MatchString(msg) {
			return entry // Allow
		}
	}

	// No allow pattern matched, filter out
	return nil
}

// FieldFilterHook filters logs by field values
type FieldFilterHook struct {
	// Field key -> allowed values (empty means any value allowed)
	allowedFields map[string]map[interface{}]bool
	// Field key -> denied values
	deniedFields map[string]map[interface{}]bool
}

// NewFieldFilterHook creates a new field filter hook
func NewFieldFilterHook() *FieldFilterHook {
	return &FieldFilterHook{
		allowedFields: make(map[string]map[interface{}]bool),
		deniedFields: make(map[string]map[interface{}]bool),
	}
}

// AllowField adds an allowed value for a field
func (h *FieldFilterHook) AllowField(key string, value interface{}) {
	if h.allowedFields[key] == nil {
		h.allowedFields[key] = make(map[interface{}]bool)
	}
	h.allowedFields[key][value] = true
}

// DenyField adds a denied value for a field
func (h *FieldFilterHook) DenyField(key string, value interface{}) {
	if h.deniedFields[key] == nil {
		h.deniedFields[key] = make(map[interface{}]bool)
	}
	h.deniedFields[key][value] = true
}

// OnWrite implements Hook interface
func (h *FieldFilterHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Fields []interface{}
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	for _, field := range e.Fields {
		if fieldMap, ok := field.(map[string]interface{}); ok {
			for key, value := range fieldMap {
				// Check denied values first
				if deniedValues, ok := h.deniedFields[key]; ok {
					if deniedValues[value] {
						return nil // Filter out
					}
				}

				// Check allowed values
				if allowedValues, ok := h.allowedFields[key]; ok {
					if len(allowedValues) > 0 && !allowedValues[value] {
						return nil // Filter out
					}
				}
			}
		}
	}

	return entry
}

// MinLengthHook filters logs shorter than minimum length
type MinLengthHook struct {
	MinLength int
}

// NewMinLengthHook creates a new minimum length hook
func NewMinLengthHook(minLength int) *MinLengthHook {
	return &MinLengthHook{MinLength: minLength}
}

// OnWrite implements Hook interface
func (h *MinLengthHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Message string
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	// Count UTF-8 characters (not bytes)
	length := utf8.RuneCountInString(e.Message)
	if length < h.MinLength {
		return nil // Filter out short messages
	}

	return entry
}

// MaxLengthHook truncates log messages longer than maximum length
type MaxLengthHook struct {
	MaxLength      int
	TruncateSuffix string
}

// NewMaxLengthHook creates a new maximum length hook
func NewMaxLengthHook(maxLength int) *MaxLengthHook {
	return &MaxLengthHook{
		MaxLength:      maxLength,
		TruncateSuffix: "...",
	}
}

// OnWrite implements Hook interface
func (h *MaxLengthHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Message string
	}

	if entry == nil {
		return nil
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	// Count UTF-8 characters
	length := utf8.RuneCountInString(e.Message)
	if length > h.MaxLength {
		// Truncate message
		runes := []rune(e.Message)
		e.Message = string(runes[:h.MaxLength]) + h.TruncateSuffix
	}

	return entry
}

// PrefixHook adds a prefix to log messages
type PrefixHook struct {
	Prefix string
}

// NewPrefixHook creates a new prefix hook
func NewPrefixHook(prefix string) *PrefixHook {
	return &PrefixHook{Prefix: prefix}
}

// OnWrite implements Hook interface
func (h *PrefixHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Message string
	}

	if entry == nil || h.Prefix == "" {
		return entry
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	e.Message = h.Prefix + e.Message
	return entry
}

// SuffixHook adds a suffix to log messages
type SuffixHook struct {
	Suffix string
}

// NewSuffixHook creates a new suffix hook
func NewSuffixHook(suffix string) *SuffixHook {
	return &SuffixHook{Suffix: suffix}
}

// OnWrite implements Hook interface
func (h *SuffixHook) OnWrite(entry interface{}) interface{} {
	type entryLike struct {
		Message string
	}

	if entry == nil || h.Suffix == "" {
		return entry
	}

	e, ok := entry.(entryLike)
	if !ok {
		return entry
	}

	e.Message = e.Message + h.Suffix
	return entry
}

// ConditionalHook executes a hook only when condition is met
type ConditionalHook struct {
	condition func(interface{}) bool
	hook      constant.Hook
}

// NewConditionalHook creates a new conditional hook
func NewConditionalHook(condition func(interface{}) bool, hook constant.Hook) *ConditionalHook {
	return &ConditionalHook{
		condition: condition,
		hook:      hook,
	}
}

// OnWrite implements Hook interface
func (h *ConditionalHook) OnWrite(entry interface{}) interface{} {
	if entry == nil || !h.condition(entry) {
		return entry
	}
	return h.hook.OnWrite(entry)
}

// ChainHook executes multiple hooks in sequence
type ChainHook struct {
	hooks []constant.Hook
}

// NewChainHook creates a new chain hook
func NewChainHook(hooks ...constant.Hook) *ChainHook {
	return &ChainHook{hooks: hooks}
}

// OnWrite implements Hook interface
func (h *ChainHook) OnWrite(entry interface{}) interface{} {
	if entry == nil {
		return nil
	}

	for _, hook := range h.hooks {
		entry = hook.OnWrite(entry)
		if entry == nil {
			return nil // Stop chain if hook filtered
		}
	}

	return entry
}
