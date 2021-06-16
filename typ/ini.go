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
	"gopkg.in/ini.v1"
)

type Ini struct{}

func NewIni() Unmarshaler {
	return &Ini{}
}

func (i *Ini) Unmarshal(data []byte, v *map[string]interface{}) error {
	load, err := ini.Load(data)
	if err != nil {
		return err
	}
	load.BlockMode = false

	var m = make(map[string]interface{})

	for _, ss := range load.SectionStrings() {
		sec, err := load.GetSection(ss)
		if err != nil {
			return err
		}
		for _, s := range sec.KeyStrings() {
			if ss == ini.DefaultSection {
				m[s] = sec.Key(s).String()
				continue
			}
			k := ss + "." + s
			m[k] = sec.Key(s).String()
		}
	}
	*v = m
	return nil
}
