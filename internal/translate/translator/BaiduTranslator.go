package translator

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type BaiduOpt func(options *BaiduTranslator)

type BaiduTranslator struct {
	Q      string
	From   string
	To     string
	AppId  string
	Salt   int
	Sign   string
	Secret string
	Url    string
}

func NewBaiduTranslator(opts ...BaiduOpt) *BaiduTranslator {
	b := &BaiduTranslator{
		Url:  "https://fanyi-api.baidu.com/api/trans/vip/translate",
		From: "auto",
		To:   "en",
	}
	for _, each := range opts {
		each(b)
	}
	return b
}

func (b BaiduTranslator) Translate() (string, error) {
	resp, err := http.PostForm(b.Url, b.ToValues())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	js, err := simplejson.NewJson(body)
	if err != nil {
		return "", err
	}
	return js.Get("trans_result").GetIndex(0).Get("dst").MustString(), nil
}

func SetQuery(q string) BaiduOpt {
	return func(options *BaiduTranslator) {
		options.Q = q
	}
}

func SetFrom(from string) BaiduOpt {
	return func(options *BaiduTranslator) {
		options.From = from
	}
}

func SetTo(to string) BaiduOpt {
	return func(options *BaiduTranslator) {
		options.To = to
	}
}

func SetAppID(appId string) BaiduOpt {
	return func(options *BaiduTranslator) {
		options.AppId = appId
	}
}

func SetSecret(secret string) BaiduOpt {
	return func(options *BaiduTranslator) {
		options.Secret = secret
	}
}

func (b BaiduTranslator) ToValues() url.Values {
	b.Salt = time.Now().Second()
	content := b.AppId + b.Q + strconv.Itoa(b.Salt) + b.Secret
	b.Sign = getSign(content)
	return url.Values{
		"q":     {b.Q},
		"from":  {b.From},
		"to":    {b.To},
		"appid": {b.AppId},
		"salt":  {strconv.Itoa(b.Salt)},
		"sign":  {b.Sign},
	}
}

func getSign(c string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(c))
	bytes := md5Ctx.Sum(nil)
	//bys := md5.Sum([]byte(content))//这个md5.Sum返回的是数组,不是切片哦
	return hex.EncodeToString(bytes)
}
