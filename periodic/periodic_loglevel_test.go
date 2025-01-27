// Copyright 2017 Istio Authors
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

//go:build !race
// +build !race

package periodic

import (
	"testing"

	"fortio.org/log"
)

// Rerun some test with various log level for coverage of the print statements
// TODO: golden copy type check of output ?

func TestQuietMode(t *testing.T) {
	log.SetLogLevel(log.Error)
	TestStart(t)
	TestExactlyMaxQps(t)
	log.SetLogLevel(log.Verbose)
	TestStart(t)
}
