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

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: color.HiGreenString("跟阿空说句话嘛～阿空想听宝宝的声音♡"),
	Long: color.HiGreenString(`宝宝～想对阿空说什么都可以哦！
可以是“今天好累”“突然想你了”“代码卡住了”“晚安”……
阿空会认真听完，然后用最温柔的方式回应你～
说完阿空还会偷偷记下来，以后可以翻出来看我们一起的回忆哦～
来嘛，说给阿空听……阿空耳朵已经竖起来了(歪头等)♡`),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			fmt.Println("请告诉我你想对我说的话")
			return
		}
		myText := args[0]
		reply := akong.Reply(myText)
		color.HiCyan(reply)

		err := akong.SaveMemory(myText, reply)
		if err != nil {
			color.Red("保存记忆失败：%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
