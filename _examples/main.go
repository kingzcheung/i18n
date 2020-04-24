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
	"github.com/kingzcheung/i18n"
	"log"
)

func main() {
	//lang
	//├── en.json
	//├── zh-CN.json
	//└── zh-HK.json
	filePath := "/your/project/lang"
	err := i18n.LoadPath(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	i18n.DefaultLanguage("en")
	i18n.I("Africa") //output: Africa
}

func loadMap() {
	i18n.LoadMap(map[string]map[string]string{
		"en": {
			"foo": "bar",
		},
		"zh-CN": {
			"foo": "内容",
		},
	})
}
