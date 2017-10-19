package cmd

import "github.com/spf13/cobra"

var Webhulk = &cobra.Command{
	Use:   "webulk [command]",
	Short: "Webulk - A lightweight API for managing webhooks",
}

func defaultTo(defaultValue, value string) string {
	if value == "" {
		return defaultValue
	}

	return value
}
