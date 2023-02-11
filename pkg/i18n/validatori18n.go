package i18n

import (
	"errors"
	"reflect"

	"github.com/clz.skywalker/event.shop/kernal/pkg/consts"
	"github.com/clz.skywalker/event.shop/kernal/pkg/datatypesx"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

type validation struct {
	lang       string
	validation *validator.Validate
	trans      ut.Translator
}

type ParaValidation struct {
	validate map[string]*validation
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-10
 * @Description    : 初始化请求参数校验解析器
 * @return          {*}
 */
func NewParaValidation() *ParaValidation {
	uni := ut.New(zh.New(), en.New())
	zhTrans, _ := uni.GetTranslator(consts.LangChinese)
	enTrans, _ := uni.GetTranslator(consts.LangEnglish)
	enValidator := validator.New()
	zhValidator := validator.New()
	enValidator.RegisterCustomTypeFunc(validateJSONDateTypeLocalTime, datatypesx.LocalTime{})
	zhValidator.RegisterCustomTypeFunc(validateJSONDateTypeLocalTime, datatypesx.LocalTime{})
	if err := enTranslation.RegisterDefaultTranslations(enValidator, enTrans); err != nil {
		panic(err)
	}
	if err := zhTranslation.RegisterDefaultTranslations(zhValidator, zhTrans); err != nil {
		panic(err)
	}
	enValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	zhValidator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	// 验证器注册翻译器
	var enValidation = &validation{
		lang:       consts.LangEnglish,
		validation: enValidator,
		trans:      enTrans,
	}

	var zhValidation = &validation{
		lang:       consts.LangChinese,
		validation: zhValidator,
		trans:      zhTrans,
	}
	return &ParaValidation{validate: map[string]*validation{consts.LangChinese: zhValidation, consts.LangEnglish: enValidation}}
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-11
 * @Description    : 自定义类型，gin解析
 * @param           {reflect.Value} field
 * @return          {*}
 */
func validateJSONDateTypeLocalTime(field reflect.Value) interface{} {
	if field.Type() != reflect.TypeOf(datatypesx.LocalTime{}) {
		return nil
	}
	timeStr := field.Interface().(datatypesx.LocalTime).String()
	if timeStr == "0001-01-01 00:00:00" {
		return nil
	}
	return timeStr
}

/**
 * @Author         : Angular
 * @Date           : 2023-02-10
 * @Description    : 检验参数
 * @param           {string} locale
 * @param           {interface{}} models
 * @return          {*}
 */
func (p ParaValidation) ValidateParam(locale string, models interface{}) error {
	var validate *validation
	var ok bool
	if validate, ok = p.validate[locale]; !ok {
		validate = p.validate[consts.LangEnglish]
	}
	err := validate.validation.Struct(models)
	if err != nil {
		switch err1 := err.(type) {
		case validator.ValidationErrors:
			for _, err2 := range err1 {
				return errors.New(err2.Translate(validate.trans))
			}
		case *validator.InvalidValidationError:
			return errors.New(err1.Error())
		default:
			return errors.New(err1.Error())
		}

	}
	return nil
}
