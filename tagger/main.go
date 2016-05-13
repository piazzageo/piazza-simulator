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

// This app reads in a CSV export of a Redmine project and "validates" all
// the issues based on various rules such as:
//    a Story must have a parent, and that parent must be an Epic
//    a Task in the current milestone must have an estimate, and that
//      estimate must be <= 16 hours
// This app is still underdevelopment -- will be adding rules as we need them.
//
// To use:
//   - in Redmine, click on "View all issues" (right-hand column, top)
//   - then click on "Also available in... CSV" (main panel, bottom-right)
//   - run the app:  $ go run redmine-scrub.go downloadedfile.csv
//

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

var DEBUG = false

func getApiKey() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", fmt.Errorf("$HOME not found")
	}

	key, err := ioutil.ReadFile(home + "/.git-token")
	if err != nil {
		return "", err
	}

	s := strings.TrimSpace(string(key))
	//log.Printf("API Key: %s", s)

	return s, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage:  $ tagger tag")
	}

	tag := os.Args[1]

	apiKey, err := getApiKey()
	if err != nil {
		log.Fatalf("Failed to get API key: %s", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: apiKey},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	//log.Printf("client: %#v", client)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.ListByOrg("venicegeo", nil)
	if err != nil {
		log.Fatalf("oauth failed: %s", err)
	}
	//log.Printf("%#v", repos)

	/*repos, err := getRepoNames(apiKey, 0, 100)
	if err != nil {
		log.Fatalf("Failed to get repo names: %s", err)
	}*/

	for _, repo := range repos {
		err = tagRepo(client, repo, tag)
		if err != nil {
			log.Fatalf("Failed to get repo names: %s", err)
		}
	}
}

func getRepoNames(apiKey string, pageStart int, pageCount int) ([]string, error) {
	return nil, nil
}

func GetLatestCommit(owner, repo string, sgc *github.Client) (string, string, error) {
	commits, res, err := sgc.Repositories.ListCommits(owner, repo, &github.CommitsListOptions{})

	if err != nil {
		log.Printf("err: %s res: %s", err, res)
		return "", "", err
	}

	log.Printf("last commit: %s %s", *commits[0].SHA, *commits[0].Commit.URL)

	return *commits[0].SHA, *commits[0].Commit.URL, nil
}

func tagRepo(client *github.Client, repo github.Repository, tag string) error {

	log.Printf("%s %s %s", *repo.Name, *repo.FullName, *repo.URL)

	sha, url, err := GetLatestCommit("venicegeo", *repo.Name, client)
	if err != nil {
		log.Fatalf("get latest commit failed: %s", err)
	}

	tag = "testtag03"
	mssg := "mssg"

	now := time.Now()
	who := "mpgerlek"
	email := "mpg@flaxen.com"

	typ := "commit"

	tagobj := &github.Tag{
		Tag:     &tag,
		Message: &mssg,
		Object:  &github.GitObject{Type: &typ, SHA: &sha, URL: &url},
		Tagger:  &github.CommitAuthor{Name: &who, Email: &email, Date: &now},
	}

	tagobj, resp, err := client.Git.CreateTag("venicegeo", *repo.Name, tagobj)
	if err != nil {
		log.Fatalf("create tag failed: %s", err)
	}
	log.Printf("tagobj: %#v", *tagobj.URL)
	refstr := "refs/heads/master"
	ref := github.Reference{Ref: &refstr, Object: &github.GitObject{SHA: tagobj.SHA}}
	_, resp, err = client.Git.CreateRef("venicegeo", *repo.Name, &ref)
	if err != nil {
		log.Fatalf("create ref failed: %s", err)
	}

	log.Printf("%#v", resp)

	os.Exit(9)

	return nil
}
