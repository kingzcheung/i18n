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
