// Copyright © 2019 Antoine Cordier <com.acordier@gmail.com>
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

package cmd

import (
	"testing"
)

func Test_legalize(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			"leading number should be removed",
			args{"1resources"},
			"resources",
			false,
		},
		{
			"name should be normalized to camel case",
			args{"resources_directory"},
			"resourcesDirectory",
			false,
		},
		{
			"name should be normalized to camel case",
			args{"resources directory"},
			"resourcesDirectory",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := legalize(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("legalize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("legalize() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_fileID(t *testing.T) {
	type args struct {
		path string
		src  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"parent directory should be removed",
			args{ "resources/svg/play.svg", "resources" },
			"svg/play.svg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileID(tt.args.path, tt.args.src); got != tt.want {
				t.Errorf("fileID() = %v, want %v", got, tt.want)
			}
		})
	}
}
