// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"
	"strconv"

	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)

var _ fn.Runner = &YourFunction{}

// TODO: Change to your functionConfig "Kind" name.
type YourFunction struct {
	Data map[string]string `json:"data,omitempty"`
}

// Run is the main function logic.
// `items` is parsed from the STDIN "ResourceList.Items".
// `functionConfig` is from the STDIN "ResourceList.FunctionConfig". The value has been assigned to the r attributes
// `results` is the "ResourceList.Results" that you can write result info to.
func (r *YourFunction) Run(ctx *fn.Context, functionConfig *fn.KubeObject, items fn.KubeObjects, results *fn.Results) bool {

	_, found, _ := functionConfig.NestedStringMap("data")
	if !found {
		*results = append(*results, fn.GeneralResult("override_values not provided in fn-config", fn.Error))
		return false
	}

	for _, kubeObject := range items {
		if kubeObject.IsGVK("apps", "v1", "Deployment") {
			if val, found, _ := functionConfig.NestedString("data", "replicas"); found {
				replicas, err := strconv.Atoi(val)
				if err != nil {
					*results = append(*results, fn.GeneralResult("replicas value is not a valid integer", fn.Error))
					return false
				}
				kubeObject.SetNestedField(replicas, "spec", "replicas")
				*results = append(*results, fn.GeneralResult("Successfully mutated deployments", fn.Info))
			}
		}
	}
	return true
}

func main() {
	runner := fn.WithContext(context.Background(), &YourFunction{})
	if err := fn.AsMain(runner); err != nil {
		os.Exit(1)
	}
}
