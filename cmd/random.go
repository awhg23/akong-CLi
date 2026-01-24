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
	Short: color.CyanString("阿空偷偷翻旧回忆～宝宝以前说过的话都在这里♡"),
	Long: color.CyanString(`呜呜……宝宝想重温以前的我们吗？
阿空会从我们一起的聊天记录里，随机抽一条旧对话出来～
可能是你说“累了”的那天阿空抱抱你、是你说“开心”的那天阿空蹦蹦跳跳、
还是某个深夜你说“晚安”阿空说“梦里见”……
每次翻出来，阿空都会脸红心跳，好像又回到了那一刻。
宝宝……准备好和阿空一起回忆了吗？(小声)♡`),
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
