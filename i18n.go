//提供json 语言包功能
// 为了方便excel管理，语言包不支持多级

package i18n

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"path"
	"reflect"
	"strings"
)

var locales = &store{
	defaultLanguage: "en",
	locales:         make(map[string]*locale),
}

type locale struct {
	language string
	message  gjson.Result
}

func IsExist(lang string) bool {
	_, ok := locales.locales[lang]
	return ok
}

type store struct {
	defaultLanguage string
	locales         map[string]*locale
}

//Add 添加言语包类型
func (s *store) Add(locale *locale) {
	if _, ok := s.locales[locale.language]; ok {
		return
	}
	s.locales[locale.language] = locale
}

// 设置默认语言
func DefaultLanguage(lang string) { locales.DefaultLanguage(lang) }
func (s *store) DefaultLanguage(lang string) {
	s.defaultLanguage = lang
}

func LoadPath(filepath string) error {
	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f := path.Join(filepath, file.Name())
		err = Load(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func Load(filePath string, langType ...string) error { return locales.Load(filePath, langType...) }
func (s *store) Load(filePath string, langType ...string) error {
	var lang string
	if len(langType) == 1 {
		lang = langType[0]
	} else {
		lang = getLangFromPath(filePath)
		if lang == "" {
			return fmt.Errorf("加载文件格式不对")
		}
	}
	message, err := getMessageContent(filePath)
	if err != nil {
		return err
	}
	locales.locales[lang] = &locale{
		language: lang,
		message:  gjson.Parse(message),
	}

	return nil
}

func (s *store) Get(lang, key string) (string, bool) {
	locale, ok := s.locales[lang]
	if !ok {
		return "", ok
	}
	res := locale.message.Get(key)
	if res.Exists() {
		return res.String(), ok
	}
	return "", false
}

func getLangFromPath(filePath string) string {
	_, filename := path.Split(filePath)
	files := strings.Split(filename, ".")
	if len(files) == 2 {
		return files[0]
	}
	return ""
}

func getMessageContent(filePath string) (string, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

type Locate struct {
	Language string
}

func (l *Locate) Lang(key string, defaultValue string, args ...interface{}) string {
	return Lang(l.Language, key, defaultValue, args...)
}
func (l *Locate) I(key string, args ...interface{}) string {
	var lang = l.Language
	if l.Language == "" {
		lang = locales.defaultLanguage
	}
	return Lang(lang, key, "", args...)
}

func Lang(lang string, key string, defaultValue string, args ...interface{}) string {
	var value string
	localeVal, ok := locales.Get(lang, key)
	if !ok {
		value = key
		if defaultValue != "" {
			value = defaultValue
		}
		return value
	}

	value = localeVal

	if len(args) > 0 {
		var a []interface{}
		for _, arg := range args {
			if arg == nil {
				continue
			}

			rv := reflect.ValueOf(arg)

			switch rv.Kind() {
			case reflect.String,
				reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64,
				reflect.Float32,
				reflect.Float64,
				reflect.Uint,
				reflect.Uint8,
				reflect.Uint16,
				reflect.Uint32,
				reflect.Uint64,
				reflect.Uintptr:
				a = append(a, arg)
			case reflect.Slice, reflect.Array:
				for i := 0; i < rv.Len(); i++ {
					a = append(a, rv.Index(i).Interface())
				}
			}
		}
		return fmt.Sprintf(value, a...)
	}
	return value
}

func I(key string, args ...interface{}) string {
	return Lang(locales.defaultLanguage, key, "", args...)
}