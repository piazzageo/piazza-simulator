package scrubber

import (
	"fmt"
	"sort"
	"strings"
)

// TagReport verfiies and counts the subject/title prefixes
type TagReport struct {
	list   *IssueList
	counts map[string]map[string][]*Issue // tag -> (status -> []Issues)
}

// NewTagReport makes a new TagReport
func NewTagReport(list *IssueList) *TagReport {
	t := &TagReport{
		list: list,
	}

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

// Report reports
func (t *TagReport) Report() string {
	result := ""

	sort.Strings(TitleTags)

	atoCount := 0
	atoEngineeringCount := 0

	result += fmt.Sprintf("```\n")

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

	result += fmt.Sprintf("There are %d open ATO tickets, of which %d are engineering-related.\n",
		atoCount, atoEngineeringCount)
	result += fmt.Sprintf("```\n")

	return result
}

// Run runs
func (t *TagReport) Run() error {

	for _, issue := range t.list.GetMap() {

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

	return nil
}
