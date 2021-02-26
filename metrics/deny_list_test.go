// Copyright 2021 Google Inc. All Rights Reserved.
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

package metrics

import (
	//"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	//"os"
	"testing"
)

var (
	ignoreSpecificMetrics = []string{"^machine_(memory|cpu).*", "^container_fs_*", "^container_cpu_*"}
)

func TestNewDenyList(t *testing.T) {
	denyList, err := NewDenyList(ignoreSpecificMetrics)
	assert.Nil(t, err)

	testDenyListIsDenied(t, denyList, "testdata/prometheus_all_metrics_name")

}

func testDenyListIsDenied(t *testing.T, denyList *DenyList, metricsFile string) {
	//wantMetrics, err := os.Open(metricsFile)
	//if err != nil {
	//	t.Fatalf("unable to read input test file %s", metricsFile)
	//}
	//
	//assert.Equal()
}

