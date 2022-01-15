// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package credentials

// SMTP represents SMTP credentials.
type SMTP struct {
	Username string `json:"username,omitempty" xml:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" xml:"password,omitempty" yaml:"password,omitempty"`
	Address  string `json:"address,omitempty" xml:"address,omitempty" yaml:"address,omitempty"`
}
