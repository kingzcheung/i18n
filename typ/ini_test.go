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

package typ

import (
	"fmt"
	"github.com/kingzcheung/i18n/v2/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIni_Unmarshal(t *testing.T) {
	iniBytes, err := testdata.TestDataFs.ReadFile("en.ini")
	assert.NoError(t, err)
	type args struct {
		data []byte
		v    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "ini parse",
			args: args{
				data: iniBytes,
				v:    map[string]interface{}{},
			},
			want: map[string]interface{}{
				"app_mode":              "development",
				"paths.data":            "/home/git/grafana",
				"server.db.host":        "localhost",
				"server.enforce_domain": "true",
				"server.http_port":      "9999",
				"server.protocol":       "http",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Ini{}
			if err := i.Unmarshal(tt.args.data, &tt.args.v); err != nil {
				t.Errorf("Unmarshal() error = %v, want %v", err, tt.want)
			} else {
				assert.Equal(t, 6, len(tt.args.v))
				assert.EqualValues(t, tt.want, tt.args.v)
				fmt.Println(tt.args.v)
			}
		})
	}
}
