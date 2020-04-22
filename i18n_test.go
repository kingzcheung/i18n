package i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	input string
	ok    bool
	out   string
}

func TestLoad(t *testing.T) {
	as := assert.New(t)
	data := []testCase{
		{"lang/zh-CN.json", true, "zh-CN"},
		{"lang/zh-CN", false, ""},
		{"lang/test.cc", false, ""},
	}
	for _, d := range data {
		as.Equal(Load(d.input) == nil, d.ok)
	}

}

func Test_getLangFromPath(t *testing.T) {
	as := assert.New(t)
	data := []testCase{
		{"/i18n/lang/zh-CN.json", true, "zh-CN"},
		{"/i18n/lang/zh-CN", true, ""},
		{"/i18n/lang/test.go", true, "test"},
	}
	for _, d := range data {
		as.Equal(getLangFromPath(d.input), d.out)
	}
}

func TestLoadPath(t *testing.T) {
	as := assert.New(t)
	filePath := "./lang"
	err := LoadPath(filePath)
	if err != nil {
		as.Error(err)
	}
	as.Equal(len(locales.locales), 3)
}

func TestLang(t *testing.T) {
	as := assert.New(t)
	err := LoadPath("./lang")
	if err != nil {
		as.Error(err)
	}
	as.Equal(Lang("en", "Africa", ""), "Africa")
	as.Equal(Lang("en", "foo", ""), "foo")
	as.Equal(Lang("en", "foo", "valueFoo"), "valueFoo")
}

func TestI(t *testing.T) {
	as := assert.New(t)
	err := LoadPath("./lang")
	if err != nil {
		as.Error(err)
	}
	DefaultLanguage("zh-CN")
	as.Equal(I("foo"), "foo")
	as.Equal(I("Africa"), "非洲")
}

func BenchmarkLang(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = LoadPath("./lang")

		DefaultLanguage("zh-CN")
		Lang("en", "Africa", "")

	}
}
