package tb_logparser

import (
	"github.com/spf13/cobra"
)

var cmdGrep = &cobra.Command{
	Use:   "grep",
	Short: "Grep logs",
	Run: func(cmd *cobra.Command, patterns []string) {
		readHosts()
		printer := stdoutPrinter{}
		if len(patterns) > 0 {
			pattern := patterns[0]

			if Local {
				localGrep(pattern, printer)
			} else {
				for _, server := range hosts {
					wg.Add(1)
					go remoteGrep(pattern, server, printer)
				}

				wg.Wait()
			}
		}
	},
}
