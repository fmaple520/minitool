package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"minitool/internal/word"
	"strings"
)

const (
	ModeUpper                     = iota + 1 // 全部转大写
	ModeLower                                // 全部转小写
	ModeUnderlineToUpperCamelCase            // 下划线转大驼峰
	ModeUnderlineToLowerCamelCase            // 下划线转小驼峰
	ModeCamelCaseToUnderline                 // 驼峰转下划线
)

var (
	str  string
	mode int8
)

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  wordDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderlineToUpperCamelCase:
			content = word.UnderlineToUpperCamelCase(str)
		case ModeUnderlineToLowerCamelCase:
			content = word.UnderlineToLowerCamelCase(str)
		case ModeCamelCaseToUnderline:
			content = word.CamelCaseToUnderline(str)
		default:
			log.Fatalf("暂时不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}

var wordDesc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")
