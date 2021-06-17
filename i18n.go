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

import (
	"bytes"
	"golang.org/x/text/language"
	"html/template"
	"net/http"
	"strings"
)

type Localization struct {
	bundle *Bundle
	tags   []language.Tag
}

func NewLocalization(bundle *Bundle, lang string, fallbackLangs ...string) *Localization {
	l := &Localization{bundle: bundle}
	tags := make([]language.Tag, len(fallbackLangs)+1)
	tags[0] = language.Make(lang)
	if len(fallbackLangs) > 0 {
		for i, fallbackLang := range fallbackLangs {
			tags[i+1] = language.Make(fallbackLang)
		}
	}
	l.tags = tags
	return l
}

func (l *Localization) WithRequest(r *http.Request) *Localization {
	accept := r.Header.Get("Accept-Language")
	return l.With(accept)
}

func (l *Localization) With(lang string) *Localization {
	tag := language.Make(lang)
	return l.WithTag(tag)
}

func (l *Localization) WithTag(tag language.Tag) *Localization {
	tags := l.tags
	if tags == nil {
		l.tags = []language.Tag{tag}
	} else {
		l.tags = append([]language.Tag{tag}, tags...)
	}
	return l
}

func (l *Localization) Localize(key string, variables ...map[string]interface{}) string {
	locales := l.bundle.Locales()
	if len(locales) == 0 {
		return key
	}
	var value string
LOCALE:
	for _, tag := range l.tags {
		var ok bool
		for _, locale := range locales {
			if locale.tag.String() == tag.String() {
				value, ok = locale.message[key]
				if !ok {
					continue
				}
				break LOCALE
			}
		}
	}
	if value == "" {
		var ok bool
		for _, locale := range locales {
			if locale.tag.String() == l.bundle.defaultLanguage.String() {
				value, ok = locale.message[key]
				if !ok {
					continue
				}
			}
		}
	}

	if strings.Index(value, "{{") > -1 {
		value = l.parseTemplateValue(value, variables)
	}

	return value
}

func (l *Localization) parseTemplateValue(value string, variables []map[string]interface{}) string {
	vars := mergeVars(variables)
	bf := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(value)
	if err != nil {
		return value
	}

	err = tmpl.Execute(bf, vars)
	if err != nil {
		return value
	}
	return bf.String()
}

func mergeVars(variables []map[string]interface{}) map[string]interface{} {
	var vars = make(map[string]interface{})
	for _, variable := range variables {
		for k, v := range variable {
			vars[k] = v
		}
	}
	return vars
}
