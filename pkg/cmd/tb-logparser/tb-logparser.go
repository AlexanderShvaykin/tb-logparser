package tb_logparser

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	wg      sync.WaitGroup
	LogName string
	hosts   []string
	host    string
	errors  []string
	Local   bool
)

func Execute() {
	rootCmd.AddCommand(cmdGrep)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(errors) > 0 {
		fmt.Printf("Erros: %v", errors)
	}
}

func init() {
	cmdGrep.PersistentFlags().StringVarP(&host,
		"host", "s", "",
		"host name, you can give many host, example: cat my_hosts | tb-logparser [command]")

	cmdGrep.PersistentFlags().StringVarP(&LogName,
		"filename", "f", "production",
		`Which files need scan, example: tb-logparser grep 'mail: find@foo.ru' -f shared/log/sidekiq.log`)

	cmdGrep.PersistentFlags().BoolVarP(&Local, "local", "", false,
		"Will be scan ./**/* files")

	if len(host) > 0 {
		hosts = append(hosts, host)
	}
}

func readHosts() {
	if len(hosts) == 0 && !Local {
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			line := in.Text()
			hosts = append(hosts, line)
		}
		if err := in.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}
