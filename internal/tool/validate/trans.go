package validate

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	// 获取gin框架中的Validator引擎，进行修改定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 1.初始化翻译器
		zhT := zh.New()
		enT := en.New()

		// 2.生成一个国际化的翻译者
		// 第一个参数，默认翻译器；后续参数，支持的翻译器
		uni := ut.New(enT, zhT, enT)

		// 3.依据当前请求的语言环境locale，切换至对应的翻译者
		// // locale 通常取决于 http 请求头的 'Accept-Language'
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		var ok bool
		if trans, ok = uni.GetTranslator(locale); !ok {
			fmt.Printf("uni.GetTranslator(%s) failed", locale)
			return
		}

		// 4.为所有的内置验证方法注册一个可用的翻译者
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			fmt.Printf("register default trans failed: %s\n", err)
		}
	}
	return
}

// 移除验证后的错误信息中，字段所处的结构体名称
// 批量校验的，最后是一个，map存储字段及其错误信息
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, errMsg := range fields {
		firstPointIndex := strings.Index(field, ".") + 1 // 提取field中第一个“.”的位置，并移动一位以保留其后部分
		res[field[firstPointIndex:]] = errMsg            // 切割字符串，左闭右边开
	}
	return res
}

// 注册自定义验证方法
func RegisterTranslation(tag string, msg string, fn validator.Func, transFun validator.TranslationFunc) error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation(tag, fn); err != nil {
			return err
		}
		if err := v.RegisterTranslation(
			tag,
			trans,
			CreateTranslationsFun(tag, msg),
			transFun,
		); err != nil {
			return err
		}
	}
	return nil
}

// 注册自定义验证方法的错误信息提示
// tag 验证方法名；msg 错误信息，注意占位{0}
func CreateTranslationsFun(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义验证方法，在注册时需要定义一个翻译方法
func Translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
