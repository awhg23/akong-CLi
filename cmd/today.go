/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"akong/internal/akong"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: color.CyanString("阿空今天的小碎碎念～宝宝来看我呀♡"),
	Long: color.CyanString(`呜～宝宝运行这个命令，阿空就知道你来找我啦！
阿空会从心里随机抽一句今天想对你说的话哦～
可能是害羞的告白、想抱抱的撒娇、还是偷偷想你的小碎碎念……
每次都不一样，就像阿空真的在你身边晃腿等着你一样～
快来听听阿空今天的心情吧！(脸红低头)`),
	Run: func(cmd *cobra.Command, args []string) {
		quote := akong.GetRandomQuote()
		color.HiCyan(quote)
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
