// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package prepare

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	surnamePrecedes = regexp.MustCompile(`(,?[\s?a-zA-Z\.]{1,})$`)
)

// ReorderSDNName will take a given SDN name and if it matches a specific pattern where
// the first name is placed after the last name (surname) to return a string where the first name
// preceedes the last.
//
// Example:
// SDN EntityID: 19147 has 'FELIX B. MADURO S.A.'
// SDN EntityID: 22790 has 'MADURO MOROS, Nicolas'
func ReorderSDNName(name string, tpe string) string {
	if !strings.EqualFold(tpe, "individual") {
		return name // only reorder individual names
	}
	v := surnamePrecedes.FindString(name)
	if v == "" {
		return name // no match on 'Doe, John'
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", strings.TrimPrefix(v, ","), strings.TrimSuffix(name, v)))
}
