/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"ddd/infrastructure/setting"
	"ddd/lib"
	"ddd/router"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clean-arc",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		settings := setting.Setting{
			DB: setting.DB{
				Type:     "mysql",
				Host:     "mysql",
				Port:     3306,
				User:     "root",
				Password: "test",
				Name:     "test",
			},
		}
		r := router.Get(settings)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()
			http.ListenAndServe(":3020", r)
		}()
		fmt.Println("Start")
		lib.WaitSignal()
		wg.Wait()
		fmt.Println("Graceful Shutdown...")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clean-arc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// ここでdbの設定を書く
}
