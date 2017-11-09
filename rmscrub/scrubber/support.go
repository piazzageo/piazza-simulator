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

package scrubber

import (
	"fmt"
	"os"
	"strings"
)

// TargetProjects are the ones we will be scrubbing
var TargetProjects = []string{"Piazza", "Beachfront"}

// names of special sprints
const (
	CurrentSprint = "Sprint 46"
	ReadySprint   = "Ready"
	BacklogSprint = "Backlog"
	EpicSprint1   = "Pz Epics"
	EpicSprint2   = "BF Epics"
)

// FutureSprints are the things to come (but not used now)
var FutureSprints = []string{"Sprint 47"}

// PastSprints is names of sprints gone by
var PastSprints = []string{
	"Brisket",
	"Deckle",
	"xRejected",
	"Sprint 0",
	"Sprint 01",
	"Sprint 02",
	"Sprint 03",
	"Sprint 04",
	"Sprint 05",
	"Sprint 06",
	"Sprint 07",
	"Sprint 08",
	"Sprint 09",
	"Sprint 10",
	"Sprint 11",
	"Sprint 12",
	"Sprint 13",
	"Sprint 14",
	"Sprint 15",
	"Sprint 16",
	"Sprint 17",
	"Sprint 18",
	"Sprint 19",
	"Sprint 20",
	"Sprint 21",
	"Sprint 22",
	"Sprint 23",
	"Sprint 24",
	"Sprint 25",
	"Sprint 26",
	"Sprint 27",
	"Sprint 28",
	"Sprint 29",
	"Sprint 30",
	"Sprint 31",
	"Sprint 32",
	"Sprint 33",
	"Sprint 34",
	"Sprint 35",
	"Sprint 36",
	"Sprint 37",
	"Sprint 38",
	"Sprint 39",
	"Sprint 40",
	"Sprint 41",
	"Sprint 42",
	"Sprint 43",
	"Sprint 44",
	"Sprint 45",
}

// TitleTags are prefixes of issue subject/title
var TitleTags = []string{
	"[PP]",
	"[ATO Engineering]",
	"[Engineering]",
	"[ATO Management]",
	"[Management]",
	"[Testing]",
	"[ATO Testing]",
	"[DevOps]",
	"[ATO DevOps]",
	"[Security]",
	"[ATO Security]",
	///
	"[NotATO Engineering]",
	"[ATO]",
	"[2.0]",
}

func isATOEngineeringTitleTag(tag string) bool {
	switch tag {
	case "[ATO Engineering]":
		fallthrough
	case "[ATO Testing]":
		fallthrough
	case "[ATO DevOps]":
		return true
	}
	return false
}

func isATOTitleTag(tag string) bool {
	return strings.HasPrefix(tag, "[ATO")
}

// GetAPIKey returns the API key
func GetAPIKey() (string, error) {
	key := os.Getenv("RMKEY")
	if key == "" {
		return "", fmt.Errorf("$RMKEY not found")
	}
	return key, nil
}
