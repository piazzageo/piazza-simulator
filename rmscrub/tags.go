package main

import (
	"fmt"
	"sort"
	"strings"
)

// TagChecker verfiies and counts the subject/title prefixes
type TagChecker struct {
	counts map[string]map[string][]*Issue // tag -> (status -> []Issues)
}

// NewTagCheker makes a new TagChecker
func NewTagCheker() *TagChecker {
	t := &TagChecker{}

	//	statusMap := map[string]int{}
	t.counts = map[string]map[string][]*Issue{}

	return t
}

func arrayString(issues []*Issue) string {
	s := ""
	for _, issue := range issues {
		s += fmt.Sprintf("%d ", issue.ID)
	}
	return s
}

func (t *TagChecker) report() {
	sort.Strings(TitleTags)

	atoCount := 0
	atoEngineeringCount := 0

	fmt.Printf("```\n")

	for _, tag := range TitleTags {
		if isATOTitleTag(tag) {
			statusMap := t.counts[tag]
			new := statusMap["New"]
			inprog := statusMap["In Progress"]
			atoCount += len(new) + len(inprog)

			if len(new)+len(inprog) > 0 {
				fmt.Printf("%-20sNew           (%d)   %s\n", tag, len(new), arrayString(new))
				fmt.Printf("%-20sIn Progress   (%d)   %s\n", tag, len(inprog), arrayString(inprog))
			}

			if isATOEngineeringTitleTag(tag) {
				atoEngineeringCount += len(new) + len(inprog)
			}
		}
	}

	fmt.Printf("```\n")

	fmt.Printf("There are %d open ATO tickets, of which %d are engineering-related.\n",
		atoCount, atoEngineeringCount)
}

func (t *TagChecker) run(list *IssueList) {

	for _, issue := range list.getMap() {

		if issue.isEpic() {
			continue
		}

		subj := issue.Subject
		status := issue.Status.Name

		for _, tag := range TitleTags {
			if strings.Contains(subj, tag) {

				_, ok := t.counts[tag]
				if !ok {
					t.counts[tag] = map[string][]*Issue{}
				}
				t.counts[tag][status] = append(t.counts[tag][status], issue)

				break
			}
		}
	}

}
