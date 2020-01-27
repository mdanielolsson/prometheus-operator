// Copyright 2020 The prometheus-operator Authors
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

package e2e

import (
	"testing"
)

func testTRCreateDeleteCluster(t *testing.T) {

	ctx := framework.NewTestCtx(t)
	defer ctx.Cleanup(t)
	ns := ctx.CreateNamespace(t, framework.KubeClient)
	ctx.SetupPrometheusRBAC(t, ns, framework.KubeClient)

	name := "test"

	if _, err := framework.CreateThanosRulerAndWaitUntilReady(ns, framework.MakeBasicThanosRuler(name)); err != nil {
		t.Fatal(err)
	}

	if err := framework.DeleteThanosRulerAndWaitUntilGone(ns, name); err != nil {
		t.Fatal(err)
	}
}
