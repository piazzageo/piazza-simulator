/*
Copyright 2016, RadiantBlue Technologies, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/venicegeo/mpg-sandbox/rmscrub/scrubber"
)

// Errorf records an error message
func Errorf(mssg string, args ...interface{}) {
	s := fmt.Sprintf(mssg, args...)
	fmt.Printf("error: %s\n", s)
	os.Exit(1)
}

func main() {

	//	if len(os.Args) == 2 {

	apiKey, err := scrubber.GetAPIKey()
	if err != nil {
		Errorf("Failed to get API key: %s", err)
	}

	projects, err := scrubber.NewProjectList(apiKey)
	if err != nil {
		Errorf("Failed to get projects: %s", err)
	}

	issues, err := scrubber.NewIssueList(apiKey, projects)
	if err != nil {
		Errorf("Failed to get issues: %s", err)
	}

	scrubReport := scrubber.NewScrubReport(issues)
	err = scrubReport.Run()
	if err != nil {
		Errorf("Failed to run ScrubReport: %s", err)
	}
	scrubberResults := scrubReport.Report()
	fmt.Printf(scrubberResults.String())

	fmt.Printf("\n")

	tagReport := scrubber.NewTagReport(issues)
	err = tagReport.Run()
	if err != nil {
		Errorf("Failed to run TagReport: %s", err)
	}
	tagResults := tagReport.Report()
	fmt.Printf(tagResults.String())
}
