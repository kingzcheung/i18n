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
