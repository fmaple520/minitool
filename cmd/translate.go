package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"minitool/internal/translate"
	"minitool/internal/word"
	"strings"
)

var (
	words      string
	from       string
	to         string
	translator string
	mod        int8
)

type TranslateS struct {
	DefaultTranslator string
	Translator        []string
	AppKey            []string
	AppSecret         []string
}

var TranslateSetting *TranslateS

var transCmd = &cobra.Command{
	Use:   "trans",
	Short: "翻译[转换]",
	Long:  transDesc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		content = translate.Translate(words, from, to, TranslateSetting.AppKey[0], TranslateSetting.AppSecret[0])
		if mod > 0 {
			switch mod {
			case ModeUpper:
				content = word.ToUpper(content)
			case ModeLower:
				content = word.ToLower(content)
			case ModeUnderlineToUpperCamelCase:
				content = word.UnderlineToUpperCamelCase(content)
			case ModeUnderlineToLowerCamelCase:
				content = word.UnderlineToLowerCamelCase(content)
			default:
				log.Fatalf("暂时不支持该转换模式，请执行 help word 查看帮助文档")
			}
		}

		log.Printf("输出结果：%s", content)
	},
}

func init() {
	transCmd.Flags().StringVarP(&translator, "translator", "T", "baidu", "翻译员")
	transCmd.Flags().StringVarP(&words, "words", "w", "", "请输入单词内容")
	transCmd.Flags().StringVarP(&to, "from", "f", "auto", "源语言")
	transCmd.Flags().StringVarP(&from, "to", "t", "zh", "目标语言")
	transCmd.Flags().Int8VarP(&mod, "mod", "m", 0, "请输入单词转换的模式")
}

var transDesc = strings.Join([]string{
	"该子命令支持对各种语言进行翻译，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：转大写驼峰",
	"4：转小写驼峰",
}, "\n")
