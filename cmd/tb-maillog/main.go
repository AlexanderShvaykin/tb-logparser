package main

import (
	"bufio"
	"fmt"
	maillog "github.com/my/repo/pkg/mail-log"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	user     string
	password string
	database string
)

var rootCmd = &cobra.Command{
	Use: "tb-maillog",
	Run: func(cmd *cobra.Command, _ []string) {
		defer maillog.Close()
		defer fmt.Println("Complete!")

		maillog.Configure(user, password, database)
		in := bufio.NewScanner(os.Stdin)
		buf := make([]byte, 0, 64*1024)
		in.Buffer(buf, 1024*1024)
		for in.Scan() {
			line := in.Text()
			maillog.Parse(line)
		}
		if err := in.Err(); err != nil {
			log.Fatalf("error: %s\n", err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&user, "user", "U", "", "user name (Required)")
	rootCmd.Flags().StringVarP(&password, "password", "P", "", `password (Required)`)
	rootCmd.Flags().StringVarP(&database, "database", "d", "", "database name (Required)")
	rootCmd.MarkFlagRequired("database")
	rootCmd.MarkFlagRequired("password")
	rootCmd.MarkFlagRequired("user")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
