i18n 
---
[![Build Status](https://cloud.drone.io/api/badges/kingzcheung/i18n/status.svg)](https://cloud.drone.io/kingzcheung/i18n)
[![GitHub license](https://img.shields.io/badge/license-Apache-blue.svg)](https://github.com/kingzcheung/i18n/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kingzcheung/i18n.svg)](https://pkg.go.dev/github.com/kingzcheung/i18n)

> v2 版本开发中

## Getting started

Installation:

```shell
go get -u github.com/kingzcheung/i18n/v2
```

example:

```go
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
```

