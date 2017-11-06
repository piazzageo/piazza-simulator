package main

import (
	"encoding/json"
	"fmt"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/venicegeo/mpg-sandbox/rmscrub/scrubber"
)

// RunTagger is a lambda
func RunTagger(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
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

	tagReport := scrubber.NewTagReport(issues)
	err = tagReport.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to run TagReport: %s", err)
	}
	tagResults := tagReport.Report()

	err = tagResults.Store()
	if err != nil {
		return nil, fmt.Errorf("Failed to store TagReport: %s", err)
	}

	return tagResults.String(), nil
}
