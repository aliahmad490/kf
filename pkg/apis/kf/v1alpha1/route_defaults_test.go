// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/kf/pkg/kf/testutil"
)

func TestRoute_SetDefaults_DedupeAppNames(t *testing.T) {
	t.Parallel()

	r := &Route{
		Spec: RouteSpec{
			AppNames: []string{
				"d", "a", "d", "a", "b", "c",
			},
		},
	}

	r.SetDefaults(context.Background())

	testutil.AssertEqual(t, "len", 4, len(r.Spec.AppNames))
	testutil.AssertContainsAll(t, strings.Join(r.Spec.AppNames, ""), []string{"a", "b", "c", "d"})
}

func ExampleRoute_SetDefaults_prefixRoutes() {
	r := &Route{}
	r.Spec.Path = "some-path"
	r.SetDefaults(context.Background())

	fmt.Println("Route:", r.Spec.Path)

	// Output: Route: /some-path
}
