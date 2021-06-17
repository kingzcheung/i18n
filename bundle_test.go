package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/kingzcheung/i18n/v2/testdata"
	"github.com/kingzcheung/i18n/v2/typ"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"reflect"
	"testing"
)

func Test_getTagFromFilepath(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want language.Tag
	}{
		{
			name: "zh",
			args: args{
				filepath: "/a/b/c/zh.json",
			},
			want: language.Chinese,
		},
		{
			name: "zh",
			args: args{
				filepath: "zh-CN.json",
			},
			want: language.Make("zh-CN"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTagFromFilepath(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTagFromFilepath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flatValues(t *testing.T) {
	type args struct {
		k string
		v interface{}
	}
	tests := []struct {
		name   string
		args   args
		wantKv map[string]string
	}{
		{
			name: "normal",
			args: args{
				k: "sha",
				v: "482713a2984e076e55164e3437cf410a28e80bf3",
			},
			wantKv: map[string]string{
				"sha": "482713a2984e076e55164e3437cf410a28e80bf3",
			},
		},
		{
			name: "deep",
			args: args{
				k: "tree",
				v: map[string]string{
					"path": "v2",
					"mode": "040000",
				},
			},
			wantKv: map[string]string{
				"tree.path": "v2",
				"tree.mode": "040000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if gotKv := flatValues(tt.args.k, tt.args.v); !reflect.DeepEqual(gotKv, tt.wantKv) {
				t.Errorf("flatValues() = %v, want %v", gotKv, tt.wantKv)

			} else {
				fmt.Println(gotKv)
			}

		})
	}
}

func TestBundle_LoadMessageFromBytes(t *testing.T) {
	data, err := testdata.TestDataFs.ReadFile("lang/en.json")
	assert.NoError(t, err)
	type fields struct {
		defaultLanguage language.Tag
		locales         []*locale
	}
	type args struct {
		s   []byte
		tag language.Tag
		umn typ.UnmarshalFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				defaultLanguage: language.English,
				locales:         nil,
			},
			args: args{
				s:   data,
				tag: language.English,
				umn: func(data []byte, i *map[string]interface{}) error {
					return json.Unmarshal(data, i)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bundle{
				defaultLanguage: tt.fields.defaultLanguage,
				locales:         tt.fields.locales,
			}
			if err := b.LoadMessageFromBytes(tt.args.s, tt.args.tag, tt.args.umn); (err != nil) != tt.wantErr {
				t.Errorf("LoadMessageFromBytes() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				for _, l := range b.locales {
					fmt.Printf("%+v\n", l)
				}
			}
		})
	}
}

func TestBundle_LoadMessageFromBytes_Ini(t *testing.T) {
	iniText := []byte(`
# possible values : production, development
app_mode = development

[paths]
# Path to where grafana can store temp files, sessions, and the sqlite3 db (if that is used)
data = /home/git/grafana

[server]
# Protocol (http or https)
protocol = http

# The http port  to use
http_port = 9999

# Redirect to correct domain if host header does not match domain
# Prevents DNS rebinding attacks
enforce_domain = true
`)
	type fields struct {
		defaultLanguage language.Tag
		locales         []*locale
	}
	type args struct {
		s   []byte
		tag language.Tag
		umn typ.UnmarshalFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				defaultLanguage: language.English,
				locales:         nil,
			},
			args: args{
				s:   iniText,
				tag: language.English,
				umn: func(bytes []byte, i *map[string]interface{}) error {
					cfg := typ.NewIni()
					return cfg.Unmarshal(bytes, i)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bundle{
				defaultLanguage: tt.fields.defaultLanguage,
				locales:         tt.fields.locales,
			}
			if err := b.LoadMessageFromBytes(tt.args.s, tt.args.tag, tt.args.umn); (err != nil) != tt.wantErr {
				t.Errorf("LoadMessageFromBytes() error = %v, wantErr %v", err, tt.wantErr)
			} else {

				for _, l := range b.locales {
					fmt.Printf("%+v\n", l)
				}
			}
		})
	}
}

func TestBundle_LoadMessageFromBytes_Yaml(t *testing.T) {
	yamlText, err := testdata.TestDataFs.ReadFile("en.yaml")
	assert.NoError(t, err)

	type fields struct {
		defaultLanguage language.Tag
		locales         []*locale
	}
	type args struct {
		s   []byte
		tag language.Tag
		umn typ.UnmarshalFunc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				defaultLanguage: language.English,
				locales:         nil,
			},
			args: args{
				s:   yamlText,
				tag: language.English,
				umn: func(bytes []byte, i *map[string]interface{}) error {
					return yaml.Unmarshal(bytes, i)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bundle{
				defaultLanguage: tt.fields.defaultLanguage,
				locales:         tt.fields.locales,
			}
			if err := b.LoadMessageFromBytes(tt.args.s, tt.args.tag, tt.args.umn); (err != nil) != tt.wantErr {
				t.Errorf("LoadMessageFromBytes() error = %v, wantErr %v", err, tt.wantErr)
			} else {

				for _, l := range b.locales {
					fmt.Printf("%+v\n", l)
				}
			}
		})
	}
}
