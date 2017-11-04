package mainx

import (
	"encoding/json"
	"fmt"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/venicegeo/mpg-sandbox/rmscrub/scrubber"
)

// HandleA is a lambda
func HandleA(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	fmt.Print(evt)

	apiKey, err := scrubber.GetAPIKey()
	if err != nil {
		return nil, fmt.Errorf("Failed to get API key: %s", err)
	}

	projects, err := scrubber.NewProjectList(apiKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to get projects: %s", err)
	}

	issues, err := scrubber.NewIssueList(apiKey, projects)
	if err != nil {
		return nil, fmt.Errorf("Failed to get issues: %s", err)
	}

	scrubReport := scrubber.NewScrubReport(issues)
	err = scrubReport.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to run ScrubReport: %s", err)
	}
	scrubberResults := scrubReport.Report()

	return scrubberResults.String(), nil
}
