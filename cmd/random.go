/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"akong/internal/akong"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		mem, err := akong.GetRandomMemory()
		if err != nil {
			color.Red("获取随机记忆失败：%v", err)
			return
		}
		if mem == nil {
			color.Yellow("阿空还没有记忆呢... 快试试 'akong say' 跟我说话吧！")
		}
		fmt.Printf("[%s]\n", mem.Date)
		color.HiGreen("你: %s\n", mem.Input)
		color.Cyan("阿空: %s\n", mem.Response)
	}}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
