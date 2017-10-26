package cmd

import "github.com/spf13/cobra"

var Webhulk = &cobra.Command{
	Use:   "webhulk [command]",
	Short: "Webhulk - A lightweight API for managing webhooks",
}

func defaultTo(defaultValue, value string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func getFlag(cmd *cobra.Command, flag string) (string, bool) {
	flags := cmd.Flags()

	return flags.Lookup(flag).Value.String(), true
}

func getConfigFile(cmd *cobra.Command) string {
	flag, _ := getFlag(cmd, "config")
	config := defaultTo("webhulk.yml", flag)

	return config
}
