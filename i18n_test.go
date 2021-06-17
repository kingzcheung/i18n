package i18n

import (
	"encoding/json"
	"github.com/kingzcheung/i18n/v2/testdata"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

func TestLocalization_Localize(t *testing.T) {
	type fields struct {
		bundle *Bundle
		tags   []language.Tag
	}
	type args struct {
		key       string
		variables []map[string]interface{}
	}
	bundle := NewBundle(language.English)
	entries, err := testdata.TestDataFs.ReadDir("lang")
	assert.NoError(t, err)
	err = bundle.LoadMessageFromFsEntries(entries, func(data []byte, m *map[string]interface{}) error {
		return json.Unmarshal(data, m)
	})
	assert.NoError(t, err)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "无参数",
			fields: fields{
				bundle,
				nil,
			},
			args: args{
				key:       "Africa",
				variables: nil,
			},
			want: "Africa, ",
		},
		{
			name: "有参数",
			fields: fields{
				bundle,
				nil,
			},
			args: args{
				key: "Africa",
				variables: []map[string]interface{}{
					{
						"welcome": "8RQ1qd",
					},
				},
			},
			want: "Africa, 8RQ1qd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Localization{
				bundle:       tt.fields.bundle,
				fallbackTags: tt.fields.tags,
			}
			if got := l.Localize(tt.args.key, tt.args.variables...); got != tt.want {
				t.Errorf("Localize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalization_With_Localize(t *testing.T) {
	type fields struct {
		bundle *Bundle
		tags   []language.Tag
	}
	type args struct {
		key       string
		variables []map[string]interface{}
		with      string
	}
	bundle := NewBundle(language.English)
	entries, err := testdata.TestDataFs.ReadDir("lang")
	assert.NoError(t, err)
	err = bundle.LoadMessageFromFsEntries(entries, func(data []byte, m *map[string]interface{}) error {
		return json.Unmarshal(data, m)
	})
	assert.NoError(t, err)
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "cn",
			fields: fields{
				bundle,
				nil,
			},
			args: args{
				key:       "Antarctica",
				variables: nil,
				with:      "zh-CN",
			},
			want: "南极洲",
		},
		{
			name: "hk",
			fields: fields{
				bundle,
				nil,
			},
			args: args{
				key:       "Antarctica",
				variables: nil,
				with:      "zh-HK",
			},
			want: "南極洲",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Localization{
				bundle:       tt.fields.bundle,
				fallbackTags: tt.fields.tags,
			}
			if got := l.With(tt.args.with).Localize(tt.args.key, tt.args.variables...); got != tt.want {
				t.Errorf("Localize() = %v, want %v", got, tt.want)
			}
		})
	}
}
