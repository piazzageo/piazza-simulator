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
	"log"
	"os"
)

const CURRENT_SPRINT = "Sprint 07"
const BACKLOG_SPRINT = "Backlog"
const EPIC_SPRINT = "Pz Epics"

var PastSprints = []string{
	"Sprint 0",
	"Sprint 01",
	"Sprint 02",
	"Sprint 03",
	"Sprint 04",
	"Sprint 05",
	"Sprint 06",
}

var FutureSprints = []string{
	"Sprint 08",
	"Sprint 09",
}

func Errorf(mssg string, args ...interface{}) {
	s := fmt.Sprintf(mssg, args...)
	fmt.Printf("error: %s\n", s)
	if DEBUG {
		panic(1)
	}
	os.Exit(1)
}

func Logf(mssg string, args ...interface{}) {
	if DEBUG {
		log.Printf(mssg, args...)
	}
}

func Printf(mssg string, args ...interface{}) {
	fmt.Printf(mssg, args...)
	fmt.Printf("\n")
}
