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

const DEBUG = false
const VENICE = "venicegeo"

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
	repos, err := getRepoNames(client)

	for _, repo := range repos {
		err, ref := tagRepo(client, repo, tag)
		if err != nil {
			log.Fatalf("Failed to get repo names: %s", err)
		}

		_ = ref

		os.Exit(9)
	}
}

func getRepoNames(client *github.Client) ([]github.Repository, error) {
	opts0 := github.ListOptions{
		Page:    0,
		PerPage: 512,
	}
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: opts0,
		Type:        "all",
	}

	repos, _, err := client.Repositories.ListByOrg(VENICE, opts)
	if err != nil {
		log.Fatalf("get repos failed: %s", err)
	}

	//log.Printf("%#v", repos)

	return repos, nil
}

func getLatestCommit(client *github.Client, repo string) (string, string, error) {

	opts := &github.CommitsListOptions{}

	commits, res, err := client.Repositories.ListCommits(VENICE, repo, opts)
	if err != nil {
		log.Fatalf("err: %s res: %s", err, res)
		return "", "", err
	}

	log.Printf("LAST COMMIT: %s %s", *commits[0].SHA, *commits[0].Commit.URL)

	return *commits[0].SHA, *commits[0].Commit.URL, nil
}

func createTag(client *github.Client,
	repo github.Repository,
	url string,
	sha string) (*github.Tag, error) {

	tag := "testtag03"
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

	tagobj, resp, err := client.Git.CreateTag(VENICE, *repo.Name, tagobj)
	if err != nil {
		return nil, err
	}

	log.Printf("CREATE TARG RESP: %#v", resp)

	return tagobj, nil
}

func createReference(client *github.Client,
	repo github.Repository,
	tagobj *github.Tag) (*github.Reference, error) {

	refstr := "refs/heads/master"

	ref := &github.Reference{Ref: &refstr, Object: &github.GitObject{SHA: tagobj.SHA}}

	ref2, resp, err := client.Git.CreateRef(VENICE, *repo.Name, ref)
	if err != nil {
		return nil, err
	}

	log.Printf("CREATE REF RESP: %#v", resp)

	return ref2, nil
}

func tagRepo(client *github.Client, repo github.Repository, tag string) (*github.Reference, error) {

	log.Printf("TAG REPO: %s %s %s", *repo.Name, *repo.FullName, *repo.URL)

	sha, url, err := getLatestCommit(client, *repo.Name)
	if err != nil {
		log.Fatalf("get latest commit failed: %s", err)
	}

	tagobj, err := createTag(client, repo, url, sha)
	if err != nil {
		log.Fatalf("create tag failed: %s", err)
	}

	log.Printf("tagobj: %#v", *tagobj.URL)

	refobj, err := createReference(client, repo, tagobj)
	if err != nil {
		log.Fatalf("create ref failed: %s", err)
	}

	return refobj, nil
}
