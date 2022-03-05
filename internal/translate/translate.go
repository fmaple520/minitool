package translate

import "minitool/internal/translate/translator"

type Translator interface {
	Translate() (string, error)
}

// todo 支持多种 translator
func NewTranslator(q string, from string, to string, appId string, secret string) Translator {
	return translator.NewBaiduTranslator(translator.SetQuery(q), translator.SetFrom(from), translator.SetTo(to), translator.SetAppID(appId), translator.SetSecret(secret))
}

func Translate(words, from, to, appId, secret string) string {
	translator := NewTranslator(words, from, to, appId, secret)
	res, err := translator.Translate() // 翻译结果处理
	if err != nil {
		return "error"
	}
	return res
}
