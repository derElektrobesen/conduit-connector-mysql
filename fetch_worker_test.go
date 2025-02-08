// Copyright © 2025 Meroxa, Inc.
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

package mysql

import (
	"testing"

	gover "github.com/hashicorp/go-version"
	"github.com/matryer/is"
)

func TestReadOnlyTransactionsAllowed(t *testing.T) {
	tests := []struct {
		version string
		want    bool
	}{
		{"5.5.62-log", false},
		{"5.7.0", true},
		{"5.7.44-log", true},
		{"8.0.41", true},
		{"8.0", true},
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			is := is.New(t)

			ver, err := gover.NewVersion(tt.version)
			is.NoErr(err)

			is.Equal(readOnlyTransactionsAllowed(ver), tt.want)
		})
	}
}
