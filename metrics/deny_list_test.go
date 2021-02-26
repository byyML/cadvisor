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
	"bytes"
	"github.com/google/cadvisor/container"
	info "github.com/google/cadvisor/info/v1"
	v2 "github.com/google/cadvisor/info/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var (
	ignoreSpecificMetrics = []string{"^machine_(memory|cpu).*", "^container_fs_*", "^container_cpu_*", "^container_blkio.*"}
)

func TestNewDenyList(t *testing.T) {
	denyList, _ := NewDenyList(ignoreSpecificMetrics)
	c := NewPrometheusCollector(testSubcontainersInfoProvider{}, func(container *info.ContainerInfo) map[string]string {
		s := DefaultContainerLabels(container)
		s["zone.name"] = "hello"
		return s
	}, container.AllMetrics, now, v2.RequestOptions{}, denyList)
	reg := prometheus.NewRegistry()
	reg.MustRegister(c)

	testDenyListIsDenied(t, reg, "testdata/prometheus_metrics_denylist")

}

func testDenyListIsDenied(t *testing.T, gatherer prometheus.Gatherer, metricsFile string) {
	wantMetrics, err := os.Open(metricsFile)
	if err != nil {
		t.Fatalf("unable to read input test file %s", metricsFile)
	}

	err = testutil.GatherAndCompare(gatherer, wantMetrics)
	if err != nil {
		t.Fatalf("Metric comparison failed: %s", err)
	}
}

func TestNewDenyListWithMachine(t *testing.T) {
	denyList, _ := NewDenyList(ignoreSpecificMetrics)
	collector := NewPrometheusMachineCollector(testSubcontainersInfoProvider{}, container.AllMetrics, denyList)
	registry := prometheus.NewRegistry()
	registry.MustRegister(collector)

	metricsFamily, err := registry.Gather()
	assert.Nil(t, err)

	var metricBuffer bytes.Buffer
	for _, metricFamily := range metricsFamily {
		_, err := expfmt.MetricFamilyToText(&metricBuffer, metricFamily)
		assert.Nil(t, err)
	}
	collectedMetrics := metricBuffer.String()

	expectedMetrics, err := ioutil.ReadFile("testdata/prometheus_machine_metrics_denylist")
	assert.Nil(t, err)
	assert.Equal(t, string(expectedMetrics), collectedMetrics)
}
