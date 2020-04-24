i18n 
---
i18n 通过解析json实现国际化和本地化,也支持通过 `map[string]map[string][string]` 数据（通过工具生成）。

### 安装
```
go get github.com/kingzcheung/i18n
```
### 用法

导入包
```go
import "github.com/kingzcheung/i18n"
```

```go
    filePath := "/your/project/lang"
	err := i18n.LoadPath(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	i18n.DefaultLanguage("en")
	i18n.I("Africa") //output: Africa
```

LoadMap :
```go
var locale = map[string]map[string]string{
		"en": {
			"key_already_exists": "KEY [%s] already exists",
		},
	}
	LoadMap(locale)
	DefaultLanguage("en")
    fmt.Println(I("key_already_exists","bar")) //output: KEY [bar] already exists

```