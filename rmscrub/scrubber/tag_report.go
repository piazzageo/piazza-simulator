package scrubber

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// TagReport verfiies and counts the subject/title prefixes
type TagReport struct {
	list   *IssueList
	counts map[string]map[string][]*Issue // tag -> (status -> []Issues)
}

// TagResults is what gets stored
type TagResults struct {
	Date                time.Time        `json:"date"`
	Data                map[string][]int `json:"data"`
	AtoCount            int              `json:"ato_count"`
	AtoEngineeringCount int              `json:"ato_engineering"`
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

// Report reports
func (t *TagReport) Report() *TagResults {
	result := &TagResults{
		Date: time.Now(),
		Data: map[string][]int{},
	}

	atoCount := 0
	atoEngineeringCount := 0

	for _, tag := range TitleTags {
		if isATOTitleTag(tag) {
			statusMap := t.counts[tag]
			new := statusMap["New"]
			inprog := statusMap["In Progress"]
			atoCount += len(new) + len(inprog)

			if len(new)+len(inprog) > 0 {
				result.Data[tag] = []int{}
				for _, issue := range new {
					result.Data[tag] = append(result.Data[tag], issue.ID)
				}
				for _, issue := range inprog {
					result.Data[tag] = append(result.Data[tag], issue.ID)
				}
			}

			if isATOEngineeringTitleTag(tag) {
				atoEngineeringCount += len(new) + len(inprog)
			}
		}
	}

	result.AtoCount = atoCount
	result.AtoEngineeringCount = atoEngineeringCount

	return result
}

func (r *TagResults) String() string {
	s := ""

	for tag, ids := range r.Data {
		s += fmt.Sprintf("%s: %d", tag, len(ids))
		/*for _, id := range ids {
			s += fmt.Sprintf(" %d", id)
		}*/
		s += fmt.Sprintf("\n")
	}

	//s += fmt.Sprintf("Date: %s\n", r.Date.Format(time.RFC3339))
	s += fmt.Sprintf("Engineering issues: %d\n", r.AtoEngineeringCount)
	s += fmt.Sprintf("Total issues: %d\n", r.AtoCount)

	return s
}

// Store writes to the database
func (r *TagResults) Store() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(r)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("RedmineStatusTable"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		return err
	}

	return nil
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
