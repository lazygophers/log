package cli

import "github.com/spf13/cobra"

func getBool(key string, cmds ...*cobra.Command) (ok bool) {
	c := rootCmd
	if len(cmds) > 0 {
		c = cmds[0]
	}

	ok, _ = c.Flags().GetBool(key)
	return ok
}

func getString(key string, cmds ...*cobra.Command) string {
	c := rootCmd
	if len(cmds) > 0 {
		c = cmds[0]
	}

	value, _ := c.Flags().GetString(key)
	return value
}
