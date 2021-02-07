package metrics

import (
	"regexp"
	"strings"
)

// DenyList encapsulates the logic needed to filter based on a string.
type DenyList struct {
	list  map[string]struct{}
	rList []*regexp.Regexp
}

// New constructs a new DenyList based on a white- and a
// DenyList. Only one of them can be not empty.
func New(b map[string]struct{}) (*DenyList, error) {
	black := copyList(b)
	var list map[string]struct{}
	list = black

	return &DenyList{
		list: list,
	}, nil
}

// Parse parses and compiles all of the regexes in the DenyList.
func (l *DenyList) Parse() error {
	var regexes []*regexp.Regexp
	for item := range l.list {
		r, err := regexp.Compile(item)
		if err != nil {
			return err
		}
		regexes = append(regexes, r)
	}
	l.rList = regexes
	return nil
}

// IsIncluded returns if the given item is included.
func (l *DenyList) IsIncluded(item string) bool {
	var matched bool
	for _, r := range l.rList {
		matched = r.MatchString(item)
		if matched {
			break
		}
	}

	return matched
}

// Status returns the status of the DenyList that can e.g. be passed into
// a logger.
func (l *DenyList) Status() string {
	items := []string{}
	for key := range l.list {
		items = append(items, key)
	}

	return "DenyListing the following items: " + strings.Join(items, ", ")
}

func copyList(l map[string]struct{}) map[string]struct{} {
	newList := map[string]struct{}{}
	for k, v := range l {
		newList[k] = v
	}
	return newList
}
