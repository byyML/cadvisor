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
	"regexp"
)

// DenyList encapsulates the logic needed to filter based on a string.
type DenyList struct {
	list  map[string]struct{}
	rList []*regexp.Regexp
}

// New constructs a new DenyList based on a white- and a
// DenyList. Only one of them can be not empty.
// New constructs a new DenyList based on a white- and a
// DenyList. Only one of them can be not empty.
func NewDenyList(b map[string]struct{}) (*DenyList, error) {
	black := copy(b)
	list := black
	l := &DenyList{
		list: list,
	}
	var regs []*regexp.Regexp
	for item := range l.list {
		r, err := regexp.Compile(item)
		if err != nil {
			return nil, err
		}
		regs = append(regs, r)
	}
	l.rList = regs
	return l, nil
}


// IsIncluded returns if the given item is included.
func (l *DenyList) IsDenied(item string) bool {
	var matched bool
	for _, r := range l.rList {
		matched = r.MatchString(item)
		if matched {
			break
		}
	}

	return matched
}

func copy(l map[string]struct{}) map[string]struct{} {
	newList := map[string]struct{}{}
	for k, v := range l {
		newList[k] = v
	}
	return newList
}
