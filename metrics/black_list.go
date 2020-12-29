package metrics

import (
	"regexp"
	"strings"
)

// BlackList encapsulates the logic needed to filter based on a string.
type BlackList struct {
	list        map[string]struct{}
	rList       []*regexp.Regexp
}

// New constructs a new BlackList based on a white- and a
// blacklist. Only one of them can be not empty.
func New(b map[string]struct{}) (*BlackList, error) {
	black := copyList(b)
	var list map[string]struct{}
	list = black

	return &BlackList{
		list:        list,
	}, nil
}

// Parse parses and compiles all of the regexes in the BlackList.
func (l *BlackList) Parse() error {
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
func (l *BlackList) IsIncluded(item string) bool {
	var matched bool
	for _, r := range l.rList {
		matched = r.MatchString(item)
		if matched {
			break
		}
	}

	return matched
}


// Status returns the status of the BlackList that can e.g. be passed into
// a logger.
func (l *BlackList) Status() string {
	items := []string{}
	for key := range l.list {
		items = append(items, key)
	}

	return "blacklisting the following items: " + strings.Join(items, ", ")
}

func copyList(l map[string]struct{}) map[string]struct{} {
	newList := map[string]struct{}{}
	for k, v := range l {
		newList[k] = v
	}
	return newList
}