/*
Copyright 2021 The Flux authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package repository

import (
	"strings"

	helmreg "helm.sh/helm/v3/pkg/registry"
)

// NormalizeURL normalizes a ChartRepository URL by its scheme.
func NormalizeURL(repositoryURL string) string {
	if repositoryURL == "" {
		return ""
	}

	if strings.Contains(repositoryURL, helmreg.OCIScheme) {
		return strings.TrimRight(repositoryURL, "/")
	}

	return strings.TrimRight(repositoryURL, "/") + "/"

}
