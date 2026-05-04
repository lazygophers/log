package main

import (
	"fmt"

	"github.com/lazygophers/log"
	"github.com/lazygophers/log/constant"
	"github.com/lazygophers/log/hooks"
)

func main() {
	fmt.Println("=== Go Log Hooks System Demo ===")

	// 1. Basic Hook
	fmt.Println("1. Basic Hook:")
	logger1 := log.New()
	logger1.EnableCaller(false)
	logger1.EnableTrace(false)

	basicHook := constant.HookFunc(func(entry interface{}) interface{} {
		fmt.Println("  → Hook executed!")
		return entry
	})

	logger1.AddHook(basicHook).Info("Message with hook")

	// 2. Message Filter
	fmt.Println("\n2. Message Filter (only 'important'):")
	logger2 := log.New()
	logger2.EnableCaller(false)
	logger2.EnableTrace(false)

	filterHook := hooks.NewMessageFilterHook()
	filterHook.AddAllowPattern("important")
	logger2.AddHook(filterHook)

	logger2.Info("This is important")
	logger2.Info("This is not important")

	// 3. Sensitive Data Masking
	fmt.Println("\n3. Sensitive Data Masking:")
	logger3 := log.New()
	logger3.EnableCaller(false)
	logger3.EnableTrace(false)

	maskHook := hooks.NewSensitiveDataMaskHook()
	logger3.AddHook(maskHook)

	logger3.Info("Email: test@example.com")
	logger3.Info("Credit card: 4111-1111-1111-1111")
	logger3.Info("Password: secret123")

	// 4. Level Filter
	fmt.Println("\n4. Level Filter (WARN and above only):")
	logger4 := log.New()
	logger4.EnableCaller(false)
	logger4.EnableTrace(false)

	levelHook := hooks.NewLevelFilterHook(int(log.WarnLevel))
	logger4.AddHook(levelHook)

	logger4.Info("This info won't show")
	logger4.Warn("This warning will show")
	logger4.Error("This error will show")

	// 5. Context Enrichment
	fmt.Println("\n5. Context Enrichment:")
	logger5 := log.New()
	logger5.EnableCaller(false)
	logger5.EnableTrace(false)

	fields := map[string]interface{}{
		"service": "demo-service",
		"version": "1.0.0",
	}
	enrichHook := hooks.NewContextEnrichHook(fields)
	logger5.AddHook(enrichHook)

	logger5.Info("Request processed")

	// 6. Prefix and Suffix
	fmt.Println("\n6. Prefix and Suffix:")
	logger6 := log.New()
	logger6.EnableCaller(false)
	logger6.EnableTrace(false)

	prefixHook := hooks.NewPrefixHook("[APP] ")
	suffixHook := hooks.NewSuffixHook(" [END]")
	logger6.AddHooks(prefixHook, suffixHook)

	logger6.Info("Message with prefix and suffix")

	// 7. Max Length
	fmt.Println("\n7. Max Length (truncated at 20 chars):")
	logger7 := log.New()
	logger7.EnableCaller(false)
	logger7.EnableTrace(false)

	maxHook := hooks.NewMaxLengthHook(20)
	logger7.AddHook(maxHook)

	logger7.Info("This is a very long message that should be truncated")

	// 8. Conditional Hook
	fmt.Println("\n8. Conditional Hook (prefix for ERROR only):")
	logger8 := log.New()
	logger8.EnableCaller(false)
	logger8.EnableTrace(false)

	condition := func(entry interface{}) bool {
		if e, ok := entry.(*log.Entry); ok {
			return e.Level == log.ErrorLevel
		}
		return false
	}

	conditionalPrefix := hooks.NewPrefixHook("[ERROR] ")
	conditionalHook := hooks.NewConditionalHook(condition, conditionalPrefix)
	logger8.AddHook(conditionalHook)

	logger8.Info("Normal message")
	logger8.Error("Error message")

	// 9. Chain Hook
	fmt.Println("\n9. Chain Hook (prefix + truncate):")
	logger9 := log.New()
	logger9.EnableCaller(false)
	logger9.EnableTrace(false)

	chain := hooks.NewChainHook(
		hooks.NewPrefixHook("[CHAINED] "),
		hooks.NewMaxLengthHook(25),
	)
	logger9.AddHook(chain)

	logger9.Info("This message goes through multiple hooks")

	// 10. Custom Mask Pattern
	fmt.Println("\n10. Custom Phone Number Mask:")
	logger10 := log.New()
	logger10.EnableCaller(false)
	logger10.EnableTrace(false)

	customMask := hooks.NewSensitiveDataMaskHook()
	customMask.AddPattern(`\b\d{3}-\d{3}-\d{4}\b`)
	customMask.SetMask("[PHONE]")
	logger10.AddHook(customMask)

	logger10.Info("Call us at 555-123-4567")

	// 11. Field Filter
	fmt.Println("\n11. Field Filter:")
	logger11 := log.New()
	logger11.EnableCaller(false)
	logger11.EnableTrace(false)

	fieldFilter := hooks.NewFieldFilterHook()
	fieldFilter.AllowField("environment", "production")
	logger11.AddHook(fieldFilter)

	logger11.Infow("Message", "environment", "production")
	logger11.Infow("Message", "environment", "development")

	fmt.Println("\n=== Demo Complete ===")
}
