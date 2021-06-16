//Copyright [2020] kingzcheung <kingzcheung@gmail.com>.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//提供json 语言包功能
// 为了方便excel管理，语言包不支持多级

package i18n

import "golang.org/x/text/language"

type Localization struct {
	bundle *Bundle
	tags   []language.Tag
}

func NewLocalization(bundle *Bundle, lang string, fallbackLangs ...string) *Localization {
	l := &Localization{bundle: bundle}
	tags := make([]language.Tag, len(fallbackLangs)+1)
	tags = append(tags, language.Make(lang))
	if len(fallbackLangs) > 0 {
		for _, fallbackLang := range fallbackLangs {
			tags = append(tags, language.Make(fallbackLang))
		}
	}
	l.tags = tags
	return l
}

func (l *Localization) Localize(key string, variables ...map[string]interface{}) string {
	locales := l.bundle.Locales()
	if len(locales) == 0 {
		return key
	}

	var value string

	for _, locale := range locales {
		var ok bool
		for _, tag := range l.tags {
			if locale.tag.String() == tag.String() {
				value, ok = locale.message[key]
				if !ok {
					continue
				}
			}
		}
		if value == "" {
			if locale.tag.String() == l.bundle.defaultLanguage.String() {
				value, ok = locale.message[key]
				if !ok {
					value = key
				}
			}

		}
	}

	//TODO 解析变量

	return value
}
