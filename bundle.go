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

package i18n

import (
	"fmt"
	"github.com/kingzcheung/i18n/v3/testdata"
	"github.com/kingzcheung/i18n/v3/typ"
	"golang.org/x/text/language"
	"io/fs"
	"io/ioutil"
	"path"
	"strings"
)

type locale struct {
	tag     language.Tag
	message map[string]string
}

type Bundle struct {
	defaultLanguage language.Tag
	locales         []*locale
}

func (b *Bundle) DefaultLanguage() language.Tag {
	return b.defaultLanguage
}

func (b *Bundle) Locales() []*locale {
	return b.locales
}

func (b *Bundle) addLocale(locale *locale) {
	b.locales = append(b.locales, locale)
}

func (b *Bundle) addLocaleMessage(tag language.Tag, message map[string]string) {
	l := new(locale)
	l.tag = tag
	l.message = message
	b.addLocale(l)
}

func (b *Bundle) LoadMessageFromFile(filename string, umn typ.UnmarshalFunc) error {
	tag := getTagFromFilepath(filename)

	rf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return b.LoadMessageFromBytes(rf, tag, umn)
}

func (b *Bundle) LoadMessageFromDir(dir string, umn typ.UnmarshalFunc) error {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, info := range infos {
		err = b.LoadMessageFromFile(path.Join(dir, info.Name()), umn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bundle) LoadMessageFromFsEntries(entries []fs.DirEntry, umn typ.UnmarshalFunc) error {
	for _, entry := range entries {
		data, err := testdata.TestDataFs.ReadFile(path.Join("lang", entry.Name()))
		if err != nil {
			return err
		}
		name := strings.Replace(entry.Name(), path.Ext(entry.Name()), "", 1)
		err = b.LoadMessageFromBytes(data, language.Make(name), umn)
		if err != nil {
			return err
		}
	}
	return nil
}

//getTagFromFilepath 从文件中获取语言特征
func getTagFromFilepath(filepath string) language.Tag {
	_, filename := path.Split(filepath)
	ext := path.Ext(filepath)

	name := strings.ReplaceAll(filename, ext, "")
	if name == "" {
		return language.English
	}
	return language.Make(name)
}

func (b *Bundle) LoadMessageFromBytes(s []byte, tag language.Tag, umn typ.UnmarshalFunc) error {
	msg, err := b.ParseMessageFromBytes(s, umn)
	if err != nil {
		return err
	}

	b.addLocaleMessage(tag, msg)

	return nil
}

func (b *Bundle) LoadMessageFromString(s string, tag language.Tag, umn typ.UnmarshalFunc) error {
	return b.LoadMessageFromBytes([]byte(s), tag, umn)
}

func (b *Bundle) ParseMessageFromBytes(rf []byte, umn typ.UnmarshalFunc) (map[string]string, error) {
	var m map[string]interface{}
	err := umn.Unmarshal(rf, &m)
	if err != nil {
		return nil, err
	}
	var data = make(map[string]string)
	for k, v := range m {
		kv := flatValues(k, v)
		for s, s2 := range kv {
			data[s] = s2
		}
	}

	return data, nil
}

//{
//  "sha": "482713a2984e076e55164e3437cf410a28e80bf3",
//  "url": "https://api.github.com",
//  "tree":
//    {
//      "path": "v2",
//      "mode": "040000"
//    },
//  "truncated": false
//}
// "sha":"482713a2984e076e55164e3437cf410a28e80bf3"
// "url":"https://api.github.com"
//"tree.path":"v2"
//"tree.mode":"040000"
//"truncated":false
func flatValues(k string, v interface{}) (kv map[string]string) {
	kv = make(map[string]string)
	switch v.(type) {
	case string:
		kv = map[string]string{k: v.(string)}
		return
	case bool:
		var value = "false"
		if v.(bool) {
			value = "true"
		}
		kv[k] = value
		return
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		kv[k] = fmt.Sprintf("%v", v)
		return

	case map[string]interface{}:
		for s, i := range v.(map[string]interface{}) {
			m := flatValues(k+"."+s, i)
			for s2, s3 := range m {
				kv[s2] = s3
			}
		}
		return
	case map[string]string:
		for s, i := range v.(map[string]string) {
			m := flatValues(k+"."+s, i)
			for s2, s3 := range m {
				kv[s2] = s3
			}
		}
	case map[interface{}]interface{}:
		for s, i := range v.(map[interface{}]interface{}) {
			m := flatValues(k+"."+s.(string), i)
			for s2, s3 := range m {
				kv[s2] = s3
			}
		}
	}

	return kv
}

func NewBundle(defaultLanguage language.Tag) *Bundle {
	return &Bundle{defaultLanguage: defaultLanguage}
}
