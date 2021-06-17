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

package main

import (
	"encoding/json"
	"fmt"
	"github.com/kingzcheung/i18n/v2"
	"github.com/kingzcheung/i18n/v2/testdata"
	"golang.org/x/text/language"
)

func main() {

	bundle := i18n.NewBundle(language.English)
	//err := bundle.LoadMessageFromFile("lang/en.json", func(bytes []byte, m *map[string]interface{}) error {
	//	return json.Unmarshal(bytes,m)
	//})
	//if err != nil {
	//	panic(err)
	//}
	data, err := testdata.TestDataFs.ReadFile("lang/en.json")
	if err != nil {
		panic(err)
	}
	err = bundle.LoadMessageFromBytes(data, language.English, func(bytes []byte, m *map[string]interface{}) error {
		return json.Unmarshal(bytes, m)
	})
	if err != nil {
		panic(err)
	}

	// or
	entries, err := testdata.TestDataFs.ReadDir("lang")
	if err != nil {
		panic(err)
	}
	_ = bundle.LoadMessageFromFsEntries(entries, func(bytes []byte, m *map[string]interface{}) error {
		return json.Unmarshal(bytes, m)
	})

	e := i18n.NewLocalization(bundle, "zh-CN")

	fmt.Println(e.Localize("Antarctica"))                                      //南极洲
	fmt.Println(e.Localize("Africa", map[string]interface{}{"welcome": "你好"})) //非洲,你好
	fmt.Println(e.With("zh-HK").Localize("Asia"))                              //亞洲
	fmt.Println(e.With("en").Localize("Asia"))
}
