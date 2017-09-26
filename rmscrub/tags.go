package main

import (
	"fmt"
	"sort"
	"strings"
)

type TagChecker struct {
	counts map[string]map[string]int // tag -> status -> int
}

func NewTagCheker() *TagChecker {
	t := &TagChecker{}

	//	statusMap := map[string]int{}
	t.counts = map[string]map[string]int{}

	return t
}

func (t *TagChecker) Report() {
	sort.Strings(TitleTags)

	for _, tag := range TitleTags {
		statusMap := t.counts[tag]

		new := statusMap["New"]
		inprog := statusMap["In Progress"]
		if new+inprog > 0 {
			fmt.Printf("%s:\n", tag)
			fmt.Printf("\tNew:         %d\n", new)
			fmt.Printf("\tIn Progress: %d\n", inprog)
		}
	}
}

func (t *TagChecker) Run(list *IssueList) {

	for _, issue := range list.GetMap() {

		subj := issue.Subject
		status := issue.Status.Name

		for _, tag := range TitleTags {
			if strings.Contains(subj, tag) {

				_, ok := t.counts[tag]
				if !ok {
					t.counts[tag] = map[string]int{}
				}
				t.counts[tag][status]++
				break
			}
		}
	}

}
